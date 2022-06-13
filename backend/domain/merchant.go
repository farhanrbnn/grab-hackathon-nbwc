package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AWS Merchant ...
type Merchant struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	Coordinate  Coordinate         `json:"coordinate,omitempty" bson:"coordinate,omitempty"`
	Rating      float32            `json:"rating,omitempty" bson:"rating,omitempty"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	Thumbnail   string             `json:"thumbnail,omitempty" bson:"thumbnail,omitempty"`
	IsAvailable bool               `json:"is_available,omitempty" bson:"is_available,omitempty"`
	CreatedAt   int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt   int64              `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type MerchantStoreRequest struct {
	Name        string     `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=8,max=100"`
	Phone       string     `json:"phone,omitempty" bson:"phone,omitempty" validate:"required,min=8,max=16"`
	Email       string     `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Coordinate  Coordinate `json:"coordinate,omitempty" bson:"coordinate,omitempty"`
	Rating      float32    `json:"rating,omitempty" bson:"rating,omitempty" validate:"required,gte=0,lte=5"`
	Address     string     `json:"address,omitempty" bson:"address,omitempty" validate:"required,min=8,max=500"`
	Thumbnail   string     `json:"thumbnail,omitempty" bson:"thumbnail,omitempty" validate:"required"`
	IsAvailable bool       `json:"is_available,omitempty" bson:"is_available,omitempty" validate:"required"`
}

type MerchantUpdateRequest struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=8,max=100"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"required,min=8,max=16"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Coordinate  Coordinate         `json:"coordinate,omitempty" bson:"coordinate,omitempty"`
	Rating      float32            `json:"rating,omitempty" bson:"rating,omitempty" validate:"required,gte=0,lte=5"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty" validate:"required,min=8,max=500"`
	Thumbnail   string             `json:"thumbnail,omitempty" bson:"thumbnail,omitempty" validate:"required"`
	IsAvailable bool               `json:"is_available,omitempty" bson:"is_available,omitempty" validate:"required"`
}

type MerchantUsecase interface {
	First(ctx context.Context, merchantFilter *Merchant) (Merchant, error)
	Fetch(ctx context.Context) (res []Merchant, err error)
	Store(context.Context, *MerchantStoreRequest) error
	Update(context.Context, *Merchant, *MerchantUpdateRequest) error
	Delete(ctx context.Context, user *Merchant) error
}

type MerchantRepository interface {
	First(ctx context.Context, merchantFilter *Merchant) (Merchant, error)
	Fetch(ctx context.Context) ([]Merchant, error)
	Store(ctx context.Context, c *Merchant) error
	Update(context.Context, *Merchant) error
	Delete(ctx context.Context, user *Merchant) error
}
