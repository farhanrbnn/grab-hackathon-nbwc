package domain

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username     string             `json:"username,omitempty" bson:"username,omitempty"`
	PasswordHash string             `json:"password_hash,omitempty" bson:"password_hash,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`
	Phone        string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	IsVerified   bool               `json:"is_verified,omitempty" bson:"is_verified,omitempty"`
	IsAdmin      bool               `json:"is_admin,omitempty" bson:"is_admin,omitempty"`
	CreatedAt    int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt    int64              `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type UserTokenData struct {
	Id         string `json:"id,omitempty" bson:"_id,omitempty"`
	Username   string `json:"username,omitempty" bson:"username,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Phone      string `json:"phone,omitempty" bson:"phone,omitempty"`
	Email      string `json:"email,omitempty" bson:"email,omitempty"`
	IsVerified bool   `json:"is_verified,omitempty" bson:"is_verified,omitempty"`
}

type UserStoreRequest struct {
	Username string `json:"username,omitempty" bson:"username,omitempty" validate:"required,min=8,max=24"`
	Password string `json:"password,omitempty" bson:"password,omitempty" validate:"required,min=8,max=24"`
	Name     string `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=4"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty" validate:"required,min=8,max=16"`
	Email    string `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
}

type UserUpdateRequest struct {
	Username string `json:"username,omitempty" bson:"username,omitempty" validate:"required,min=8,max=24"`
	Name     string `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=4"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty" validate:"required,min=8,max=16"`
	Email    string `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
}

type Credential struct {
	Phone        string `json:"phone,omitempty" bson:"phone,omitempty" validate:"required,min=8,max=16"`
	Password     string `json:"password,omitempty" bson:"password,omitempty" validate:"required"`
	PasswordHash string `json:"password_hash,omitempty" bson:"password_hash,omitempty"`
}

type Claims struct {
	UserTokenData *UserTokenData `json:"user,omitempty" bson:"user,omitempty"`
	IsAdmin       bool           `json:"is_admin,omitempty" bson:"is_admin,omitempty"`
	jwt.RegisteredClaims
}

type JwtToken struct {
	Token     string `json:"token,omitempty" bson:"token,omitempty"`
	ExpiresAt int64  `json:"expires_at,omitempty" bson:"expires_at,omitempty"`
}

type UserUsecase interface {
	Signin(context.Context, *Credential, *User) (JwtToken, error)
	First(context.Context, *User) (User, error)
	Fetch(context.Context) ([]User, error)
	Store(context.Context, *UserStoreRequest) error
	Update(context.Context, *User, *UserUpdateRequest) error
	Delete(context.Context, *User) error
}

type UserRepository interface {
	First(context.Context, *User) (User, error)
	Fetch(context.Context) ([]User, error)
	Store(context.Context, *User) error
	Update(context.Context, *User) error
	Delete(context.Context, *User) error
}
