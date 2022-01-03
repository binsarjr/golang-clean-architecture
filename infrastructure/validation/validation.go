package validation

import (
	"giapps/servisin/infrastructure/exception"
	"net/http"

	"github.com/go-chi/render"
)

func NewValidation(r *http.Request, request render.Binder) {
	if err := render.Bind(r, request); err != nil {
		panic(exception.ErrResponse{Code: http.StatusUnprocessableEntity, Message: err.Error()})
	}
}
