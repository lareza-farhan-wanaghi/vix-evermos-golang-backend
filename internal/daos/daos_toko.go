package daos

import "gorm.io/gorm"

type Toko struct {
	gorm.Model
	IdUser   uint
	NamaToko string
	UrlFoto  string
}

type FilterToko struct {
	Limit, Offset int
	NamaToko      string
}
