package repository

import (
	"fmt"
	"giapps/servisin/domain/entity"
	"giapps/servisin/domain/vo"
	"giapps/servisin/infrastructure/database"
	"giapps/servisin/infrastructure/exception"

	"github.com/jackc/pgx/v4/pgxpool"
)

type WilayahRepository interface {
	GetProvinsi() []entity.WilayahEntity
	GetKota(provinsi vo.KodeProvinsi) []entity.WilayahEntity
	GetKecamatan(provinsi vo.KodeProvinsi, kota vo.KodeKota) []entity.WilayahEntity
	GetKelurahan(provinsi vo.KodeProvinsi, kota vo.KodeKota, kecamatan vo.KodeKecamatan) []entity.WilayahEntity
}

type wilayahRepositortImpl struct {
	db *pgxpool.Pool
}

func NewWilayahRepository(db *pgxpool.Pool) WilayahRepository {
	return &wilayahRepositortImpl{db: db}
}

func (repo wilayahRepositortImpl) GetProvinsi() []entity.WilayahEntity {
	ctx, cancel := database.NewDatabaseContext()
	defer cancel()

	rows, err := repo.db.Query(ctx, `
        SELECT
            wilayah_id,coalesce(provinsi,''),coalesce(kota,''),coalesce(kecamatan,''), coalesce(kelurahan,'')
        FROM manajemen.wilayah
        where 
            length(wilayah_id) = 2
    `)
	exception.PanicIfNeeded(err)
	defer rows.Close()

	provinsi := []entity.WilayahEntity{}
	for rows.Next() {
		prov := entity.WilayahEntity{}
		err := rows.Scan(&prov.WilayahId, &prov.Provinsi, &prov.Kota, &prov.Kecamatan, &prov.Kelurahan)
		exception.PanicIfNeeded(err)

		provinsi = append(provinsi, prov)
	}
	return provinsi
}

func (repo wilayahRepositortImpl) GetKota(provinsi vo.KodeProvinsi) []entity.WilayahEntity {
	ctx, cancel := database.NewDatabaseContext()
	defer cancel()

	rows, err := repo.db.Query(ctx, `SELECT wilayah_id,coalesce(provinsi,''),coalesce(kota,''),coalesce(kecamatan,''), coalesce(kelurahan,'') FROM manajemen.wilayah where length(wilayah_id) = 5 and wilayah_id ilike $1`, string(provinsi.Value())+".%")
	exception.PanicIfNeeded(err)
	defer rows.Close()

	kota := []entity.WilayahEntity{}
	for rows.Next() {
		kotaEntity := entity.WilayahEntity{}
		err := rows.Scan(&kotaEntity.WilayahId, &kotaEntity.Provinsi, &kotaEntity.Kota, &kotaEntity.Kecamatan, &kotaEntity.Kelurahan)
		exception.PanicIfNeeded(err)

		kota = append(kota, kotaEntity)
	}
	return kota
}

func (repo wilayahRepositortImpl) GetKecamatan(provinsi vo.KodeProvinsi, kota vo.KodeKota) []entity.WilayahEntity {

	ctx, cancel := database.NewDatabaseContext()
	defer cancel()

	rows, err := repo.db.Query(ctx, `
        SELECT
            wilayah_id,coalesce(provinsi,''),coalesce(kota,''),coalesce(kecamatan,''), coalesce(kelurahan,'')
        FROM manajemen.wilayah
        where 
            length(wilayah_id) = 8 and wilayah_id ilike $1
    `, fmt.Sprintf("%s.%s%s", provinsi.Value(), kota.Value(), "%"))
	exception.PanicIfNeeded(err)
	defer rows.Close()

	kecamatan := []entity.WilayahEntity{}
	for rows.Next() {
		kecamatanEntity := entity.WilayahEntity{}
		err := rows.Scan(&kecamatanEntity.WilayahId, &kecamatanEntity.Provinsi, &kecamatanEntity.Kota, &kecamatanEntity.Kecamatan, &kecamatanEntity.Kelurahan)
		exception.PanicIfNeeded(err)

		kecamatan = append(kecamatan, kecamatanEntity)
	}
	return kecamatan
}

func (repo wilayahRepositortImpl) GetKelurahan(provinsi vo.KodeProvinsi, kota vo.KodeKota, kecamatan vo.KodeKecamatan) []entity.WilayahEntity {

	ctx, cancel := database.NewDatabaseContext()
	defer cancel()

	rows, err := repo.db.Query(ctx, `
        SELECT
            wilayah_id,coalesce(provinsi,''),coalesce(kota,''),coalesce(kecamatan,''), coalesce(kelurahan,'')
        FROM manajemen.wilayah
        where 
            length(wilayah_id) = 13 and wilayah_id ilike $1
    `, fmt.Sprintf("%s.%s.%s%s", provinsi.Value(), kota.Value(), kecamatan.Value(), "%"))
	exception.PanicIfNeeded(err)
	defer rows.Close()

	kelurahan := []entity.WilayahEntity{}
	for rows.Next() {
		kelurahanEntity := entity.WilayahEntity{}
		err := rows.Scan(&kelurahanEntity.WilayahId, &kelurahanEntity.Provinsi, &kelurahanEntity.Kota, &kelurahanEntity.Kecamatan, &kelurahanEntity.Kelurahan)
		exception.PanicIfNeeded(err)

		kelurahan = append(kelurahan, kelurahanEntity)
	}
	return kelurahan
}
