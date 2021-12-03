package application

import (
	"giapps/servisin/domain/entity"
	"giapps/servisin/domain/repository"
	"giapps/servisin/domain/vo"
	"giapps/servisin/infrastructure/exception"
)

type WilayahAppInterface interface {
	GetProvinsi() []entity.WilayahEntity
	GetKota(request *entity.WilayahRequest) []entity.WilayahEntity
	GetKecamatan(request *entity.WilayahRequest) []entity.WilayahEntity
	GetKelurahan(request *entity.WilayahRequest) []entity.WilayahEntity
}

type wilayahAppInterfaceImpl struct {
	wilayahRepo repository.WilayahRepository
}

func NewWilayahAppInterface(wilayahRepo *repository.WilayahRepository) WilayahAppInterface {
	return &wilayahAppInterfaceImpl{wilayahRepo: *wilayahRepo}
}

func (app *wilayahAppInterfaceImpl) GetProvinsi() []entity.WilayahEntity {
	provinsi := app.wilayahRepo.GetProvinsi()
	return provinsi
}

func (app *wilayahAppInterfaceImpl) GetKota(request *entity.WilayahRequest) []entity.WilayahEntity {
	kdprov, err := vo.NewKodeProvinsi(request.ProvId)
	exception.ErrValidationIfNeeded(err)
	kota := app.wilayahRepo.GetKota(kdprov)
	return kota
}

func (app *wilayahAppInterfaceImpl) GetKecamatan(request *entity.WilayahRequest) []entity.WilayahEntity {
	kdprov, err := vo.NewKodeProvinsi(request.ProvId)
	exception.ErrValidationIfNeeded(err)
	kdkota, err := vo.NewKodeKota(request.KotaId)
	exception.ErrValidationIfNeeded(err)
	kecamatan := app.wilayahRepo.GetKecamatan(kdprov, kdkota)
	return kecamatan
}

func (app *wilayahAppInterfaceImpl) GetKelurahan(request *entity.WilayahRequest) []entity.WilayahEntity {
	kdprov, err := vo.NewKodeProvinsi(request.ProvId)
	exception.ErrValidationIfNeeded(err)
	kdkota, err := vo.NewKodeKota(request.KotaId)
	exception.ErrValidationIfNeeded(err)
	kdkecamatan, err := vo.NewKodeKecamatan(request.KecId)
	exception.ErrValidationIfNeeded(err)

	kelurahan := app.wilayahRepo.GetKelurahan(kdprov, kdkota, kdkecamatan)
	return kelurahan
}
