package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Transaction ...
type Transaction struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	OrderId      primitive.ObjectID `json:"order_id,omitempty" bson:"order_id,omitempty"`
	UserId       primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	UserWalletId primitive.ObjectID `json:"user_wallet_id,omitempty" bson:"user_wallet_id,omitempty"`
	Status       string             `json:"status,omitempty" bson:"status,omitempty"`
	Amount       int64              `json:"amount,omitempty" bson:"amount,omitempty"`
	CreatedAt    int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt    int64              `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type TransactionStoreRequest struct {
	OrderId      primitive.ObjectID `json:"order_id,omitempty" bson:"order_id,omitempty" validate:"required"`
	UserId       primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty" validate:"required"`
	UserWalletId primitive.ObjectID `json:"user_wallet_id,omitempty" bson:"user_wallet_id,omitempty" validate:"required"`
	Status       string             `json:"status,omitempty" bson:"status,omitempty" validate:"required,min=0,max=100"`
	Amount       int64              `json:"amount,omitempty" bson:"amount,omitempty" validate:"required,gte=0"`
}

type TransactionUpdateRequest struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Status string             `json:"status,omitempty" bson:"status,omitempty" validate:"required,min=0,max=100"`
	Amount int64              `json:"amount,omitempty" bson:"amount,omitempty" validate:"required,gte=0"`
}

type TransactionUsecase interface {
	First(ctx context.Context, transactionFilter *Transaction) (Transaction, error)
	FetchByUserId(ctx context.Context, transactionFilter *Transaction) ([]Transaction, error)
	Fetch(ctx context.Context) (res []Transaction, err error)
	Store(context.Context, *TransactionStoreRequest) (*Transaction, error)
	Update(context.Context, *Transaction, *TransactionUpdateRequest) error
	Delete(ctx context.Context, user *Transaction) error
}

type TransactionRepository interface {
	First(ctx context.Context, transactionFilter *Transaction) (Transaction, error)
	FetchByUserId(ctx context.Context, transactionFilter *Transaction) ([]Transaction, error)
	Fetch(ctx context.Context) ([]Transaction, error)
	Store(ctx context.Context, c *Transaction) error
	Update(context.Context, *Transaction) error
	Delete(ctx context.Context, user *Transaction) error
}
