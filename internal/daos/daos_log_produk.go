package daos

import "gorm.io/gorm"

type LogProduk struct {
	gorm.Model
	IdProduk      uint `gorm:"index"`
	NamaProduk    string
	Slug          string
	HargaReseller string
	HargaKonsumen string
	Deskripsi     string `gorm:"type:text"`
	IdToko        uint
	IdCategory    uint

	Produk   *Produk   `gorm:"foreignKey:IdProduk"`
	Toko     *Toko     `gorm:"foreignKey:IdToko"`
	Category *Category `gorm:"foreignKey:IdCategory"`
}
