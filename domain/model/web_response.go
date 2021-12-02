package model

import (
	"net/http"

	"github.com/go-chi/render"
)

type WebResponse struct {
	Code   int32       `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func (web *WebResponse) Render(w http.ResponseWriter, r *http.Request) error {
	if web.Code != 0 {
		render.Status(r, int(web.Code))
	}
	return nil
}
