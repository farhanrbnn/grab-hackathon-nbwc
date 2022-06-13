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

// ProductHandler  represent the httphandler for product
type ProductHandler struct {
	ProductUsecase domain.ProductUsecase
}

func NewProductHandler(e *echo.Echo, r *echo.Group, cs domain.ProductUsecase) {
	handler := &ProductHandler{
		ProductUsecase: cs,
	}

	r.GET("/api/v1/products", handler.Fetch) // Testing, Admin only
	r.GET("/api/v1/products/:merchant_id", handler.FetchByMerchantId)
	r.GET("/api/v1/product/:id", handler.FindByUrlId)
	r.POST("/api/v1/product/create", handler.Store)
	r.PATCH("/api/v1/product/update", handler.Update)
	r.DELETE("/api/v1/product/delete", handler.Delete)
}

func (handler *ProductHandler) FindByUrlId(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return handler.FindByUrl(c, &domain.Product{Id: id})
}

func (handler *ProductHandler) FindByUrl(c echo.Context, productId *domain.Product) (err error) {
	ctx := c.Request().Context()

	product, err := handler.ProductUsecase.First(ctx, productId)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, product)
}

func (handler *ProductHandler) FetchByMerchantId(c echo.Context) error {
	ctx := c.Request().Context()

	merchantId, err := primitive.ObjectIDFromHex(c.Param("merchant_id"))
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	res, err := handler.ProductUsecase.FetchByMerchantId(ctx, &domain.Product{
		MerchantId: merchantId,
	})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *ProductHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := handler.ProductUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *ProductHandler) Get(c echo.Context) error {
	user := helper.GetUserByToken(c)

	return c.JSON(http.StatusOK, user)
}

func (handler *ProductHandler) Store(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var productStoreRequest domain.ProductStoreRequest
	err = c.Bind(&productStoreRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err = c.Validate(productStoreRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = handler.ProductUsecase.Store(ctx, &productStoreRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, helper.HttpStatusCreatedMessage)
}

func (handler *ProductHandler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var productUpdateRequest domain.ProductUpdateRequest
	err = c.Bind(&productUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(productUpdateRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	fmt.Print(productUpdateRequest.Id)

	product, err := handler.ProductUsecase.First(ctx, &domain.Product{Id: productUpdateRequest.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.ProductUsecase.Update(ctx, &product, &productUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusUpdatedMessage)
}

// Delete will remove the Product by given id
func (handler *ProductHandler) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var productData domain.Product

	err = c.Bind(&productData)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	product, err := handler.ProductUsecase.First(ctx, &domain.Product{Id: productData.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.ProductUsecase.Delete(ctx, &product)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusDeletedMessage)
}
