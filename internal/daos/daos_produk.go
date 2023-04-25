package daos

import "gorm.io/gorm"

type Produk struct {
	gorm.Model
	NamaProduk    string
	Slug          string
	HargaReseller string
	HargaKonsumen string
	Stok          int
	Deskripsi     string `gorm:"type:text"`
	IdToko        uint
	IdCategory    uint

	FotoProduks []*FotoProduk `gorm:"foreignKey:IdProduk"`
	Toko        *Toko         `gorm:"foreignKey:IdToko"`
	Category    *Category     `gorm:"foreignKey:IdCategory"`
}

type FilterProduk struct {
	Limit, Offset int
	NamaProduk    string
	CategoryId    uint
	TokoId        uint
	MaxHarga      int
	MinHarga      int
}
