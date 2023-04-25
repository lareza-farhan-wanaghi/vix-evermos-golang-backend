package daos

import "gorm.io/gorm"

type DetailTrx struct {
	gorm.Model
	IdTrx       uint
	IdLogProduk uint
	IdToko      uint
	Kuantitas   int
	HargaTotal  int

	LogProduk *LogProduk `gorm:"foreignKey:IdLogProduk"`
}
