package interfaces

import (
	"giapps/servisin/application"
	"giapps/servisin/domain/model"
	"giapps/servisin/domain/repository"
	"giapps/servisin/infrastructure/validation"

	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

type Authenticate struct {
	AuthenticateApp application.AuthenticateAppInterface
	UserRepo        repository.UserRepository
	tokenAuth       *jwtauth.JWTAuth
}

func NewAuthenticate(authenticateApp application.AuthenticateAppInterface, userRepo *repository.UserRepository, tokenAuth *jwtauth.JWTAuth) Authenticate {
	return Authenticate{AuthenticateApp: authenticateApp, tokenAuth: tokenAuth, UserRepo: *userRepo}
}

func (handler *Authenticate) Login(w http.ResponseWriter, r *http.Request) {
	request := &model.AuthLoginRequest{}

	validation.NewValidation(w, r, request)

	response := handler.AuthenticateApp.Login(request)
	render.Render(w, r, &response)
}

func (handler *Authenticate) Register(w http.ResponseWriter, r *http.Request) {
	request := &model.AuthRegisterRequest{}
	validation.NewValidation(w, r, request)

	response := handler.AuthenticateApp.Register(request)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &response)
}
