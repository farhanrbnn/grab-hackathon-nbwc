package http

// Need to use echo jwt version
import (
	"net/http"

	"github.com/labstack/echo/v4"

	"grab-hack-for-good/domain"
	"grab-hack-for-good/helper"
	"grab-hack-for-good/modules/grab_express/usecase"
)

// GrabExpressHandler  represent the httphandler for grabExpress
type GrabExpressHandler struct {
	GrabExpressUsecase domain.GrabExpressUsecase
}

func NewGrabExpressHandler(e *echo.Echo, r *echo.Group, cs domain.GrabExpressUsecase) {
	handler := &GrabExpressHandler{
		GrabExpressUsecase: cs,
	}

	r.GET("/api/v1/grab-express/get-token", handler.GetToken) // Testing, Admin only
}

func (handler *GrabExpressHandler) GetToken(c echo.Context) (err error) {
	ctx := c.Request().Context()

	err = handler.GrabExpressUsecase.GetToken(ctx)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, usecase.GrabExpressToken)
}

func (handler *GrabExpressHandler) CreateDeliveryRequest(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var grabExpressDeliveryRequest domain.GrabExpressDeliveryRequest
	err = c.Bind(&grabExpressDeliveryRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err = c.Validate(grabExpressDeliveryRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	grabExpressDeliveryResponse, err := handler.GrabExpressUsecase.CreateDeliveryRequest(ctx, &grabExpressDeliveryRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, grabExpressDeliveryResponse)
}
