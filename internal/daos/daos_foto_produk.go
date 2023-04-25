package daos

import "gorm.io/gorm"

type FotoProduk struct {
	gorm.Model
	IdProduk uint
	Url      string
}
