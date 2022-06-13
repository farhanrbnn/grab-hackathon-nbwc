package mongo

import (
	"context"
	"time"

	"grab-hack-for-good/domain"
	"grab-hack-for-good/helper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoWalletRepository struct {
	Wallet *mongo.Collection
}

// NewMongoWalletRepository will create an implementation of wallet.Repository
func NewMongoWalletRepository(wallet *mongo.Collection) domain.WalletRepository {
	return &mongoWalletRepository{
		Wallet: wallet,
	}
}

func (m *mongoWalletRepository) First(ctx context.Context, walletFilter *domain.Wallet) (wallet domain.Wallet, err error) {
	filter := bson.D{{"$and", bson.A{
		walletFilter,
		bson.M{"deleted_at": nil}},
	}}

	if err = m.Wallet.FindOne(ctx, filter).Decode(&wallet); err != nil {
		return domain.Wallet{}, helper.ErrNotFound
	}

	return
}

func (m *mongoWalletRepository) Fetch(ctx context.Context) (res []domain.Wallet, err error) {
	res = make([]domain.Wallet, 0)

	cursor, err := m.Wallet.Find(ctx, bson.M{"deleted_at": nil})

	if err != nil {
		return make([]domain.Wallet, 0), err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var wallet bson.M
		var pd domain.Wallet

		if err = cursor.Decode(&wallet); err != nil {
			return make([]domain.Wallet, 0), err
		}

		bsonBytes, _ := bson.Marshal(wallet)
		bson.Unmarshal(bsonBytes, &pd)

		res = append(res, pd)
	}

	return
}

func (m *mongoWalletRepository) Store(ctx context.Context, wallet *domain.Wallet) (err error) {
	wallet.CreatedAt = time.Now().Unix()

	res, err := m.Wallet.InsertOne(ctx, wallet)

	wallet.Id = res.InsertedID.(primitive.ObjectID)

	return
}

func (m *mongoWalletRepository) Update(ctx context.Context, wallet *domain.Wallet) (err error) {
	_, err = m.Wallet.ReplaceOne(ctx, bson.M{"_id": wallet.Id}, wallet)

	if err != nil {
		return err
	}

	return
}

func (m *mongoWalletRepository) Delete(ctx context.Context, wallet *domain.Wallet) (err error) {
	wallet.DeletedAt = time.Now().Unix()

	_, err = m.Wallet.UpdateOne(
		ctx,
		bson.M{"_id": wallet.Id},
		bson.D{
			{"$set", bson.D{{"deleted_at", wallet.DeletedAt}}},
		},
	)

	if err != nil {
		return err
	}

	return
}
