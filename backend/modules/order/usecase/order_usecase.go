package usecase

import (
	"context"
	"time"

	"grab-hack-for-good/domain"

	"github.com/spf13/viper"
)

type orderUsecase struct {
	orderRepo      domain.OrderRepository
	contextTimeout time.Duration
}

// NewOrderUsecase will create new an orderUsecase object representation of domain.OrderUsecase interface
func NewOrderUsecase(cr domain.OrderRepository, timeout time.Duration) domain.OrderUsecase {
	return &orderUsecase{
		orderRepo:      cr,
		contextTimeout: timeout,
	}
}

func (usecase *orderUsecase) First(c context.Context, orderFilter *domain.Order) (order domain.Order, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	order, err = usecase.orderRepo.First(ctx, orderFilter)
	if err != nil {
		return domain.Order{}, err
	}

	return
}

func (usecase *orderUsecase) FetchByUserId(c context.Context, orderFilter *domain.Order) (res []domain.Order, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.orderRepo.FetchByUserId(ctx, orderFilter)
	if err != nil {
		return make([]domain.Order, 0), err
	}

	return
}

func (usecase *orderUsecase) Fetch(c context.Context) (res []domain.Order, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.orderRepo.Fetch(ctx)
	if err != nil {
		return make([]domain.Order, 0), err
	}

	return
}

// func (usecase *orderUsecase) CreateDeliveryRequest(c context.Context, order *domain.Order, transaction *domain.Transaction) (err error) {
// 	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
// 	defer cancel()

// 	current_time := time.Now()
// 	grabExpressDeliveryRequest := domain.GrabExpressDeliveryRequest{
// 		MerchantOrderId: order.Id.Hex(),
// 		ServiceType:     "INSTANT",
// 		PaymentMethod:   "CASHLESS",
// 	}

// 	var _package domain.Package
// 	// var _dimensions domain.Dimensions
// 	for i := 0; i < len(order.Manifest.Products); i++ {
// 		_package.Name = order.Manifest.Products[i].Product.Name
// 		_package.Description = order.Manifest.Products[i].Product.Description
// 		_package.Price = order.Manifest.Products[i].Product.Price
// 		_package.Quantity = order.Manifest.Products[i].Quantity
// 		_package.Dimensions = order.Manifest.Products[i].Product.Dimensions
// 		// _package.Dimensions = domain.Dimensions{
// 		// 	Height: 0,
// 		// 	Width:  0,
// 		// 	Depth:  0,
// 		// 	Weight: 0,
// 		// }
// 		_package.Dimensions = order.Manifest.Products[i].Product.Dimensions

// 		grabExpressDeliveryRequest.Packages = append(grabExpressDeliveryRequest.Packages, _package)
// 	}

// 	grabExpressDeliveryRequest.Origin = domain.Origin{
// 		Address:     order.Manifest.Merchant.Address,
// 		Coordinates: order.Manifest.Merchant.Coordinate,
// 	}

// 	grabExpressDeliveryRequest.Destination = domain.Destination{
// 		Address:     order.DropOffLocation.Address,
// 		Coordinates: order.DropOffLocation.Coordinate,
// 	}
// 	grabExpressDeliveryRequest.Sender = domain.Sender{
// 		FirstName:   order.Manifest.Merchant.Name,
// 		CompanyName: order.Manifest.Merchant.Name,
// 		Email:       order.Manifest.Merchant.Email,
// 		Phone:       order.Manifest.Merchant.Phone,
// 		SmsEnabled:  true,
// 	}

// 	grabExpressDeliveryRequest.Recipient = domain.Recipient{
// 		FirstName:   order.DropOffLocation.PIC,
// 		CompanyName: order.DropOffLocation.Name,
// 		Phone:       order.DropOffLocation.Phone,
// 		Email:       order.DropOffLocation.Email,
// 		SmsEnabled:  true,
// 	}

// 	grabExpressDeliveryRequest.Schedule = domain.Schedule{
// 		PickupTimeFrom: string(current_time.Add(time.Hour * 0).Format(time.RFC3339)),
// 		PickupTimeTo:   string(current_time.Add(time.Hour * 2).Format(time.RFC3339)),
// 	}

// 	sqsRequest := domain.SQSGrabExpressDeliveryRequest{
// 		Payload:            grabExpressDeliveryRequest,
// 		GrabEndpointURL:    viper.GetString("aws.grab_endpoint_url"),
// 		BackendEndpointURL: viper.GetString("aws.backend_endpoint_url"),
// 	}

// 	queueUrl := viper.GetString("aws.grab_create_delivery_queue_url")
// 	messageBody, err := json.Marshal(sqsRequest)
// 	if err != nil {
// 		return
// 	}

// 	sess := session.Must(session.NewSessionWithOptions(session.Options{
// 		SharedConfigState: session.SharedConfigEnable,
// 	}))
// 	svc := sqs.New(sess)

// 	_, err = svc.SendMessage(&sqs.SendMessageInput{
// 		DelaySeconds: aws.Int64(10),
// 		MessageAttributes: map[string]*sqs.MessageAttributeValue{
// 			"Title": &sqs.MessageAttributeValue{
// 				DataType:    aws.String("String"),
// 				StringValue: aws.String("The Whistler"),
// 			},
// 			"Author": &sqs.MessageAttributeValue{
// 				DataType:    aws.String("String"),
// 				StringValue: aws.String("John Grisham"),
// 			},
// 			"WeeksOn": &sqs.MessageAttributeValue{
// 				DataType:    aws.String("Number"),
// 				StringValue: aws.String("6"),
// 			},
// 		},
// 		MessageBody: aws.String(string(messageBody)),
// 		QueueUrl:    &queueUrl,
// 	})

// 	return

// 	// grabExpressDeliveryResponse, err := handler.GrabExpressUsecase.CreateDeliveryRequest(ctx, &grabExpressDeliveryRequest)
// 	// if err != nil {
// 	// 	return
// 	// }

// 	// var orderUpdateRequest domain.OrderUpdateRequest
// 	// copier.Copy(&orderUpdateRequest, &order)
// 	// orderUpdateRequest.Transaction = transaction
// 	// orderUpdateRequest.DeliveryId = &grabExpressDeliveryResponse.DeliveryId
// 	// orderUpdateRequest.MerchantOrderId = &grabExpressDeliveryResponse.MerchantOrderId

// 	// err = usecase.Update(ctx, order, &orderUpdateRequest)
// 	// if err != nil {
// 	// 	return
// 	// }
// }

func (usecase *orderUsecase) Store(c context.Context, orderStoreRequest *domain.OrderStoreRequest) (order *domain.Order, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	var basketPrice int64
	basketPrice = 0
	for i := 0; i < len(orderStoreRequest.Manifest.Products); i++ {
		basketPrice += orderStoreRequest.Manifest.Products[i].Product.Price * orderStoreRequest.Manifest.Products[i].Quantity
	}

	orderStoreRequest.Manifest.BasketPrice = basketPrice
	orderStoreRequest.Manifest.PICFee = viper.GetInt64("pricing.pic_fee")

	subtotal := orderStoreRequest.Manifest.BasketPrice + orderStoreRequest.Manifest.DeliveryFee + orderStoreRequest.Manifest.PICFee

	orderStoreRequest.Manifest.AdminFee = viper.GetInt64("pricing.admin_fee_percentage") * subtotal / 100
	orderStoreRequest.Manifest.TotalPrice = subtotal + orderStoreRequest.Manifest.AdminFee
	orderStoreRequest.Manifest.RemainingFund = orderStoreRequest.FundingSource.Amount - orderStoreRequest.Manifest.TotalPrice
	orderStoreRequest.Manifest.RemainingFundPercentage = (float64(orderStoreRequest.Manifest.RemainingFund) / float64(orderStoreRequest.Manifest.TotalPrice)) * 100

	order = &domain.Order{
		UserId:          &orderStoreRequest.UserId,
		FundingSource:   &orderStoreRequest.FundingSource,
		DropOffLocation: &orderStoreRequest.DropOffLocation,
		Manifest:        &orderStoreRequest.Manifest,
	}

	err = usecase.orderRepo.Store(ctx, order)

	return
}

func (usecase *orderUsecase) Update(c context.Context, order *domain.Order, orderUpdateRequest *domain.OrderUpdateRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	order.Transaction = orderUpdateRequest.Transaction
	order.UserId = orderUpdateRequest.UserId
	order.MerchantOrderId = orderUpdateRequest.MerchantOrderId
	order.DeliveryId = orderUpdateRequest.DeliveryId
	order.DeliveryStatus = orderUpdateRequest.DeliveryStatus
	order.DeliveryFailedReason = orderUpdateRequest.DeliveryFailedReason
	order.FundingSource = orderUpdateRequest.FundingSource
	order.DropOffLocation = orderUpdateRequest.DropOffLocation
	order.Manifest = orderUpdateRequest.Manifest
	order.Driver = orderUpdateRequest.Driver
	order.ProofOfImages = orderUpdateRequest.ProofOfImages

	return usecase.orderRepo.Update(ctx, order)
}

func (usecase *orderUsecase) Delete(c context.Context, order *domain.Order) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	err = usecase.orderRepo.Delete(ctx, order)

	return
}
