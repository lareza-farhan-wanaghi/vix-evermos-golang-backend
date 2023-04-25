package daos

import "gorm.io/gorm"

type Trx struct {
	gorm.Model
	IdUser           uint
	AlamatPengiriman uint
	HargaTotal       int
	KodeInvoice      string
	MethodBayar      string

	User       *User        `gorm:"foreignKey:IdUser"`
	DetailTrxs []*DetailTrx `gorm:"foreignKey:IdTrx"`
	Alamat     *Alamat      `gorm:"foreignKey:AlamatPengiriman"`
}

type FilterTrx struct {
	Limit, Offset int
	KodeInvoice   string
}
