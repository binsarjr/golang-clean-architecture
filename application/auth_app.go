package application

import (
	"fmt"
	"giapps/servisin/domain/entity"
	"giapps/servisin/domain/model"
	"giapps/servisin/domain/repository"
	"giapps/servisin/infrastructure/exception"

	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticateAppInterface interface {
	Login(request *model.AuthLoginRequest) (response model.AuthResponse)
	Register(request *model.AuthRegisterRequest) (response model.AuthResponse)
}

type authenticateAppInterfaceImpl struct {
	UserRepo  repository.UserRepository
	tokenAuth *jwtauth.JWTAuth
}

func NewAuthenticateAppInterface(userRepo *repository.UserRepository, tokenAuth *jwtauth.JWTAuth) AuthenticateAppInterface {
	return &authenticateAppInterfaceImpl{UserRepo: *userRepo, tokenAuth: tokenAuth}
}

func (app *authenticateAppInterfaceImpl) CreateJWTToken(user entity.UserEntity) (token string) {
	_, tokenString, _ := app.tokenAuth.Encode(map[string]interface{}{
		"user_id": user.UserId,
	})
	return tokenString
}

func (app *authenticateAppInterfaceImpl) response(token string) (response model.AuthResponse) {

	response = model.AuthResponse{
		Type:        "Bearer",
		AccessToken: token,
	}

	return response
}

func (app *authenticateAppInterfaceImpl) Login(request *model.AuthLoginRequest) (response model.AuthResponse) {
	userauth, err := app.UserRepo.FindUserByUsername(request.Username)
	exception.PanicIfNeeded(err)

	err = bcrypt.CompareHashAndPassword([]byte(userauth.Password), []byte(request.Password))
	if passwordMatch := err != nil; passwordMatch {
		panic(exception.ErrResponse{Code: http.StatusUnauthorized, Message: "username atau password tidak ditemukan"})
	}

	tokenString := app.CreateJWTToken(userauth)
	response = app.response(tokenString)
	return response
}

func (app *authenticateAppInterfaceImpl) Register(request *model.AuthRegisterRequest) (response model.AuthResponse) {
	var user entity.UserEntity
	user, _ = app.UserRepo.FindUserByUsernameOrEmail(request.Username, request.Email)

	if userNotFound := user != (entity.UserEntity{}); userNotFound {
		var field string
		if userNotNull := user.Username != ""; userNotNull {
			field = "username"
		}
		if emailNotNull := user.Email != ""; emailNotNull {
			field = "email"
		}
		panic(exception.ErrResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: fmt.Sprintf("%s sudah digunakan", field),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	exception.PanicIfNeeded(err)

	userauth := entity.UserEntity{
		UserId:   uuid.New().String(),
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
	}
	app.UserRepo.Insert(userauth)

	tokenString := app.CreateJWTToken(userauth)
	response = app.response(tokenString)
	return response
}
