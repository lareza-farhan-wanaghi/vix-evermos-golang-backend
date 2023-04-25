package daos

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama         string
	KataSandi    string
	Notelp       string `gorm:"unique"`
	TanggalLahir time.Time
	JenisKelamin string
	Tentang      string `gorm:"type:text"`
	Pekerjaan    string
	Email        string `gorm:"unique"`
	IdProvinsi   string
	IdKota       string
	IsAdmin      bool

	Toko    *Toko     `gorm:"foreignKey:IdUser"`
	Alamats []*Alamat `gorm:"foreignKey:IdUser"`
}
