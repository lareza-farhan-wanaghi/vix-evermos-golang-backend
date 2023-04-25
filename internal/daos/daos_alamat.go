package daos

import "gorm.io/gorm"

type Alamat struct {
	gorm.Model
	IdUser       uint
	JudulAlamat  string
	NamaPenerima string
	Notelp       string
	DetailAlamat string
}

type FilterAlamat struct {
	JudulAlamat string
}
