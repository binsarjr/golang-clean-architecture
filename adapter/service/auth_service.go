package service

import (
	"fmt"
	"giapps/servisin/adapter/repository"
	"giapps/servisin/domain/entity"
	"giapps/servisin/domain/model"
	"giapps/servisin/exception"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignIn(request model.AuthSignInRequest) (response model.AuthResponse)
	SignUp(request model.AuthSignUpRequest) (response model.AuthResponse)
	CreateJWTToken(user *entity.UserEntity) (token string)
}

type authServiceImpl struct {
	UserRepository repository.UserRepository
	tokenAuth      *jwtauth.JWTAuth
}

func (service *authServiceImpl) CreateJWTToken(user *entity.UserEntity) (token string) {
	_, tokenString, _ := service.tokenAuth.Encode(map[string]interface{}{
		"user_id": user.UserId,
	})
	return tokenString
}

func (service *authServiceImpl) SignIn(request model.AuthSignInRequest) (response model.AuthResponse) {

	userauth, err := service.UserRepository.FindUserByUsername(request.Username)
	exception.PanicIfNeeded(err)

	err = bcrypt.CompareHashAndPassword([]byte(userauth.Password), []byte(request.Password))
	if err != nil {
		panic(exception.ErrResponse{Code: http.StatusUnauthorized, Message: "username dan password tidak ditemukan"})
	}

	tokenString := service.CreateJWTToken(&userauth)

	return service.response(tokenString)
}

func (service *authServiceImpl) SignUp(request model.AuthSignUpRequest) (response model.AuthResponse) {
	var user entity.UserEntity
	user, _ = service.UserRepository.FindUserByUsernameOrEmail(request.Username, request.Email)
	if user != (entity.UserEntity{}) {
		var field string
		if user.Username != "" {
			field = "username"
		}
		if user.Email != "" {
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
	service.UserRepository.Insert(&userauth)

	tokenString := service.CreateJWTToken(&userauth)

	return service.response(tokenString)
}

func (service *authServiceImpl) response(token string) (response model.AuthResponse) {

	response = model.AuthResponse{
		Type:        "Bearer",
		AccessToken: token,
	}

	return response
}

func NewAuthSerivce(repository *repository.UserRepository, tokenAuth *jwtauth.JWTAuth) AuthService {
	return &authServiceImpl{
		UserRepository: *repository,
		tokenAuth:      tokenAuth,
	}
}
