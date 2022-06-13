package helper

import (
	"encoding/json"
	"grab-hack-for-good/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetJSONRawBody(c echo.Context) (res map[string]interface{}, err error) {

	res = make(map[string]interface{})
	err = json.NewDecoder(c.Request().Body).Decode(&res)
	if err != nil {
		log.Error("empty json body")
	}

	return
}

func GetUserByToken(c echo.Context) (user domain.UserTokenData) {
	data := c.Get("user").(*jwt.Token)
	claims := data.Claims.(jwt.MapClaims)

	userClaim := claims["user"]

	mapstructure.Decode(userClaim, &user)

	return user
}

func GetUserId(c echo.Context) (id primitive.ObjectID, err error) {
	user := GetUserByToken(c)

	id, err = primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return
	}

	return
}
