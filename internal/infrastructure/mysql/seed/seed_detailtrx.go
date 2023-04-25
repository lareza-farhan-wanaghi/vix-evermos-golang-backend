package seed

import (
	"tugas_akhir_example/internal/daos"
)

var DetailTrxSeed = []daos.DetailTrx{
	{
		IdTrx:       1,
		IdLogProduk: 1,
		IdToko:      1,
		Kuantitas:   1,
		HargaTotal:  75000,
	},
	{
		IdTrx:       2,
		IdLogProduk: 2,
		IdToko:      2,
		Kuantitas:   1,
		HargaTotal:  100000,
	},
	{
		IdTrx:       2,
		IdLogProduk: 3,
		IdToko:      2,
		Kuantitas:   2,
		HargaTotal:  60000,
	},
}
