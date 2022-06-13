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

type mongoTransactionRepository struct {
	Transaction *mongo.Collection
}

// NewMongoTransactionRepository will create an implementation of transaction.Repository
func NewMongoTransactionRepository(transaction *mongo.Collection) domain.TransactionRepository {
	return &mongoTransactionRepository{
		Transaction: transaction,
	}
}

func (m *mongoTransactionRepository) First(ctx context.Context, transactionFilter *domain.Transaction) (transaction domain.Transaction, err error) {
	filter := bson.D{{"$and", bson.A{
		bson.M{"_id": transactionFilter.Id},
		bson.M{"deleted_at": nil}},
	}}

	if err = m.Transaction.FindOne(ctx, filter).Decode(&transaction); err != nil {
		return domain.Transaction{}, helper.ErrNotFound
	}

	return
}

func (m *mongoTransactionRepository) FetchByUserId(ctx context.Context, transactionFilter *domain.Transaction) (res []domain.Transaction, err error) {
	res = make([]domain.Transaction, 0)

	filter := bson.D{{"$and", bson.A{
		transactionFilter,
		bson.M{"deleted_at": nil}},
	}}

	cursor, err := m.Transaction.Find(ctx, filter)

	if err != nil {
		return make([]domain.Transaction, 0), err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var transaction bson.M
		var pd domain.Transaction

		if err = cursor.Decode(&transaction); err != nil {
			return make([]domain.Transaction, 0), err
		}

		bsonBytes, _ := bson.Marshal(transaction)
		bson.Unmarshal(bsonBytes, &pd)

		res = append(res, pd)
	}

	return
}

func (m *mongoTransactionRepository) Fetch(ctx context.Context) (res []domain.Transaction, err error) {
	res = make([]domain.Transaction, 0)

	cursor, err := m.Transaction.Find(ctx, bson.M{"deleted_at": nil})

	if err != nil {
		return make([]domain.Transaction, 0), err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var transaction bson.M
		var pd domain.Transaction

		if err = cursor.Decode(&transaction); err != nil {
			return make([]domain.Transaction, 0), err
		}

		bsonBytes, _ := bson.Marshal(transaction)
		bson.Unmarshal(bsonBytes, &pd)

		res = append(res, pd)
	}

	return
}

func (m *mongoTransactionRepository) Store(ctx context.Context, transaction *domain.Transaction) (err error) {
	transaction.CreatedAt = time.Now().Unix()

	res, err := m.Transaction.InsertOne(ctx, transaction)

	transaction.Id = res.InsertedID.(primitive.ObjectID)

	return
}

func (m *mongoTransactionRepository) Update(ctx context.Context, transaction *domain.Transaction) (err error) {
	_, err = m.Transaction.ReplaceOne(ctx, bson.M{"_id": transaction.Id}, transaction)

	if err != nil {
		return err
	}

	return
}

func (m *mongoTransactionRepository) Delete(ctx context.Context, transaction *domain.Transaction) (err error) {
	transaction.DeletedAt = time.Now().Unix()

	_, err = m.Transaction.UpdateOne(
		ctx,
		bson.M{"_id": transaction.Id},
		bson.D{
			{"$set", bson.D{{"deleted_at", transaction.DeletedAt}}},
		},
	)

	if err != nil {
		return err
	}

	return
}
