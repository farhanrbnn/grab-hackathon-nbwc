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

type mongoUserRepository struct {
	User *mongo.Collection
}

// NewMongoUserRepository will create an implementation of user.Repository
func NewMongoUserRepository(user *mongo.Collection) domain.UserRepository {
	return &mongoUserRepository{
		User: user,
	}
}

func (m *mongoUserRepository) First(ctx context.Context, userFilter *domain.User) (user domain.User, err error) {
	filter := bson.D{{"$and", bson.A{
		userFilter,
		bson.M{"deleted_at": nil}},
	}}

	if err = m.User.FindOne(ctx, filter).Decode(&user); err != nil {
		return domain.User{}, helper.ErrNotFound
	}

	return
}

func (m *mongoUserRepository) Fetch(ctx context.Context) (res []domain.User, err error) {
	res = make([]domain.User, 0)

	cursor, err := m.User.Find(ctx, bson.M{"deleted_at": nil})

	if err != nil {
		return make([]domain.User, 0), err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user bson.M
		var pd domain.User

		if err = cursor.Decode(&user); err != nil {
			return make([]domain.User, 0), err
		}

		bsonBytes, _ := bson.Marshal(user)
		bson.Unmarshal(bsonBytes, &pd)

		res = append(res, pd)
	}

	return
}

func (m *mongoUserRepository) Store(ctx context.Context, user *domain.User) (err error) {
	user.CreatedAt = time.Now().Unix()

	res, err := m.User.InsertOne(ctx, user)

	user.Id = res.InsertedID.(primitive.ObjectID)

	return
}

func (m *mongoUserRepository) Update(ctx context.Context, user *domain.User) (err error) {
	_, err = m.User.ReplaceOne(ctx, bson.M{"_id": user.Id}, user)

	if err != nil {
		return err
	}

	return
}

func (m *mongoUserRepository) Delete(ctx context.Context, user *domain.User) (err error) {
	user.DeletedAt = time.Now().Unix()

	_, err = m.User.UpdateOne(
		ctx,
		bson.M{"_id": user.Id},
		bson.D{
			{"$set", bson.D{{"deleted_at", user.DeletedAt}}},
		},
	)

	if err != nil {
		return err
	}

	return
}
