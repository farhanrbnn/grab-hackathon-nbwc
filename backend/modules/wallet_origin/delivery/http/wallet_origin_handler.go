package http

// Need to use echo jwt version
import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"grab-hack-for-good/domain"
	"grab-hack-for-good/helper"
)

// WalletOriginHandler  represent the httphandler for walletOrigin
type WalletOriginHandler struct {
	WalletOriginUsecase domain.WalletOriginUsecase
}

func NewWalletOriginHandler(e *echo.Echo, r *echo.Group, cs domain.WalletOriginUsecase) {
	handler := &WalletOriginHandler{
		WalletOriginUsecase: cs,
	}

	r.GET("/api/v1/wallet-origins", handler.Fetch) // Testing, Admin only
	r.GET("/api/v1/wallet-origin/:id", handler.FindByUrlId)
	r.POST("/api/v1/wallet-origin/create", handler.Store)
	r.PATCH("/api/v1/wallet-origin/update", handler.Update)
	r.DELETE("/api/v1/wallet-origin/delete", handler.Delete)
}

func (handler *WalletOriginHandler) FindByUrlId(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return handler.FindByUrl(c, &domain.WalletOrigin{Id: id})
}

func (handler *WalletOriginHandler) FindByUrl(c echo.Context, walletOriginId *domain.WalletOrigin) (err error) {
	ctx := c.Request().Context()

	walletOrigin, err := handler.WalletOriginUsecase.First(ctx, walletOriginId)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, walletOrigin)
}

func (handler *WalletOriginHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := handler.WalletOriginUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *WalletOriginHandler) Store(c echo.Context) (err error) {
	var walletOriginRequest domain.WalletOriginRequest

	err = c.Bind(&walletOriginRequest)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err = c.Validate(walletOriginRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	err = handler.WalletOriginUsecase.Store(ctx, &walletOriginRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, helper.HttpStatusCreatedMessage)
}

// Update will store the WalletOrigin by given request body
func (handler *WalletOriginHandler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var walletOriginRequest domain.WalletOriginRequest

	err = c.Bind(&walletOriginRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(walletOriginRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	walletOrigin, err := handler.WalletOriginUsecase.First(ctx, &domain.WalletOrigin{Id: walletOriginRequest.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.WalletOriginUsecase.Update(ctx, &walletOrigin, &walletOriginRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusUpdatedMessage)
}

// Delete will remove the WalletOrigin by given id
func (handler *WalletOriginHandler) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var walletOriginData domain.WalletOrigin

	err = c.Bind(&walletOriginData)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	walletOrigin, err := handler.WalletOriginUsecase.First(ctx, &domain.WalletOrigin{Id: walletOriginData.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.WalletOriginUsecase.Delete(ctx, &walletOrigin)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusDeletedMessage)
}
