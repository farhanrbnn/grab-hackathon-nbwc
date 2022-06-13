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

type mongoMerchantRepository struct {
	Merchant *mongo.Collection
}

// NewMongoMerchantRepository will create an implementation of merchant.Repository
func NewMongoMerchantRepository(merchant *mongo.Collection) domain.MerchantRepository {
	return &mongoMerchantRepository{
		Merchant: merchant,
	}
}

func (m *mongoMerchantRepository) First(ctx context.Context, merchantFilter *domain.Merchant) (merchant domain.Merchant, err error) {
	filter := bson.D{{"$and", bson.A{
		bson.M{"_id": merchantFilter.Id},
		bson.M{"deleted_at": nil}},
	}}

	if err = m.Merchant.FindOne(ctx, filter).Decode(&merchant); err != nil {
		return domain.Merchant{}, helper.ErrNotFound
	}

	return
}

func (m *mongoMerchantRepository) Fetch(ctx context.Context) (res []domain.Merchant, err error) {
	res = make([]domain.Merchant, 0)

	cursor, err := m.Merchant.Find(ctx, bson.M{"deleted_at": nil})

	if err != nil {
		return make([]domain.Merchant, 0), err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var merchant bson.M
		var pd domain.Merchant

		if err = cursor.Decode(&merchant); err != nil {
			return make([]domain.Merchant, 0), err
		}

		bsonBytes, _ := bson.Marshal(merchant)
		bson.Unmarshal(bsonBytes, &pd)

		res = append(res, pd)
	}

	return
}

func (m *mongoMerchantRepository) Store(ctx context.Context, merchant *domain.Merchant) (err error) {
	merchant.CreatedAt = time.Now().Unix()

	res, err := m.Merchant.InsertOne(ctx, merchant)

	merchant.Id = res.InsertedID.(primitive.ObjectID)

	return
}

func (m *mongoMerchantRepository) Update(ctx context.Context, merchant *domain.Merchant) (err error) {
	_, err = m.Merchant.ReplaceOne(ctx, bson.M{"_id": merchant.Id}, merchant)

	if err != nil {
		return err
	}

	return
}

func (m *mongoMerchantRepository) Delete(ctx context.Context, merchant *domain.Merchant) (err error) {
	merchant.DeletedAt = time.Now().Unix()

	_, err = m.Merchant.UpdateOne(
		ctx,
		bson.M{"_id": merchant.Id},
		bson.D{
			{"$set", bson.D{{"deleted_at", merchant.DeletedAt}}},
		},
	)

	if err != nil {
		return err
	}

	return
}
