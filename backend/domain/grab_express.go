package domain

import (
	"context"
)

type GrabExpressCredential struct {
	ClientId     string `json:"client_id,omitempty" bson:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty" bson:"client_secret,omitempty"`
	GrantType    string `json:"grant_type,omitempty" bson:"grant_type,omitempty"`
	Scope        string `json:"scope,omitempty" bson:"scope,omitempty"`
}

type GrabExpressToken struct {
	AccessToken string `json:"access_token,omitempty" bson:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty" bson:"token_type,omitempty"`
	ExpiresIn   int64  `json:"expires_in,omitempty" bson:"expires_in,omitempty"`
}

type Package struct {
	Name        string     `json:"name,omitempty" bson:"name,omitempty"`
	Description string     `json:"description,omitempty" bson:"description,omitempty"`
	Price       int64      `json:"price,omitempty" bson:"price,omitempty"`
	Quantity    int64      `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Dimensions  Dimensions `json:"dimensions,omitempty" bson:"dimensions,omitempty"`
}

type Origin struct {
	Address     string     `json:"address,omitempty" bson:"address,omitempty" validate:"required,min=8,max=500"`
	Coordinates Coordinate `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
}

type Destination struct {
	Address     string     `json:"address,omitempty" bson:"address,omitempty" validate:"required,min=8,max=500"`
	Coordinates Coordinate `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
}

type Recipient struct {
	FirstName   string `json:"firstName,omitempty" bson:"firstName,omitempty" validate:"required,min=3,max=100"`
	LastName    string `json:"lastName,omitempty" bson:"lastName,omitempty" validate:"required,min=3,max=100"`
	Title       string `json:"title,omitempty" bson:"title,omitempty"`
	CompanyName string `json:"companyName,omitempty" bson:"companyName,omitempty"`
	Phone       string `json:"phone,omitempty" bson:"phone,omitempty" validate:"required,min=8,max=16"`
	Email       string `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	SmsEnabled  bool   `json:"smsEnabled,omitempty" bson:"smsEnabled,omitempty" validate:"required,email"`
}

type Sender struct {
	FirstName   string `json:"firstName,omitempty" bson:"firstName,omitempty" validate:"required,min=3,max=100"`
	LastName    string `json:"lastName,omitempty" bson:"lastName,omitempty" validate:"min=3,max=100"`
	Title       string `json:"title,omitempty" bson:"title,omitempty"`
	CompanyName string `json:"companyName,omitempty" bson:"companyName,omitempty" validate:"required,min=8,max=100"`
	Phone       string `json:"phone,omitempty" bson:"phone,omitempty" validate:"required,min=8,max=16"`
	Email       string `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	SmsEnabled  bool   `json:"smsEnabled,omitempty" bson:"smsEnabled,omitempty" validate:"required,email"`
	Instruction string `json:"instruction,omitempty" bson:"instruction,omitempty" validate:"required,max=1000"`
}

type Schedule struct {
	PickupTimeFrom string `json:"pickupTimeFrom,omitempty" bson:"pickupTimeFrom,omitempty"`
	PickupTimeTo   string `json:"pickupTimeTo,omitempty" bson:"pickupTimeTo,omitempty"`
}

type Vehicle struct {
	LicensePlate        string `json:"licensePlate,omitempty" bson:"licensePlate,omitempty"`
	Model               string `json:"model,omitempty" bson:"model,omitempty"`
	PhysicalVehicleType string `json:"physicalVehicleType,omitempty" bson:"physicalVehicleType,omitempty"`
}

type Courier struct {
	Name        string     `json:"name,omitempty" bson:"name,omitempty"`
	Phone       string     `json:"phone,omitempty" bson:"phone,omitempty"`
	PictureURL  string     `json:"pictureURL,omitempty" bson:"pictureURL,omitempty"`
	Rating      float64    `json:"rating,omitempty" bson:"rating,omitempty"`
	Coordinates Coordinate `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
}

type Service struct {
	Id   int64  `json:"id,omitempty" bson:"id,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

type Currency struct {
	Code     string `json:"code,omitempty" bson:"code,omitempty"`
	Symbol   string `json:"symbol,omitempty" bson:"symbol,omitempty"`
	Exponent int64  `json:"exponent,omitempty" bson:"exponent,omitempty"`
}

type EstimatedTimeline struct {
	Pickup  string `json:"pickup,omitempty" bson:"pickup,omitempty"`
	Dropoff string `json:"dropoff,omitempty" bson:"dropoff,omitempty"`
}

type DiscountInfo struct {
	Amount   float64 `json:"amount,omitempty" bson:"amount,omitempty"`
	Success  bool    `json:"success,omitempty" bson:"success,omitempty"`
	ErrorMsg string  `json:"errorMsg,omitempty" bson:"errorMsg,omitempty"`
}

type Quote struct {
	Service           Service           `json:"service,omitempty" bson:"service,omitempty"`
	Currency          Currency          `json:"currency,omitempty" bson:"currency,omitempty"`
	Amount            float64           `json:"amount,omitempty" bson:"amount,omitempty"`
	EstimatedTimeline EstimatedTimeline `json:"estimatedTimeline,omitempty" bson:"estimatedTimeline,omitempty"`
	Distance          int64             `json:"distance,omitempty" bson:"distance,omitempty"`
	Packages          []Package         `json:"packages,omitempty" bson:"packages,omitempty"`
	Origin            Origin            `json:"origin,omitempty" bson:"origin,omitempty"`
	Destination       Destination       `json:"destination,omitempty" bson:"destination,omitempty"`
	DiscountInfo      DiscountInfo      `json:"discountInfo,omitempty" bson:"discountInfo,omitempty"`
}

type GrabExpressDeliveryQuotesRequest struct {
	ServiceType string      `json:"serviceType,omitempty" bson:"serviceType,omitempty" validate:"required,min=3,max=500"`
	Packages    []Package   `json:"packages,omitempty" bson:"packages,omitempty"`
	Origin      Origin      `json:"origin,omitempty" bson:"origin,omitempty"`
	Destination Destination `json:"destination,omitempty" bson:"destination,omitempty"`
}

type GrabExpressDeliveryQuotesResponse struct {
	Quotes      []Quote     `json:"quotes,omitempty" bson:"quotes,omitempty"`
	Packages    []Package   `json:"packages,omitempty" bson:"packages,omitempty"`
	Origin      Origin      `json:"origin,omitempty" bson:"origin,omitempty"`
	Destination Destination `json:"destination,omitempty" bson:"destination,omitempty"`
}

type GrabExpressDeliveryRequest struct {
	MerchantOrderId string      `json:"merchantOrderID,omitempty" bson:"merchantOrderID,omitempty" validate:"required,min=8,max=500"`
	ServiceType     string      `json:"serviceType,omitempty" bson:"serviceType,omitempty" validate:"required,min=8,max=500"`
	PaymentMethod   string      `json:"paymentMethod,omitempty" bson:"paymentMethod,omitempty" validate:"required,min=8,max=500"`
	Packages        []Package   `json:"packages,omitempty" bson:"packages,omitempty"`
	Origin          Origin      `json:"origin,omitempty" bson:"origin,omitempty"`
	Destination     Destination `json:"destination,omitempty" bson:"destination,omitempty"`
	Recipient       Recipient   `json:"recipient,omitempty" bson:"recipient,omitempty"`
	Sender          Sender      `json:"sender,omitempty" bson:"sender,omitempty"`
	Schedule        Schedule    `json:"schedule,omitempty" bson:"schedule,omitempty"`
}

type GrabExpressDeliveryResponse struct {
	DeliveryId      string                 `json:"deliveryID,omitempty" bson:"delivery_id,omitempty" validate:"required,min=8,max=500"`
	MerchantOrderId string                 `json:"merchantOrderID,omitempty" bson:"merchantOrderID,omitempty" validate:"required,min=8,max=500"`
	Payer           string                 `json:"payer,omitempty" bson:"payer,omitempty"`
	Status          string                 `json:"status,omitempty" bson:"status,omitempty" validate:"required"`
	FailedReason    string                 `json:"failedReason,omitempty" bson:"failed_reason,omitempty" validate:"required"`
	TrackingURL     string                 `json:"trackingURL,omitempty" bson:"trackingURL,omitempty"`
	Courier         Courier                `json:"courier,omitempty" bson:"courier,omitempty"`
	Timeline        map[string]interface{} `json:"timeline,omitempty" bson:"timeline,omitempty"`
	Schedule        map[string]interface{} `json:"schedule,omitempty" bson:"schedule,omitempty"`
	CashOnDelivery  map[string]interface{} `json:"cashOnDelivery,omitempty" bson:"cashOnDelivery,omitempty"`
	InvoiceNo       string                 `json:"invoiceNo,omitempty" bson:"invoiceNo,omitempty"`
	PickupPin       string                 `json:"pickupPin,omitempty" bson:"pickupPin,omitempty"`
	AdvanceInfo     map[string]interface{} `json:"advanceInfo,omitempty" bson:"advanceInfo,omitempty"`
	Sender          Sender                 `json:"sender,omitempty" bson:"sender,omitempty"`
	Recipient       Recipient              `json:"recipient,omitempty" bson:"recipient,omitempty"`
	Quote           Quote                  `json:"quote,omitempty" bson:"quote,omitempty"`
}

type SQSGrabExpressDeliveryRequest struct {
	Payload            GrabExpressDeliveryRequest `json:"payload,omitempty"`
	GrabToken          string                     `json:"grab_credentials,omitempty"`
	GrabEndpointURL    string                     `json:"grab_endpoint_url,omitempty"`
	BackendEndpointURL string                     `json:"backend_endpoint_url,omitempty"`
	UserToken          string                     `json:"backend_credentials,omitempty"`
}

type GrabExpressUsecase interface {
	GetToken(ctx context.Context) error
	GetDeliveryQuotes(ctx context.Context, grabExpressDeliveryQuotesRequest *GrabExpressDeliveryQuotesRequest) (GrabExpressDeliveryQuotesResponse, error)
	CreateDeliveryRequest(ctx context.Context, grabExpressDeliveryRequest *GrabExpressDeliveryRequest) (GrabExpressDeliveryResponse, error)
	QueueCreateDeliveryRequest(ctx context.Context, order *Order, transaction *Transaction) error
	CancelDelivery(ctx context.Context, order Order) error
}

type GrabExpressRepository interface {
}

// type GrabExpress struct {
// 	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
// 	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
// 	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty"`
// 	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
// 	Coordinate  Coordinate         `json:"coordinate,omitempty" bson:"coordinate,omitempty"`
// 	Rating      float32            `json:"rating,omitempty" bson:"rating,omitempty"`
// 	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
// 	Thumbnail   string             `json:"thumbnail,omitempty" bson:"thumbnail,omitempty"`
// 	IsAvailable bool               `json:"is_available,omitempty" bson:"is_available,omitempty"`
// 	CreatedAt   int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
// 	UpdatedAt   int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
// 	DeletedAt   int64              `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
// }
