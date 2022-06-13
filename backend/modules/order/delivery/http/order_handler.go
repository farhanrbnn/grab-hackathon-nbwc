package http

// Need to use echo jwt version
import (
	"fmt"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"grab-hack-for-good/domain"
	"grab-hack-for-good/helper"
)

// OrderHandler  represent the httphandler for order
type OrderHandler struct {
	OrderUsecase           domain.OrderUsecase
	GrabExpressUsecase     domain.GrabExpressUsecase
	TransactionUsecase     domain.TransactionUsecase
	UserWalletUsecase      domain.UserWalletUsecase
	DropOffLocationUsecase domain.DropOffLocationUsecase
}

func NewOrderHandler(e *echo.Echo, r *echo.Group, cs domain.OrderUsecase, ge domain.GrabExpressUsecase, tu domain.TransactionUsecase, uwu domain.UserWalletUsecase, dolu domain.DropOffLocationUsecase) {
	handler := &OrderHandler{
		OrderUsecase:           cs,
		GrabExpressUsecase:     ge,
		TransactionUsecase:     tu,
		UserWalletUsecase:      uwu,
		DropOffLocationUsecase: dolu,
	}

	// r.GET("/api/v1/orders", handler.Fetch) // Testing, Admin only
	r.GET("/api/v1/orders", handler.FetchByUserId) // Testing, Admin only
	r.GET("/api/v1/order/:id", handler.FindByUrlId)
	r.DELETE("/api/v1/order/cancel-delivery/:id", handler.CancelDelivery)
	e.POST("/api/v1/order/update-status", handler.UpdateStatus)
	e.POST("/api/v1/order/add-proof-of-image", handler.AddProofOfImages)
	r.POST("/api/v1/order/create", handler.Store)
	r.PATCH("/api/v1/order/update", handler.Update)
	r.DELETE("/api/v1/order/delete", handler.Delete)
}

func (handler *OrderHandler) FindByUrlId(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return handler.FindByUrl(c, &domain.Order{Id: id})
}

func (handler *OrderHandler) FindByUrl(c echo.Context, orderId *domain.Order) (err error) {
	ctx := c.Request().Context()

	order, err := handler.OrderUsecase.First(ctx, orderId)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, order)
}

func (handler *OrderHandler) FetchByUserId(c echo.Context) error {
	ctx := c.Request().Context()

	userId, err := helper.GetUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	res, err := handler.OrderUsecase.FetchByUserId(ctx, &domain.Order{
		UserId: &userId,
	})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *OrderHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := handler.OrderUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *OrderHandler) Get(c echo.Context) error {
	user := helper.GetUserByToken(c)

	return c.JSON(http.StatusOK, user)
}

func (handler *OrderHandler) UpdateStatus(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var grabExpressDeliveryResponse domain.GrabExpressDeliveryResponse
	err = c.Bind(&grabExpressDeliveryResponse)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	orderId, err := primitive.ObjectIDFromHex(grabExpressDeliveryResponse.MerchantOrderId)
	order, err := handler.OrderUsecase.First(ctx, &domain.Order{
		Id: orderId,
	})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	var orderUpdateRequest domain.OrderUpdateRequest
	copier.Copy(&orderUpdateRequest, &order)

	deliveryStatus := grabExpressDeliveryResponse.Status
	deliveryFailedReason := grabExpressDeliveryResponse.FailedReason

	orderUpdateRequest.DeliveryStatus = &deliveryStatus
	orderUpdateRequest.DeliveryFailedReason = &deliveryFailedReason

	err = handler.OrderUsecase.Update(ctx, &order, &orderUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	userWallet, err := handler.UserWalletUsecase.First(ctx, &domain.UserWallet{
		Id: order.Transaction.UserWalletId,
	})

	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if deliveryStatus == "RETURNED" || deliveryStatus == "SENDER_CANCELLED" || deliveryStatus == "SCHEDULE_FAILED" || deliveryStatus == "DRIVER_CANCELED" || deliveryStatus == "OPERATOR_CANCELLED" || deliveryStatus == "FAILED" {
		dropOffLocation, err := handler.DropOffLocationUsecase.First(ctx, &domain.DropOffLocation{
			Id: order.DropOffLocation.Id,
		})
		if err != nil {
			return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
		}

		var dropOffLocationUpdateRequest domain.DropOffLocationUpdateRequest
		copier.Copy(&dropOffLocationUpdateRequest, &dropOffLocationUpdateRequest)
		dropOffLocation.CurrentSupply = dropOffLocation.CurrentSupply - order.Manifest.ProductsAmount

		err = handler.DropOffLocationUsecase.Update(ctx, &dropOffLocation, &dropOffLocationUpdateRequest)
		if err != nil {
			return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
		}

		err = handler.UserWalletUsecase.Update(ctx, &userWallet, &domain.UserWalletUpdateRequest{
			Id:              userWallet.Id,
			EffectiveAmount: userWallet.EffectiveAmount + order.Transaction.Amount,
			OnHoldAmount:    userWallet.OnHoldAmount - order.Transaction.Amount,
		})
		if err != nil {
			return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
		}
	} else if deliveryStatus == "COMPLETED" {
		err = handler.UserWalletUsecase.Update(ctx, &userWallet, &domain.UserWalletUpdateRequest{
			Id:              userWallet.Id,
			EffectiveAmount: userWallet.EffectiveAmount,
			OnHoldAmount:    userWallet.OnHoldAmount - order.Transaction.Amount,
		})
		if err != nil {
			return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
		}
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusUpdatedMessage)
}

func (handler *OrderHandler) CancelDelivery(c echo.Context) (err error) {
	ctx := c.Request().Context()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	order, err := handler.OrderUsecase.First(ctx, &domain.Order{
		Id: id,
	})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	userId, err := helper.GetUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	userWallet, err := handler.UserWalletUsecase.First(ctx, &domain.UserWallet{
		Id: order.Transaction.UserWalletId,
	})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	dropOffLocation, err := handler.DropOffLocationUsecase.First(ctx, &domain.DropOffLocation{
		Id: order.DropOffLocation.Id,
	})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if userId != *order.UserId || userId != userWallet.UserId {
		err = helper.ErrUnauthorized
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}
	err = handler.GrabExpressUsecase.CancelDelivery(ctx, order)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	var orderUpdateRequest domain.OrderUpdateRequest
	copier.Copy(&orderUpdateRequest, &order)

	deliveryStatus := "FAILED"
	deliveryFailedReason := "Canceled by User"

	orderUpdateRequest.DeliveryStatus = &deliveryStatus
	orderUpdateRequest.DeliveryFailedReason = &deliveryFailedReason

	err = handler.OrderUsecase.Update(ctx, &order, &orderUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	var dropOffLocationUpdateRequest domain.DropOffLocationUpdateRequest
	copier.Copy(&dropOffLocationUpdateRequest, &dropOffLocationUpdateRequest)
	dropOffLocation.CurrentSupply = dropOffLocation.CurrentSupply - order.Manifest.ProductsAmount

	err = handler.DropOffLocationUsecase.Update(ctx, &dropOffLocation, &dropOffLocationUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.UserWalletUsecase.Update(ctx, &userWallet, &domain.UserWalletUpdateRequest{
		Id:              userWallet.Id,
		EffectiveAmount: userWallet.EffectiveAmount + order.Transaction.Amount,
		OnHoldAmount:    userWallet.OnHoldAmount - order.Transaction.Amount,
	})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, order)
}

func (handler *OrderHandler) Store(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var orderStoreRequest domain.OrderStoreRequest
	err = c.Bind(&orderStoreRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	orderStoreRequest.UserId, err = helper.GetUserId(c)
	if err != nil {
		return
	}

	if err = c.Validate(orderStoreRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var grabExpressDeliveryQuotesRequest domain.GrabExpressDeliveryQuotesRequest
	grabExpressDeliveryQuotesRequest.ServiceType = "INSTANT"

	var _package domain.Package
	var amount int64
	amount = 0
	// var _dimensions domain.Dimensions
	for i := 0; i < len(orderStoreRequest.Manifest.Products); i++ {
		_package.Name = orderStoreRequest.Manifest.Products[i].Product.Name
		_package.Description = orderStoreRequest.Manifest.Products[i].Product.Description
		_package.Price = orderStoreRequest.Manifest.Products[i].Product.Price
		_package.Quantity = orderStoreRequest.Manifest.Products[i].Quantity
		_package.Dimensions = orderStoreRequest.Manifest.Products[i].Product.Dimensions

		grabExpressDeliveryQuotesRequest.Packages = append(grabExpressDeliveryQuotesRequest.Packages, _package)
		amount += _package.Quantity
	}
	orderStoreRequest.Manifest.ProductsAmount = amount

	grabExpressDeliveryQuotesRequest.Origin = domain.Origin{
		Address:     orderStoreRequest.Manifest.Merchant.Address,
		Coordinates: orderStoreRequest.Manifest.Merchant.Coordinate,
	}

	grabExpressDeliveryQuotesRequest.Destination = domain.Destination{
		Address:     orderStoreRequest.DropOffLocation.Address,
		Coordinates: orderStoreRequest.DropOffLocation.Coordinate,
	}

	grabExpressDeliveryQuotesResponse, err := handler.GrabExpressUsecase.GetDeliveryQuotes(ctx, &grabExpressDeliveryQuotesRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	orderStoreRequest.Manifest.DeliveryFee = int64(grabExpressDeliveryQuotesResponse.Quotes[0].Amount)

	order, err := handler.OrderUsecase.Store(ctx, &orderStoreRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	dropOffLocation, err := handler.DropOffLocationUsecase.First(ctx, &domain.DropOffLocation{
		Id: order.DropOffLocation.Id,
	})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	var dropOffLocationUpdateRequest domain.DropOffLocationUpdateRequest
	copier.Copy(&dropOffLocationUpdateRequest, &dropOffLocation)
	dropOffLocationUpdateRequest.CurrentSupply = dropOffLocation.CurrentSupply + amount

	err = handler.DropOffLocationUsecase.Update(ctx, &dropOffLocation, &dropOffLocationUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	transactionStoreRequest := domain.TransactionStoreRequest{
		OrderId:      order.Id,
		UserId:       orderStoreRequest.UserId,
		UserWalletId: orderStoreRequest.FundingSource.UserWalletId,
		Status:       "Waiting",
		Amount:       order.Manifest.TotalPrice,
	}

	transaction, err := handler.TransactionUsecase.Store(ctx, &transactionStoreRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, struct {
		Order       *domain.Order       `json:"order,omitempty"`
		Transaction *domain.Transaction `json:"transaction,omitempty"`
	}{
		order,
		transaction,
	})
}

func (handler *OrderHandler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var orderUpdateRequest domain.OrderUpdateRequest
	err = c.Bind(&orderUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(orderUpdateRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	fmt.Print(orderUpdateRequest.Id)

	order, err := handler.OrderUsecase.First(ctx, &domain.Order{Id: orderUpdateRequest.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.OrderUsecase.Update(ctx, &order, &orderUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusUpdatedMessage)
}

func (handler *OrderHandler) AddProofOfImages(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var proofOfImageAddRequest domain.ProofOfImageAddRequest
	err = c.Bind(&proofOfImageAddRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(proofOfImageAddRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	order, err := handler.OrderUsecase.First(ctx, &domain.Order{Id: proofOfImageAddRequest.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if *order.DeliveryStatus != "COMPLETED" {
		err = helper.ErrUnauthorized
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	var orderUpdateRequest domain.OrderUpdateRequest
	copier.Copy(&orderUpdateRequest, &order)
	fmt.Println(orderUpdateRequest)
	fmt.Println(proofOfImageAddRequest)

	var proofOfImages []string
	if orderUpdateRequest.ProofOfImages == nil {
		proofOfImages = append(proofOfImages, proofOfImageAddRequest.ProofOfImage)
	} else {
		proofOfImages = append(*orderUpdateRequest.ProofOfImages, proofOfImageAddRequest.ProofOfImage)
	}
	fmt.Println(proofOfImages)
	orderUpdateRequest.ProofOfImages = &proofOfImages

	err = handler.OrderUsecase.Update(ctx, &order, &orderUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusUpdatedMessage)
}

// Delete will remove the Order by given id
func (handler *OrderHandler) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var orderData domain.Order

	err = c.Bind(&orderData)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	order, err := handler.OrderUsecase.First(ctx, &domain.Order{Id: orderData.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.OrderUsecase.Delete(ctx, &order)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusDeletedMessage)
}
