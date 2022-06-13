package http

// Need to use echo jwt version
import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"grab-hack-for-good/domain"
	"grab-hack-for-good/helper"
)

// MerchantHandler  represent the httphandler for merchant
type MerchantHandler struct {
	MerchantUsecase domain.MerchantUsecase
}

func NewMerchantHandler(e *echo.Echo, r *echo.Group, cs domain.MerchantUsecase) {
	handler := &MerchantHandler{
		MerchantUsecase: cs,
	}

	r.GET("/api/v1/merchants", handler.Fetch) // Testing, Admin only
	r.GET("/api/v1/merchant/:id", handler.FindByUrlId)
	r.POST("/api/v1/merchant/create", handler.Store)
	r.PATCH("/api/v1/merchant/update", handler.Update)
	r.DELETE("/api/v1/merchant/delete", handler.Delete)
}

func (handler *MerchantHandler) FindByUrlId(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return handler.FindByUrl(c, &domain.Merchant{Id: id})
}

func (handler *MerchantHandler) FindByUrl(c echo.Context, merchantId *domain.Merchant) (err error) {
	ctx := c.Request().Context()

	merchant, err := handler.MerchantUsecase.First(ctx, merchantId)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, merchant)
}

func (handler *MerchantHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := handler.MerchantUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *MerchantHandler) Get(c echo.Context) error {
	user := helper.GetUserByToken(c)

	return c.JSON(http.StatusOK, user)
}

func (handler *MerchantHandler) Store(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var merchantStoreRequest domain.MerchantStoreRequest
	err = c.Bind(&merchantStoreRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err = c.Validate(merchantStoreRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = handler.MerchantUsecase.Store(ctx, &merchantStoreRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, helper.HttpStatusCreatedMessage)
}

func (handler *MerchantHandler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var merchantUpdateRequest domain.MerchantUpdateRequest
	err = c.Bind(&merchantUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(merchantUpdateRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	fmt.Print(merchantUpdateRequest.Id)

	merchant, err := handler.MerchantUsecase.First(ctx, &domain.Merchant{Id: merchantUpdateRequest.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.MerchantUsecase.Update(ctx, &merchant, &merchantUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusUpdatedMessage)
}

// Delete will remove the Merchant by given id
func (handler *MerchantHandler) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var merchantData domain.Merchant

	err = c.Bind(&merchantData)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	merchant, err := handler.MerchantUsecase.First(ctx, &domain.Merchant{Id: merchantData.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.MerchantUsecase.Delete(ctx, &merchant)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusDeletedMessage)
}
