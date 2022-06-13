package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wallet struct {
	Id             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	WalletOriginId primitive.ObjectID `json:"wallet_origin_id,omitempty" bson:"wallet_origin_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Image          string             `json:"image,omitempty" bson:"image,omitempty"`
	CreatedAt      int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt      int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt      int64              `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type WalletStoreRequest struct {
	WalletOriginId primitive.ObjectID `json:"wallet_origin_id,omitempty" bson:"wallet_origin_id,omitempty" validate:"required"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=3,max=50"`
	Image          string             `json:"image,omitempty" bson:"image,omitempty" validate:"required"`
}

type WalletUpdateRequest struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=3,max=50"`
	Image string             `json:"image,omitempty" bson:"image,omitempty" validate:"required"`
}

type WalletUsecase interface {
	First(ctx context.Context, walletFilter *Wallet) (Wallet, error)
	Fetch(ctx context.Context) (res []Wallet, err error)
	Store(ctx context.Context, wallet *WalletStoreRequest) error
	Update(ctx context.Context, wallet *Wallet, walletRequest *WalletUpdateRequest) error
	Delete(ctx context.Context, wallet *Wallet) error
}

type WalletRepository interface {
	First(ctx context.Context, walletFilter *Wallet) (Wallet, error)
	Fetch(ctx context.Context) ([]Wallet, error)
	Store(context.Context, *Wallet) error
	Update(context.Context, *Wallet) error
	Delete(ctx context.Context, wallet *Wallet) error
}
