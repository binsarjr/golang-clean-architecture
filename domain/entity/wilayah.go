package entity

import "net/http"

type WilayahEntity struct {
	WilayahId string `json:"wilayah_id"`
	Provinsi  string `json:"provinsi,omitempty"`
	Kota      string `json:"kota,omitempty"`
	Kecamatan string `json:"kecamatan,omitempty"`
	Kelurahan string `json:"kelurahan,omitempty"`
}

func (entity *WilayahEntity) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type WilayahRequest struct {
	ProvId string
	KotaId string
	KecId  string
	KelId  string
}
