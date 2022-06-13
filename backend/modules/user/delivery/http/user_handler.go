package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"grab-hack-for-good/domain"
	"grab-hack-for-good/helper"
)

// UserHandler  represent the httphandler for user
type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(e *echo.Echo, r *echo.Group, cs domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: cs,
	}

	e.POST("/api/v1/user/signin", handler.Signin)
	r.GET("/api/v1/user", handler.Get)
	r.GET("/api/v1/users", handler.Fetch) // Testing, Admin only
	r.GET("/api/v1/user/:id", handler.FindByUrlId)
	e.POST("/api/v1/user/create", handler.Store)
	r.PATCH("/api/v1/user/update", handler.Update)
	r.DELETE("/api/v1/user/delete", handler.Delete)
}

func (handler *UserHandler) Signin(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var credential domain.Credential

	err = c.Bind(&credential)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err = c.Validate(credential); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var user domain.User

	user, err = handler.UserUsecase.First(ctx, &domain.User{Phone: credential.Phone})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	token, err := handler.UserUsecase.Signin(ctx, &credential, &user)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, token)
}

func (handler *UserHandler) Get(c echo.Context) error {
	user := helper.GetUserByToken(c)

	return c.JSON(http.StatusOK, user)
}

func (handler *UserHandler) FindByUrlId(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return handler.FindByUrl(c, &domain.User{Id: id})
}

func (handler *UserHandler) FindByUrl(c echo.Context, userId *domain.User) (err error) {
	ctx := c.Request().Context()

	user, err := handler.UserUsecase.First(ctx, userId)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (handler *UserHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := handler.UserUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *UserHandler) Store(c echo.Context) (err error) {
	var userRequest domain.UserStoreRequest

	err = c.Bind(&userRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err = c.Validate(userRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()

	duplicatedUser, err := handler.UserUsecase.First(ctx, &domain.User{
		Username: userRequest.Username,
	})
	if duplicatedUser != (domain.User{}) && err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if duplicatedUser == (domain.User{}) {
		err = handler.UserUsecase.Store(ctx, &userRequest)
		if err != nil {
			return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
		}
	} else {
		err = helper.ErrConflict
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, helper.HttpStatusCreatedMessage)
}

// Update will store the User by given request body
func (handler *UserHandler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var userRequest domain.UserUpdateRequest

	err = c.Bind(&userRequest)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(userRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	userId, err := helper.GetUserId(c)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	user, err := handler.UserUsecase.First(ctx, &domain.User{Id: userId})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	var duplicateUser domain.User
	if user.Username != userRequest.Username {
		duplicateUser, err = handler.UserUsecase.First(ctx, &domain.User{
			Username: userRequest.Username,
		})
		if err != nil {
			return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
		}
	}

	if duplicateUser == (domain.User{}) {
		err = handler.UserUsecase.Update(ctx, &user, &userRequest)
		if err != nil {
			return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
		}
	} else {
		err = helper.ErrBadParamInput
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: "Username already exist"})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusUpdatedMessage)
}

// Delete will remove the User by given id
func (handler *UserHandler) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()

	userId, err := helper.GetUserId(c)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	user, err := handler.UserUsecase.First(ctx, &domain.User{Id: userId})
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	err = handler.UserUsecase.Delete(ctx, &user)
	if err != nil {
		return c.JSON(helper.GetStatusCode(err), helper.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, helper.HttpStatusDeletedMessage)
}
