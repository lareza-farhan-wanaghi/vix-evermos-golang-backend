package seed

import (
	"tugas_akhir_example/internal/daos"
)

var TrxSeed = []daos.Trx{
	{
		IdUser:           1,
		AlamatPengiriman: 1,
		HargaTotal:       75000,
		KodeInvoice:      "Kode-A",
		MethodBayar:      "bca",
	},
	{
		IdUser:           2,
		AlamatPengiriman: 2,
		HargaTotal:       5000,
		KodeInvoice:      "Kode-B",
		MethodBayar:      "bca",
	},
}
