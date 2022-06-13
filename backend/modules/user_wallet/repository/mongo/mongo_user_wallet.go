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

type mongoUserWalletRepository struct {
	UserWallet *mongo.Collection
}

// NewMongoUserWalletRepository will create an implementation of userWallet.Repository
func NewMongoUserWalletRepository(userWallet *mongo.Collection) domain.UserWalletRepository {
	return &mongoUserWalletRepository{
		UserWallet: userWallet,
	}
}

func (m *mongoUserWalletRepository) First(ctx context.Context, userWalletFilter *domain.UserWallet) (userWallet domain.UserWallet, err error) {
	filter := bson.D{{"$and", bson.A{
		userWalletFilter,
		bson.M{"deleted_at": nil}},
	}}

	if err = m.UserWallet.FindOne(ctx, filter).Decode(&userWallet); err != nil {
		return domain.UserWallet{}, helper.ErrNotFound
	}

	return
}

func (m *mongoUserWalletRepository) FetchByUserId(ctx context.Context, userWalletFilter *domain.UserWallet) (res []domain.UserWallet, err error) {
	res = make([]domain.UserWallet, 0)

	filter := bson.D{{"$and", bson.A{
		userWalletFilter,
		bson.M{"deleted_at": nil}},
	}}

	cursor, err := m.UserWallet.Find(ctx, filter)

	if err != nil {
		return make([]domain.UserWallet, 0), err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var userWallet bson.M
		var pd domain.UserWallet

		if err = cursor.Decode(&userWallet); err != nil {
			return make([]domain.UserWallet, 0), err
		}

		bsonBytes, _ := bson.Marshal(userWallet)
		bson.Unmarshal(bsonBytes, &pd)

		res = append(res, pd)
	}

	return
}

func (m *mongoUserWalletRepository) Fetch(ctx context.Context) (res []domain.UserWallet, err error) {
	res = make([]domain.UserWallet, 0)

	cursor, err := m.UserWallet.Find(ctx, bson.M{"deleted_at": nil})

	if err != nil {
		return make([]domain.UserWallet, 0), err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var userWallet bson.M
		var pd domain.UserWallet

		if err = cursor.Decode(&userWallet); err != nil {
			return make([]domain.UserWallet, 0), err
		}

		bsonBytes, _ := bson.Marshal(userWallet)
		bson.Unmarshal(bsonBytes, &pd)

		res = append(res, pd)
	}

	return
}

func (m *mongoUserWalletRepository) Store(ctx context.Context, userWallet *domain.UserWallet) (err error) {
	userWallet.CreatedAt = time.Now().Unix()

	res, err := m.UserWallet.InsertOne(ctx, userWallet)

	userWallet.Id = res.InsertedID.(primitive.ObjectID)

	return
}

func (m *mongoUserWalletRepository) Update(ctx context.Context, userWallet *domain.UserWallet) (err error) {
	_, err = m.UserWallet.ReplaceOne(ctx, bson.M{"_id": userWallet.Id}, userWallet)

	if err != nil {
		return err
	}

	return
}

func (m *mongoUserWalletRepository) Delete(ctx context.Context, userWallet *domain.UserWallet) (err error) {
	userWallet.DeletedAt = time.Now().Unix()

	_, err = m.UserWallet.UpdateOne(
		ctx,
		bson.M{"_id": userWallet.Id},
		bson.D{
			{"$set", bson.D{{"deleted_at", userWallet.DeletedAt}}},
		},
	)

	if err != nil {
		return err
	}

	return
}
