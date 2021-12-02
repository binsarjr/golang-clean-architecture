package controller

import "github.com/go-chi/chi/v5"

type Controller interface {
	Route(c chi.Router)
}
