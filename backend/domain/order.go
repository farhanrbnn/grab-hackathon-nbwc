package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order ...
type Order struct {
	Id                   primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Transaction          *Transaction        `json:"transaction" bson:"transaction"`
	UserId               *primitive.ObjectID `json:"user_id" bson:"user_id"`
	MerchantOrderId      *string             `json:"merchant_order_id,omitempty" bson:"merchant_order_id,omitempty"`
	DeliveryId           *string             `json:"delivery_id,omitempty" bson:"delivery_id,omitempty"`
	DeliveryStatus       *string             `json:"delivery_status,omitempty" bson:"delivery_status,omitempty"`
	DeliveryFailedReason *string             `json:"delivery_failed_reason,omitempty" bson:"delivery_failed_reason,omitempty"`
	FundingSource        *FundingSource      `json:"funding_source,omitempty" bson:"funding_source,omitempty"`
	DropOffLocation      *DropOffLocation    `json:"drop_off_location,omitempty" bson:"drop_off_location,omitempty"`
	Manifest             *Manifest           `json:"manifest,omitempty" bson:"manifest,omitempty"`
	Driver               *Courier            `json:"driver" bson:"driver"`
	ProofOfImages        *[]string           `json:"proof_of_images,omitempty" bson:"proof_of_images,omitempty"`
	CreatedAt            *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt            *int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt            *int64              `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type FundingSource struct {
	UserWalletId primitive.ObjectID `json:"user_wallet_id,omitempty" bson:"user_wallet_id,omitempty" validate:"required"`
	Source       string             `json:"source,omitempty" bson:"source,omitempty" validate:"required,min=3"`
	Amount       int64              `json:"amount,omitempty" bson:"amount,omitempty" validate:"required,gt=0"`
	ProofImage   string             `json:"proof_image,omitempty" bson:"proof_image,omitempty"`
}

type Manifest struct {
	Merchant                Merchant       `json:"merchant,omitempty" bson:"merchant,omitempty"`
	Products                []ProductOrder `json:"products,omitempty" bson:"products,omitempty"`
	ProductsAmount          int64          `json:"products_amount,omitempty" bson:"products_amount,omitempty"`
	BasketPrice             int64          `json:"basket_price,omitempty" bson:"basket_price,omitempty"`
	AdminFee                int64          `json:"admin_fee,omitempty" bson:"admin_fee,omitempty"`
	PICFee                  int64          `json:"pic_fee,omitempty" bson:"pic_fee,omitempty"`
	DeliveryFee             int64          `json:"delivery_fee,omitempty" bson:"delivery_fee,omitempty"`
	TotalPrice              int64          `json:"total_price,omitempty" bson:"total_price,omitempty"`
	PromoAmount             int64          `json:"promo_amount,omitempty" bson:"promo_amount,omitempty"`
	RemainingFund           int64          `json:"remaining_fund,omitempty" bson:"remaining_fund,omitempty"`
	RemainingFundPercentage float64        `json:"remaining_fund_percentage,omitempty" bson:"remaining_fund_percentage,omitempty"`
}

type ProductOrder struct {
	Product  Product `json:"product,omitempty" bson:"product,omitempty"`
	Quantity int64   `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

type OrderStoreRequest struct {
	UserId          primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	FundingSource   FundingSource      `json:"funding_source,omitempty" bson:"funding_source,omitempty"`
	DropOffLocation DropOffLocation    `json:"drop_off_location,omitempty" bson:"drop_off_location,omitempty"`
	Manifest        Manifest           `json:"manifest,omitempty" bson:"manifest,omitempty"`
}

type OrderUpdateRequest struct {
	Id                   primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Transaction          *Transaction        `json:"transaction,omitempty" bson:"transaction,omitempty"`
	UserId               *primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	MerchantOrderId      *string             `json:"merchant_order_id,omitempty" bson:"merchant_order_id,omitempty"`
	DeliveryId           *string             `json:"delivery_id,omitempty" bson:"delivery_id,omitempty"`
	DeliveryStatus       *string             `json:"delivery_status,omitempty" bson:"delivery_status,omitempty"`
	DeliveryFailedReason *string             `json:"delivery_failed_reason,omitempty" bson:"delivery_failed_reason,omitempty"`
	FundingSource        *FundingSource      `json:"funding_source,omitempty" bson:"funding_source,omitempty"`
	DropOffLocation      *DropOffLocation    `json:"drop_off_location,omitempty" bson:"drop_off_location,omitempty"`
	Manifest             *Manifest           `json:"manifest,omitempty" bson:"manifest,omitempty"`
	Driver               *Courier            `json:"driver,omitempty" bson:"driver,omitempty"`
	ProofOfImages        *[]string           `json:"proof_of_images,omitempty" bson:"proof_of_images,omitempty"`
	CreatedAt            *int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt            *int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt            *int64              `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type ProofOfImageAddRequest struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ProofOfImage string             `json:"proof_of_image,omitempty" bson:"proof_of_image,omitempty"`
}

type OrderUsecase interface {
	First(ctx context.Context, orderFilter *Order) (Order, error)
	FetchByUserId(ctx context.Context, orderFilter *Order) ([]Order, error)
	// CreateDeliveryRequest(ctx context.Context, order *Order, transaction *Transaction) error
	Fetch(ctx context.Context) (res []Order, err error)
	Store(context.Context, *OrderStoreRequest) (*Order, error)
	Update(context.Context, *Order, *OrderUpdateRequest) error
	Delete(ctx context.Context, order *Order) error
}

type OrderRepository interface {
	First(ctx context.Context, orderFilter *Order) (Order, error)
	FetchByUserId(ctx context.Context, orderFilter *Order) ([]Order, error)
	Fetch(ctx context.Context) ([]Order, error)
	Store(ctx context.Context, order *Order) error
	Update(context.Context, *Order) error
	Delete(ctx context.Context, order *Order) error
}
