package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"grab-hack-for-good/domain"
	"grab-hack-for-good/helper"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(cr domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       cr,
		contextTimeout: timeout,
	}
}

func (usecase *userUsecase) Signin(c context.Context, credential *domain.Credential, user *domain.User) (jwtToken domain.JwtToken, err error) {
	// ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	// defer cancel()

	hash := sha256.Sum256([]byte(credential.Password))
	credential.PasswordHash = hex.EncodeToString(hash[:])

	if credential.PasswordHash != user.PasswordHash {
		err = helper.ErrUnauthorized
		return domain.JwtToken{}, err
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	userTokenData := &domain.UserTokenData{
		Id:         user.Id.Hex(),
		Username:   user.Username,
		Name:       user.Name,
		Phone:      user.Phone,
		Email:      user.Email,
		IsVerified: user.IsVerified,
	}

	claims := &domain.Claims{
		userTokenData,
		user.IsAdmin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString([]byte(viper.GetString("jwt.secret_key")))

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		err = helper.ErrInternalServerError
		return domain.JwtToken{}, err
	}

	jwtToken = domain.JwtToken{
		Token:     tokenString,
		ExpiresAt: expirationTime.Unix(),
	}

	return
}

func (usecase *userUsecase) First(c context.Context, userFilter *domain.User) (user domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	user, err = usecase.userRepo.First(ctx, userFilter)
	if err != nil {
		return domain.User{}, err
	}

	return
}

func (usecase *userUsecase) Fetch(c context.Context) (res []domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.userRepo.Fetch(ctx)
	if err != nil {
		return make([]domain.User, 0), err
	}

	return
}

func (usecase *userUsecase) Store(c context.Context, userRequest *domain.UserStoreRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	user := domain.User{
		Username: userRequest.Username,
		Name:     userRequest.Name,
		Phone:    userRequest.Phone,
		Email:    userRequest.Email,
	}

	hash := sha256.Sum256([]byte(userRequest.Password))
	user.PasswordHash = hex.EncodeToString(hash[:])
	user.IsVerified = true
	user.IsAdmin = true

	err = usecase.userRepo.Store(ctx, &user)

	return
}

func (usecase *userUsecase) Update(c context.Context, user *domain.User, userData *domain.UserUpdateRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	user.Name = userData.Name
	user.Phone = userData.Phone
	user.Email = userData.Email
	user.UpdatedAt = time.Now().Unix()

	return usecase.userRepo.Update(ctx, user)
}

func (usecase *userUsecase) Delete(c context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	err = usecase.userRepo.Delete(ctx, user)

	return
}
