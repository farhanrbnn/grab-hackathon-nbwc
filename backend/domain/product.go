package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	MerchantId  primitive.ObjectID `json:"merchant_id,omitempty" bson:"merchant_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Images      []string           `json:"images,omitempty" bson:"images,omitempty"`
	Price       int64              `json:"price,omitempty" bson:"price,omitempty"`
	Dimensions  Dimensions         `json:"dimensions,omitempty" bson:"dimensions,omitempty"`
	CreatedAt   int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt   int64              `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type Dimensions struct {
	Height int64 `json:"height,omitempty" bson:"height,omitempty"`
	Width  int64 `json:"width,omitempty" bson:"width,omitempty"`
	Depth  int64 `json:"depth,omitempty" bson:"depth,omitempty"`
	Weight int64 `json:"weight,omitempty" bson:"weight,omitempty"`
}

type ProductStoreRequest struct {
	MerchantId  primitive.ObjectID `json:"merchant_id,omitempty" bson:"merchant_id,omitempty" validate:"required"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=8,max=24"`
	Description string             `json:"description,omitempty" bson:"description,omitempty" validate:"required,min=8,max=500"`
	Images      []string           `json:"images,omitempty" bson:"images,omitempty"`
	Price       int64              `json:"price,omitempty" bson:"price,omitempty" validate:"required,gte=0"`
	Dimensions  Dimensions         `json:"dimensions,omitempty" bson:"dimensions,omitempty"`
}

type ProductUpdateRequest struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=8,max=24"`
	Description string             `json:"description,omitempty" bson:"description,omitempty" validate:"required,min=8,max=500"`
	Images      []string           `json:"images,omitempty" bson:"images,omitempty"`
	Price       int64              `json:"price,omitempty" bson:"price,omitempty" validate:"required,gte=0"`
	Dimensions  Dimensions         `json:"dimensions,omitempty" bson:"dimensions,omitempty"`
}

type ProductUsecase interface {
	First(ctx context.Context, productFilter *Product) (Product, error)
	FetchByMerchantId(ctx context.Context, productFilter *Product) ([]Product, error)
	Fetch(ctx context.Context) (res []Product, err error)
	Store(context.Context, *ProductStoreRequest) error
	Update(context.Context, *Product, *ProductUpdateRequest) error
	Delete(ctx context.Context, user *Product) error
}

type ProductRepository interface {
	First(ctx context.Context, productFilter *Product) (Product, error)
	FetchByMerchantId(ctx context.Context, productFilter *Product) ([]Product, error)
	Fetch(ctx context.Context) ([]Product, error)
	Store(ctx context.Context, c *Product) error
	Update(context.Context, *Product) error
	Delete(ctx context.Context, user *Product) error
}
