package mongo

import (
	"grab-hack-for-good/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoGrabExpressRepository struct {
	GrabExpress *mongo.Collection
}

// NewMongoGrabExpressRepository will create an implementation of grabExpress.Repository
func NewMongoGrabExpressRepository(grabExpress *mongo.Collection) domain.GrabExpressRepository {
	return &mongoGrabExpressRepository{
		GrabExpress: grabExpress,
	}
}
