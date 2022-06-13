package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserWallet struct {
	Id              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId          primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	WalletId        primitive.ObjectID `json:"wallet_id,omitempty" bson:"wallet_id,omitempty"`
	EffectiveAmount int64              `json:"effective_amount,omitempty" bson:"effective_amount,omitempty"`
	OnHoldAmount    int64              `json:"on_hold_amount,omitempty" bson:"on_hold_amount,omitempty"`
	CreatedAt       int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt       int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt       int64              `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type UserWalletStoreRequest struct {
	UserId          primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty" validate:"required"`
	WalletId        primitive.ObjectID `json:"wallet_id,omitempty" bson:"wallet_id,omitempty" validate:"required"`
	EffectiveAmount int64              `json:"effective_amount,omitempty" bson:"effective_amount,omitempty" validate:"required,gte=0"`
	OnHoldAmount    int64              `json:"on_hold_amount,omitempty" bson:"on_hold_amount,omitempty" validate:"gte=0"`
}

type UserWalletUpdateRequest struct {
	Id              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	EffectiveAmount int64              `json:"effective_amount,omitempty" bson:"effective_amount,omitempty" validate:"required,gte=0"`
	OnHoldAmount    int64              `json:"on_hold_amount,omitempty" bson:"on_hold_amount,omitempty" validate:"gte=0"`
}

type UserWalletTransactionRequest struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Amount int64              `json:"amount,omitempty" bson:"amount,omitempty" validate:"required,gte=0"`
}

type UserWalletUsecase interface {
	First(ctx context.Context, walletFilter *UserWallet) (UserWallet, error)
	FetchByUserId(ctx context.Context, userWalletFilter *UserWallet) ([]UserWallet, error)
	Fetch(ctx context.Context) (res []UserWallet, err error)
	Store(ctx context.Context, wallet *UserWalletStoreRequest) error
	Pay(ctx context.Context, wallet *UserWallet, userWalletTransactionRequest *UserWalletTransactionRequest) error
	Refund(ctx context.Context, wallet *UserWallet, userWalletTransactionRequest *UserWalletTransactionRequest) error
	Commit(ctx context.Context, wallet *UserWallet, userWalletTransactionRequest *UserWalletTransactionRequest) error
	Update(ctx context.Context, wallet *UserWallet, userWalletRequest *UserWalletUpdateRequest) error
	Delete(ctx context.Context, wallet *UserWallet) error
}

type UserWalletRepository interface {
	First(ctx context.Context, walletFilter *UserWallet) (UserWallet, error)
	FetchByUserId(ctx context.Context, walletFilter *UserWallet) ([]UserWallet, error)
	Fetch(ctx context.Context) ([]UserWallet, error)
	Store(context.Context, *UserWallet) error
	Update(context.Context, *UserWallet) error
	Delete(ctx context.Context, wallet *UserWallet) error
}
