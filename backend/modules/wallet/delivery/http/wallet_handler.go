package http

// Need to use echo jwt version
import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"grab-hack-for-good/domain"
	"grab-hack-for-good/helper"
)

// WalletHandler  represent the httphandler for wallet
type WalletHandler struct {
	WalletUsecase domain.WalletUsecase
}

func NewWalletHandler(e *echo.Echo, r *echo.Group, cs domain.WalletUsecase) {
	handler := &WalletHandler{
		WalletUsecase: cs,
	}

	r.GET("/api/v1/wallets", handler.Fetch) // Testing, Admin only
	r.GET("/api/v1/wallet/:id", handler.FindByUrlId)
	r.POST("/api/v1/wallet/create", handler.Store)
	r.PATCH("/api/v1/wallet/update", handler.Update)
	r.DELETE("/api/v1/wallet/delete", handler.Delete)
}

func (handler *WalletHandler) FindByUrlId(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return handler.FindByUrl(c, &domain.Wallet{Id: id})
}

func (handler *WalletHandler) FindByUrl(c echo.Context, walletId *domain.Wallet) (err error) {
	ctx := c.Request().Context()

	wallet, err := handler.WalletUsecase.First(ctx, walletId)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, wallet)
}

func (handler *WalletHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := handler.WalletUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *WalletHandler) Store(c echo.Context) (err error) {
	var walletRequest domain.WalletStoreRequest

	err = c.Bind(&walletRequest)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err = c.Validate(walletRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	err = handler.WalletUsecase.Store(ctx, &walletRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, helper.HttpStatusCreatedMessage)
}

// Update will store the Wallet by given request body
func (handler *WalletHandler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var walletRequest domain.WalletUpdateRequest
	err = c.Bind(&walletRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(walletRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	wallet, err := handler.WalletUsecase.First(ctx, &domain.Wallet{Id: walletRequest.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.WalletUsecase.Update(ctx, &wallet, &walletRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusUpdatedMessage)
}

// Delete will remove the Wallet by given id
func (handler *WalletHandler) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var walletData domain.Wallet

	err = c.Bind(&walletData)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	wallet, err := handler.WalletUsecase.First(ctx, &domain.Wallet{Id: walletData.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.WalletUsecase.Delete(ctx, &wallet)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusDeletedMessage)
}
