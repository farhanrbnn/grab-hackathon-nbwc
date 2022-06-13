package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WalletOrigin struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Company   string             `json:"company,omitempty" bson:"company,omitempty"`
	CreatedAt int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt int64              `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type WalletOriginRequest struct {
	Id      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=3,max=50"`
	Company string             `json:"company,omitempty" bson:"company,omitempty" validate:"required,min=3,max=100"`
}

type WalletOriginUsecase interface {
	First(context.Context, *WalletOrigin) (WalletOrigin, error)
	Fetch(context.Context) ([]WalletOrigin, error)
	Store(context.Context, *WalletOriginRequest) error
	Update(context.Context, *WalletOrigin, *WalletOriginRequest) error
	Delete(context.Context, *WalletOrigin) error
}

type WalletOriginRepository interface {
	First(context.Context, *WalletOrigin) (WalletOrigin, error)
	Fetch(context.Context) ([]WalletOrigin, error)
	Store(context.Context, *WalletOrigin) error
	Update(context.Context, *WalletOrigin) error
	Delete(context.Context, *WalletOrigin) error
}
