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

// DropOffLocationHandler  represent the httphandler for dropOffLocation
type DropOffLocationHandler struct {
	DropOffLocationUsecase domain.DropOffLocationUsecase
}

func NewDropOffLocationHandler(e *echo.Echo, r *echo.Group, cs domain.DropOffLocationUsecase) {
	handler := &DropOffLocationHandler{
		DropOffLocationUsecase: cs,
	}

	r.GET("/api/v1/drop-off-locations", handler.Fetch) // Testing, Admin only
	r.GET("/api/v1/drop-off-location/:id", handler.FindByUrlId)
	r.POST("/api/v1/drop-off-location/create", handler.Store)
	r.PATCH("/api/v1/drop-off-location/update", handler.Update)
	r.DELETE("/api/v1/drop-off-location/delete", handler.Delete)
}

func (handler *DropOffLocationHandler) FindByUrlId(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return handler.FindByUrl(c, &domain.DropOffLocation{Id: id})
}

func (handler *DropOffLocationHandler) FindByUrl(c echo.Context, dropOffLocationId *domain.DropOffLocation) (err error) {
	ctx := c.Request().Context()

	dropOffLocation, err := handler.DropOffLocationUsecase.First(ctx, dropOffLocationId)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dropOffLocation)
}

func (handler *DropOffLocationHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := handler.DropOffLocationUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *DropOffLocationHandler) Get(c echo.Context) error {
	user := helper.GetUserByToken(c)

	return c.JSON(http.StatusOK, user)
}

func (handler *DropOffLocationHandler) Store(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var dropOffLocationStoreRequest domain.DropOffLocationStoreRequest
	err = c.Bind(&dropOffLocationStoreRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err = c.Validate(dropOffLocationStoreRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = handler.DropOffLocationUsecase.Store(ctx, &dropOffLocationStoreRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, helper.HttpStatusCreatedMessage)
}

func (handler *DropOffLocationHandler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var dropOffLocationUpdateRequest domain.DropOffLocationUpdateRequest
	err = c.Bind(&dropOffLocationUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(dropOffLocationUpdateRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	fmt.Print(dropOffLocationUpdateRequest.Id)

	dropOffLocation, err := handler.DropOffLocationUsecase.First(ctx, &domain.DropOffLocation{Id: dropOffLocationUpdateRequest.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.DropOffLocationUsecase.Update(ctx, &dropOffLocation, &dropOffLocationUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusUpdatedMessage)
}

// Delete will remove the DropOffLocation by given id
func (handler *DropOffLocationHandler) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var dropOffLocationData domain.DropOffLocation

	err = c.Bind(&dropOffLocationData)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	dropOffLocation, err := handler.DropOffLocationUsecase.First(ctx, &domain.DropOffLocation{Id: dropOffLocationData.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.DropOffLocationUsecase.Delete(ctx, &dropOffLocation)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusDeletedMessage)
}
