package validation

import (
	"giapps/servisin/infrastructure/exception"
	"net/http"

	"github.com/go-chi/render"
)

func NewValidation(w http.ResponseWriter, r *http.Request, request render.Binder) {
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, &exception.ErrResponse{Code: http.StatusUnprocessableEntity, Message: err.Error()})
	}
}
