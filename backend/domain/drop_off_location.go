package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AWS DropOffLocation ...
type DropOffLocation struct {
	Id            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	PIC           string             `json:"pic,omitempty" bson:"pic,omitempty"`
	Phone         string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	Coordinate    Coordinate         `json:"coordinate,omitempty" bson:"coordinate,omitempty"`
	Address       string             `json:"address,omitempty" bson:"address,omitempty"`
	MaxSupply     int64              `json:"max_supply,omitempty" bson:"max_supply,omitempty"`
	CurrentSupply int64              `json:"current_supply,omitempty" bson:"current_supply,omitempty"`
	CreatedAt     int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt     int64              `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type DropOffLocationStoreRequest struct {
	Name          string     `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=8,max=100"`
	PIC           string     `json:"pic,omitempty" bson:"pic,omitempty"`
	Phone         string     `json:"phone,omitempty" bson:"phone,omitempty" validate:"required,min=8,max=16"`
	Email         string     `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Coordinate    Coordinate `json:"coordinate,omitempty" bson:"coordinate,omitempty" validate:"required"`
	Address       string     `json:"address,omitempty" bson:"address,omitempty" validate:"required,min=8,max=500"`
	MaxSupply     int64      `json:"max_supply,omitempty" bson:"max_supply,omitempty" validate:"required,gte=0"`
	CurrentSupply int64      `json:"current_supply,omitempty" bson:"current_supply,omitempty" validate:"required,gte=0"`
}

type DropOffLocationUpdateRequest struct {
	Id            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=8,max=100"`
	PIC           string             `json:"pic,omitempty" bson:"pic,omitempty"`
	Phone         string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"required,min=8,max=16"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Coordinate    Coordinate         `json:"coordinate,omitempty" bson:"coordinate,omitempty" validate:"required"`
	Address       string             `json:"address,omitempty" bson:"address,omitempty" validate:"required,min=8,max=500"`
	MaxSupply     int64              `json:"max_supply,omitempty" bson:"max_supply,omitempty" validate:"required,gte=0"`
	CurrentSupply int64              `json:"current_supply,omitempty" bson:"current_supply,omitempty" validate:"required,gte=0"`
}

type DropOffLocationUsecase interface {
	First(ctx context.Context, dropOffLocationFilter *DropOffLocation) (DropOffLocation, error)
	Fetch(ctx context.Context) (res []DropOffLocation, err error)
	Store(context.Context, *DropOffLocationStoreRequest) error
	Update(context.Context, *DropOffLocation, *DropOffLocationUpdateRequest) error
	Delete(ctx context.Context, user *DropOffLocation) error
}

type DropOffLocationRepository interface {
	First(ctx context.Context, dropOffLocationFilter *DropOffLocation) (DropOffLocation, error)
	Fetch(ctx context.Context) ([]DropOffLocation, error)
	Store(ctx context.Context, c *DropOffLocation) error
	Update(context.Context, *DropOffLocation) error
	Delete(ctx context.Context, user *DropOffLocation) error
}
