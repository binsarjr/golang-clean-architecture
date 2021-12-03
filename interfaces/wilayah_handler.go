package interfaces

import (
	"encoding/json"
	"giapps/servisin/application"
	"giapps/servisin/domain/entity"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Wilayah struct {
	WilayahApp application.WilayahAppInterface
}

func NewWilayah(wilayahApp *application.WilayahAppInterface) Wilayah {
	return Wilayah{WilayahApp: *wilayahApp}
}

func (handler *Wilayah) GetProvinsi(w http.ResponseWriter, r *http.Request) {
	response := handler.WilayahApp.GetProvinsi()
	json.NewEncoder(w).Encode(response)
}
func (handler *Wilayah) GetKota(w http.ResponseWriter, r *http.Request) {
	request := entity.WilayahRequest{
		ProvId: chi.URLParam(r, "provId"),
	}

	response := handler.WilayahApp.GetKota(&request)
	json.NewEncoder(w).Encode(response)
}

func (handler *Wilayah) GetKecamatan(w http.ResponseWriter, r *http.Request) {
	request := entity.WilayahRequest{
		ProvId: chi.URLParam(r, "provId"),
		KotaId: chi.URLParam(r, "kotaId"),
	}

	response := handler.WilayahApp.GetKecamatan(&request)
	json.NewEncoder(w).Encode(response)
}
func (handler *Wilayah) GetKelurahan(w http.ResponseWriter, r *http.Request) {
	request := entity.WilayahRequest{
		ProvId: chi.URLParam(r, "provId"),
		KotaId: chi.URLParam(r, "kotaId"),
		KecId:  chi.URLParam(r, "kecId"),
	}

	response := handler.WilayahApp.GetKelurahan(&request)
	json.NewEncoder(w).Encode(response)
}
