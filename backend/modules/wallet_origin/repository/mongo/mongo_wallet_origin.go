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

type mongoWalletOriginRepository struct {
	WalletOrigin *mongo.Collection
}

// NewMongoWalletOriginRepository will create an implementation of walletOrigin.Repository
func NewMongoWalletOriginRepository(walletOrigin *mongo.Collection) domain.WalletOriginRepository {
	return &mongoWalletOriginRepository{
		WalletOrigin: walletOrigin,
	}
}

func (m *mongoWalletOriginRepository) First(ctx context.Context, walletOriginFilter *domain.WalletOrigin) (walletOrigin domain.WalletOrigin, err error) {
	filter := bson.D{{"$and", bson.A{
		walletOriginFilter,
		bson.M{"deleted_at": nil}},
	}}

	if err = m.WalletOrigin.FindOne(ctx, filter).Decode(&walletOrigin); err != nil {
		return domain.WalletOrigin{}, helper.ErrNotFound
	}

	return
}

func (m *mongoWalletOriginRepository) Fetch(ctx context.Context) (res []domain.WalletOrigin, err error) {
	res = make([]domain.WalletOrigin, 0)

	cursor, err := m.WalletOrigin.Find(ctx, bson.M{"deleted_at": nil})

	if err != nil {
		return make([]domain.WalletOrigin, 0), err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var walletOrigin bson.M
		var pd domain.WalletOrigin

		if err = cursor.Decode(&walletOrigin); err != nil {
			return make([]domain.WalletOrigin, 0), err
		}

		bsonBytes, _ := bson.Marshal(walletOrigin)
		bson.Unmarshal(bsonBytes, &pd)

		res = append(res, pd)
	}

	return
}

func (m *mongoWalletOriginRepository) Store(ctx context.Context, walletOrigin *domain.WalletOrigin) (err error) {
	walletOrigin.CreatedAt = time.Now().Unix()

	res, err := m.WalletOrigin.InsertOne(ctx, walletOrigin)

	walletOrigin.Id = res.InsertedID.(primitive.ObjectID)

	return
}

func (m *mongoWalletOriginRepository) Update(ctx context.Context, walletOrigin *domain.WalletOrigin) (err error) {
	_, err = m.WalletOrigin.ReplaceOne(ctx, bson.M{"_id": walletOrigin.Id}, walletOrigin)

	if err != nil {
		return err
	}

	return
}

func (m *mongoWalletOriginRepository) Delete(ctx context.Context, walletOrigin *domain.WalletOrigin) (err error) {
	walletOrigin.DeletedAt = time.Now().Unix()

	_, err = m.WalletOrigin.UpdateOne(
		ctx,
		bson.M{"_id": walletOrigin.Id},
		bson.D{
			{"$set", bson.D{{"deleted_at", walletOrigin.DeletedAt}}},
		},
	)

	if err != nil {
		return err
	}

	return
}
