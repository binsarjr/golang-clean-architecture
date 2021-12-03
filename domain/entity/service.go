package entity

type ServiceKategoriEntity struct {
	KategoriId   int    `json:"kategori_id"`
	Icon         string `json:"icon"`
	NamaKategori string `json:"nama_kategori"`
}

type ServiceMerkEntity struct {
	MerkId     int    `json:"merk_id"`
	KategoriId int    `json:"kategori_id"`
	Merk       string `json:"merk"`
	Tipe       string `json:"tipe"`
}
