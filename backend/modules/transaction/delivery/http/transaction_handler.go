package http

// Need to use echo jwt version
import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"grab-hack-for-good/domain"
	"grab-hack-for-good/helper"
)

// TransactionHandler  represent the httphandler for transaction
type TransactionHandler struct {
	TransactionUsecase domain.TransactionUsecase
	GrabExpressUsecase domain.GrabExpressUsecase
	OrderUsecase       domain.OrderUsecase
	UserWalletUsecase  domain.UserWalletUsecase
}

func NewTransactionHandler(e *echo.Echo, r *echo.Group, cs domain.TransactionUsecase, ge domain.GrabExpressUsecase, ou domain.OrderUsecase, uwu domain.UserWalletUsecase) {
	handler := &TransactionHandler{
		TransactionUsecase: cs,
		GrabExpressUsecase: ge,
		OrderUsecase:       ou,
		UserWalletUsecase:  uwu,
	}

	r.GET("/api/v1/transaction/pay/:id", handler.Pay)

	// r.GET("/api/v1/transactions", handler.Fetch) // Testing, Admin only
	r.GET("/api/v1/transactions", handler.FetchByUserId)
	r.GET("/api/v1/transaction/:id", handler.FindByUrlId)
	r.POST("/api/v1/transaction/create", handler.Store)
	r.PATCH("/api/v1/transaction/update", handler.Update)
	r.DELETE("/api/v1/transaction/delete", handler.Delete)
}

func (handler *TransactionHandler) Pay(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	userId, err := helper.GetUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()

	transaction, err := handler.TransactionUsecase.First(ctx, &domain.Transaction{
		Id: id,
	})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	userWallet, err := handler.UserWalletUsecase.First(ctx, &domain.UserWallet{
		Id: transaction.UserWalletId,
	})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if userId != transaction.UserId || userId != userWallet.UserId {
		err = helper.ErrUnauthorized
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if userWallet.EffectiveAmount < transaction.Amount {
		err = helper.ErrUnauthorized
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if transaction.Status == "Paid" {
		err = echo.ErrBadRequest
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	transactionUpdateRequest := domain.TransactionUpdateRequest{
		Status: "Paid",
	}

	err = handler.TransactionUsecase.Update(ctx, &transaction, &transactionUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	order, err := handler.OrderUsecase.First(ctx, &domain.Order{
		Id: transaction.OrderId,
	})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.GrabExpressUsecase.QueueCreateDeliveryRequest(ctx, &order, &transaction)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.UserWalletUsecase.Update(ctx, &userWallet, &domain.UserWalletUpdateRequest{
		Id:              userWallet.Id,
		EffectiveAmount: userWallet.EffectiveAmount - transaction.Amount,
		OnHoldAmount:    userWallet.OnHoldAmount + transaction.Amount,
	})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	var orderUpdateRequest domain.OrderUpdateRequest
	copier.Copy(&orderUpdateRequest, &order)
	orderUpdateRequest.Transaction = &transaction

	err = handler.OrderUsecase.Update(ctx, &order, &orderUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, transaction)
	// return c.JSON(http.StatusOK, struct {
	// 	Transaction *domain.Transaction `json:"transaction,omitempty"`
	// 	// GrabExpressDeliveryResponse *domain.GrabExpressDeliveryResponse `json:"grab_express_delivery_response,omitempty"`
	// }{
	// 	&transaction,
	// 	// &grabExpressDeliveryResponse,
	// })
}

func (handler *TransactionHandler) FindByUrlId(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return handler.FindByUrl(c, &domain.Transaction{Id: id})
}

func (handler *TransactionHandler) FindByUrl(c echo.Context, transactionId *domain.Transaction) (err error) {
	ctx := c.Request().Context()

	transaction, err := handler.TransactionUsecase.First(ctx, transactionId)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, transaction)
}

func (handler *TransactionHandler) FetchByUserId(c echo.Context) error {
	ctx := c.Request().Context()

	userId, err := helper.GetUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	res, err := handler.TransactionUsecase.FetchByUserId(ctx, &domain.Transaction{
		UserId: userId,
	})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *TransactionHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := handler.TransactionUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *TransactionHandler) Store(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var transactionStoreRequest domain.TransactionStoreRequest
	err = c.Bind(&transactionStoreRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err = c.Validate(transactionStoreRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	_, err = handler.TransactionUsecase.Store(ctx, &transactionStoreRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, helper.HttpStatusCreatedMessage)
}

func (handler *TransactionHandler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var transactionUpdateRequest domain.TransactionUpdateRequest
	err = c.Bind(&transactionUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(transactionUpdateRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	transaction, err := handler.TransactionUsecase.First(ctx, &domain.Transaction{Id: transactionUpdateRequest.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.TransactionUsecase.Update(ctx, &transaction, &transactionUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusUpdatedMessage)
}

// Delete will remove the Transaction by given id
func (handler *TransactionHandler) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var transactionData domain.Transaction

	err = c.Bind(&transactionData)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	transaction, err := handler.TransactionUsecase.First(ctx, &domain.Transaction{Id: transactionData.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.TransactionUsecase.Delete(ctx, &transaction)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusDeletedMessage)
}
