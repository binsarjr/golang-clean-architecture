package service

import (
	"giapps/newapp/adapter/repository"
	"giapps/newapp/domain/entity"
	"giapps/newapp/domain/model"
	"giapps/newapp/exception"
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
	secretKey      string
}

func (service *authServiceImpl) CreateJWTToken(user *entity.UserEntity) (token string) {
	tokenAuth := jwtauth.New("HS256", []byte(service.secretKey), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{
		"user_id": user.UserId,
	})
	return tokenString
}

func (service *authServiceImpl) SignIn(request model.AuthSignInRequest) (response model.AuthResponse) {

	userauth, err := service.UserRepository.FindUserUsername(request.Username)
	exception.PanicIfNeeded(err)

	err = bcrypt.CompareHashAndPassword([]byte(userauth.Password), []byte(request.Password))
	if err != nil {
		panic(exception.ErrResponse{Code: http.StatusUnauthorized, Message: "username dan password tidak ditemukan"})
	}

	tokenString := service.CreateJWTToken(&userauth)

	return service.response(tokenString)
}

func (service *authServiceImpl) SignUp(request model.AuthSignUpRequest) (response model.AuthResponse) {

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

func NewAuthSerivce(repository *repository.UserRepository, secretkey string) AuthService {
	return &authServiceImpl{
		UserRepository: *repository,
		secretKey:      secretkey,
	}
}
