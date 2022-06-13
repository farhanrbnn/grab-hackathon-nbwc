package http

// Need to use echo jwt version
import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"grab-hack-for-good/domain"
	"grab-hack-for-good/helper"
)

// UserWalletHandler  represent the httphandler for userWallet
type UserWalletHandler struct {
	UserWalletUsecase domain.UserWalletUsecase
}

func NewUserWalletHandler(e *echo.Echo, r *echo.Group, cs domain.UserWalletUsecase) {
	handler := &UserWalletHandler{
		UserWalletUsecase: cs,
	}

	// r.GET("/api/v1/user-wallets", handler.Fetch)         // Testing, Admin only
	r.GET("/api/v1/user-wallets", handler.FetchByUserId)
	r.GET("/api/v1/user-wallet/:id", handler.FindByUrlId)
	r.POST("/api/v1/user-wallet/create", handler.Store)
	r.PATCH("/api/v1/user-wallet/update", handler.Update)
	r.DELETE("/api/v1/user-wallet/delete", handler.Delete)
}

func (handler *UserWalletHandler) FindByUrlId(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return handler.FindByUrl(c, &domain.UserWallet{Id: id})
}

func (handler *UserWalletHandler) FindByUrl(c echo.Context, userWalletId *domain.UserWallet) (err error) {
	ctx := c.Request().Context()

	userWallet, err := handler.UserWalletUsecase.First(ctx, userWalletId)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, userWallet)
}

func (handler *UserWalletHandler) FetchByUserId(c echo.Context) error {
	ctx := c.Request().Context()

	userId, err := helper.GetUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	res, err := handler.UserWalletUsecase.FetchByUserId(ctx, &domain.UserWallet{
		UserId: userId,
	})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *UserWalletHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := handler.UserWalletUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *UserWalletHandler) Get(c echo.Context) error {
	user := helper.GetUserByToken(c)

	return c.JSON(http.StatusOK, user)
}

func (handler *UserWalletHandler) Store(c echo.Context) (err error) {
	ctx := c.Request().Context()

	userId, err := helper.GetUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var userWalletStoreRequest domain.UserWalletStoreRequest
	err = c.Bind(&userWalletStoreRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	userWalletStoreRequest.UserId = userId

	if err = c.Validate(userWalletStoreRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	userWallet, err := handler.UserWalletUsecase.First(ctx, &domain.UserWallet{UserId: userId, WalletId: userWalletStoreRequest.WalletId})
	if userWallet != (domain.UserWallet{}) && err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if userWallet == (domain.UserWallet{}) {
		err = handler.UserWalletUsecase.Store(ctx, &userWalletStoreRequest)
		if err != nil {
			return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
		}
	} else {
		err = helper.ErrConflict
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, helper.HttpStatusCreatedMessage)
}

// Update will store the UserWallet by given request body
func (handler *UserWalletHandler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()

	userId, err := helper.GetUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var userWalletUpdateRequest domain.UserWalletUpdateRequest
	err = c.Bind(&userWalletUpdateRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(userWalletUpdateRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	userWallet, err := handler.UserWalletUsecase.First(ctx, &domain.UserWallet{Id: userWalletUpdateRequest.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if userId == userWallet.UserId {
		err = handler.UserWalletUsecase.Update(ctx, &userWallet, &userWalletUpdateRequest)
		if err != nil {
			return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
		}
	} else {
		err = helper.ErrUnauthorized
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusUpdatedMessage)
}

// Delete will remove the UserWallet by given id
func (handler *UserWalletHandler) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()

	userId, err := helper.GetUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var userWalletData domain.UserWallet

	err = c.Bind(&userWalletData)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	userWallet, err := handler.UserWalletUsecase.First(ctx, &domain.UserWallet{Id: userWalletData.Id})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if userId == userWallet.UserId {
		err = handler.UserWalletUsecase.Delete(ctx, &userWallet)
		if err != nil {
			return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
		}
	} else {
		err = helper.ErrUnauthorized
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusDeletedMessage)
}
