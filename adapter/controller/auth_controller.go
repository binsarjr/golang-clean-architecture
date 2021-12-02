package controller

import (
	"giapps/newapp/adapter/service"
	"giapps/newapp/domain/model"
	"giapps/newapp/exception"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService *service.AuthService) AuthController {
	return AuthController{AuthService: *authService}
}

func (controller *AuthController) Route(r chi.Router) {
	authRouter := chi.NewRouter()
	authRouter.Post("/signin", controller.SignIn)
	authRouter.Post("/signup", controller.SignUp)

	r.Mount("/auth", authRouter)
}

func (controller *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	request := &model.AuthSignInRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, &exception.ErrResponse{Code: http.StatusUnprocessableEntity, Message: err.Error()})
		return
	}

	response := controller.AuthService.SignIn(*request)

	render.Render(w, r, &response)
}

func (controller *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	request := model.AuthSignUpRequest{}
	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, &exception.ErrResponse{Code: http.StatusUnprocessableEntity, Message: err.Error()})
		return
	}

	response := controller.AuthService.SignUp(request)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &response)
}
