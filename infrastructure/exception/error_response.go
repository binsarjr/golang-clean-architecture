package exception

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("tst", e)
	if e.Code != 0 {
		render.Status(r, e.Code)
	} else {
		render.Status(r, http.StatusBadRequest)
	}
	return nil
}

func (errResponse ErrResponse) Error() string {
	return errResponse.Message
}
