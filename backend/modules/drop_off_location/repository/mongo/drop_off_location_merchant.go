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

type mongoDropOffLocationRepository struct {
	DropOffLocation *mongo.Collection
}

// NewMongoDropOffLocationRepository will create an implementation of dropOffLocation.Repository
func NewMongoDropOffLocationRepository(dropOffLocation *mongo.Collection) domain.DropOffLocationRepository {
	return &mongoDropOffLocationRepository{
		DropOffLocation: dropOffLocation,
	}
}

func (m *mongoDropOffLocationRepository) First(ctx context.Context, dropOffLocationFilter *domain.DropOffLocation) (dropOffLocation domain.DropOffLocation, err error) {
	filter := bson.D{{"$and", bson.A{
		bson.M{"_id": dropOffLocationFilter.Id},
		bson.M{"deleted_at": nil}},
	}}

	if err = m.DropOffLocation.FindOne(ctx, filter).Decode(&dropOffLocation); err != nil {
		return domain.DropOffLocation{}, helper.ErrNotFound
	}

	return
}

func (m *mongoDropOffLocationRepository) Fetch(ctx context.Context) (res []domain.DropOffLocation, err error) {
	res = make([]domain.DropOffLocation, 0)

	cursor, err := m.DropOffLocation.Find(ctx, bson.M{"deleted_at": nil})

	if err != nil {
		return make([]domain.DropOffLocation, 0), err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var dropOffLocation bson.M
		var pd domain.DropOffLocation

		if err = cursor.Decode(&dropOffLocation); err != nil {
			return make([]domain.DropOffLocation, 0), err
		}

		bsonBytes, _ := bson.Marshal(dropOffLocation)
		bson.Unmarshal(bsonBytes, &pd)

		res = append(res, pd)
	}

	return
}

func (m *mongoDropOffLocationRepository) Store(ctx context.Context, dropOffLocation *domain.DropOffLocation) (err error) {
	dropOffLocation.CreatedAt = time.Now().Unix()

	res, err := m.DropOffLocation.InsertOne(ctx, dropOffLocation)

	dropOffLocation.Id = res.InsertedID.(primitive.ObjectID)

	return
}

func (m *mongoDropOffLocationRepository) Update(ctx context.Context, dropOffLocation *domain.DropOffLocation) (err error) {
	_, err = m.DropOffLocation.ReplaceOne(ctx, bson.M{"_id": dropOffLocation.Id}, dropOffLocation)

	if err != nil {
		return err
	}

	return
}

func (m *mongoDropOffLocationRepository) Delete(ctx context.Context, dropOffLocation *domain.DropOffLocation) (err error) {
	dropOffLocation.DeletedAt = time.Now().Unix()

	_, err = m.DropOffLocation.UpdateOne(
		ctx,
		bson.M{"_id": dropOffLocation.Id},
		bson.D{
			{"$set", bson.D{{"deleted_at", dropOffLocation.DeletedAt}}},
		},
	)

	if err != nil {
		return err
	}

	return
}
