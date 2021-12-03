package interfaces

import (
	"fmt"
	"giapps/servisin/domain/entity"
	"giapps/servisin/domain/model"
	"giapps/servisin/domain/repository"
	"giapps/servisin/infrastructure/exception"
	"giapps/servisin/infrastructure/validation"

	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Authenticate struct {
	UserRepo  repository.UserRepository
	tokenAuth *jwtauth.JWTAuth
}

func NewAuthenticate(userRepo *repository.UserRepository, tokenAuth *jwtauth.JWTAuth) Authenticate {
	return Authenticate{tokenAuth: tokenAuth, UserRepo: *userRepo}
}

func (handler *Authenticate) CreateJWTToken(user *entity.UserEntity) (token string) {
	_, tokenString, _ := handler.tokenAuth.Encode(map[string]interface{}{
		"user_id": user.UserId,
	})
	return tokenString
}

func (handler *Authenticate) response(token string) (response model.AuthResponse) {

	response = model.AuthResponse{
		Type:        "Bearer",
		AccessToken: token,
	}

	return response
}

func (handler *Authenticate) Login(w http.ResponseWriter, r *http.Request) {
	request := &model.AuthLoginRequest{}

	validation.NewValidation(w, r, request)

	userauth, err := handler.UserRepo.FindUserByUsername(request.Username)
	exception.PanicIfNeeded(err)

	err = bcrypt.CompareHashAndPassword([]byte(userauth.Password), []byte(request.Password))
	if err != nil {
		panic(exception.ErrResponse{Code: http.StatusUnauthorized, Message: "username dan password tidak ditemukan"})
	}

	tokenString := handler.CreateJWTToken(&userauth)
	response := handler.response(tokenString)
	render.Render(w, r, &response)
}

func (handler *Authenticate) Register(w http.ResponseWriter, r *http.Request) {
	request := &model.AuthRegisterRequest{}
	validation.NewValidation(w, r, request)

	var user entity.UserEntity
	user, _ = handler.UserRepo.FindUserByUsernameOrEmail(request.Username, request.Email)
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
	handler.UserRepo.Insert(&userauth)

	tokenString := handler.CreateJWTToken(&userauth)
	response := handler.response(tokenString)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &response)
}
