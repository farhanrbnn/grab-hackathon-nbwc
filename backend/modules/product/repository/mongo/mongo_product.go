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

type mongoProductRepository struct {
	Product *mongo.Collection
}

// NewMongoProductRepository will create an implementation of product.Repository
func NewMongoProductRepository(product *mongo.Collection) domain.ProductRepository {
	return &mongoProductRepository{
		Product: product,
	}
}

func (m *mongoProductRepository) First(ctx context.Context, productFilter *domain.Product) (product domain.Product, err error) {
	filter := bson.D{{"$and", bson.A{
		productFilter,
		bson.M{"deleted_at": nil}},
	}}

	if err = m.Product.FindOne(ctx, filter).Decode(&product); err != nil {
		return domain.Product{}, helper.ErrNotFound
	}

	return
}

func (m *mongoProductRepository) FetchByMerchantId(ctx context.Context, productFilter *domain.Product) (res []domain.Product, err error) {
	res = make([]domain.Product, 0)

	filter := bson.D{{"$and", bson.A{
		productFilter,
		bson.M{"deleted_at": nil}},
	}}

	cursor, err := m.Product.Find(ctx, filter)

	if err != nil {
		return make([]domain.Product, 0), err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product bson.M
		var pd domain.Product

		if err = cursor.Decode(&product); err != nil {
			return make([]domain.Product, 0), err
		}

		bsonBytes, _ := bson.Marshal(product)
		bson.Unmarshal(bsonBytes, &pd)

		res = append(res, pd)
	}

	return
}

func (m *mongoProductRepository) Fetch(ctx context.Context) (res []domain.Product, err error) {
	res = make([]domain.Product, 0)

	cursor, err := m.Product.Find(ctx, bson.M{"deleted_at": nil})

	if err != nil {
		return make([]domain.Product, 0), err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product bson.M
		var pd domain.Product

		if err = cursor.Decode(&product); err != nil {
			return make([]domain.Product, 0), err
		}

		bsonBytes, _ := bson.Marshal(product)
		bson.Unmarshal(bsonBytes, &pd)

		res = append(res, pd)
	}

	return
}

func (m *mongoProductRepository) Store(ctx context.Context, product *domain.Product) (err error) {
	product.CreatedAt = time.Now().Unix()

	res, err := m.Product.InsertOne(ctx, product)

	product.Id = res.InsertedID.(primitive.ObjectID)

	return
}

func (m *mongoProductRepository) Update(ctx context.Context, product *domain.Product) (err error) {
	_, err = m.Product.ReplaceOne(ctx, bson.M{"_id": product.Id}, product)

	if err != nil {
		return err
	}

	return
}

func (m *mongoProductRepository) Delete(ctx context.Context, product *domain.Product) (err error) {
	product.DeletedAt = time.Now().Unix()

	_, err = m.Product.UpdateOne(
		ctx,
		bson.M{"_id": product.Id},
		bson.D{
			{"$set", bson.D{{"deleted_at", product.DeletedAt}}},
		},
	)

	if err != nil {
		return err
	}

	return
}
