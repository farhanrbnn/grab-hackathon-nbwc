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

type mongoOrderRepository struct {
	Order *mongo.Collection
}

// NewMongoOrderRepository will create an implementation of order.Repository
func NewMongoOrderRepository(order *mongo.Collection) domain.OrderRepository {
	return &mongoOrderRepository{
		Order: order,
	}
}

func (m *mongoOrderRepository) First(ctx context.Context, orderFilter *domain.Order) (order domain.Order, err error) {
	filter := bson.D{{"$and", bson.A{
		bson.M{"_id": orderFilter.Id},
		bson.M{"deleted_at": nil}},
	}}

	if err = m.Order.FindOne(ctx, filter).Decode(&order); err != nil {
		return domain.Order{}, helper.ErrNotFound
	}

	return
}

func (m *mongoOrderRepository) FetchByUserId(ctx context.Context, orderFilter *domain.Order) (res []domain.Order, err error) {
	res = make([]domain.Order, 0)

	filter := bson.D{{"$and", bson.A{
		bson.M{"user_id": orderFilter.UserId},
		bson.M{"deleted_at": nil}},
	}}

	cursor, err := m.Order.Find(ctx, filter)

	if err != nil {
		return make([]domain.Order, 0), err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order bson.M
		var pd domain.Order

		if err = cursor.Decode(&order); err != nil {
			return make([]domain.Order, 0), err
		}

		bsonBytes, _ := bson.Marshal(order)
		bson.Unmarshal(bsonBytes, &pd)

		res = append(res, pd)
	}

	return
}

func (m *mongoOrderRepository) Fetch(ctx context.Context) (res []domain.Order, err error) {
	res = make([]domain.Order, 0)

	cursor, err := m.Order.Find(ctx, bson.M{"deleted_at": nil})

	if err != nil {
		return make([]domain.Order, 0), err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order bson.M
		var pd domain.Order

		if err = cursor.Decode(&order); err != nil {
			return make([]domain.Order, 0), err
		}

		bsonBytes, _ := bson.Marshal(order)
		bson.Unmarshal(bsonBytes, &pd)

		res = append(res, pd)
	}

	return
}

func (m *mongoOrderRepository) Store(ctx context.Context, order *domain.Order) (err error) {
	current_time := time.Now().Unix()
	order.CreatedAt = &current_time

	res, err := m.Order.InsertOne(ctx, order)

	order.Id = res.InsertedID.(primitive.ObjectID)

	return
}

func (m *mongoOrderRepository) Update(ctx context.Context, order *domain.Order) (err error) {
	current_time := time.Now().Unix()
	order.UpdatedAt = &current_time

	_, err = m.Order.ReplaceOne(ctx, bson.M{"_id": order.Id}, order)

	if err != nil {
		return err
	}

	return
}

func (m *mongoOrderRepository) Delete(ctx context.Context, order *domain.Order) (err error) {
	current_time := time.Now().Unix()
	order.DeletedAt = &current_time

	_, err = m.Order.UpdateOne(
		ctx,
		bson.M{"_id": order.Id},
		bson.D{
			{"$set", bson.D{{"deleted_at", order.DeletedAt}}},
		},
	)

	if err != nil {
		return err
	}

	return
}
