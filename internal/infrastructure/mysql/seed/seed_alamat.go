package seed

import (
	"tugas_akhir_example/internal/daos"
)

var AlamatSeed = []daos.Alamat{
	{
		IdUser:       1,
		JudulAlamat:  "AlamatA1",
		NamaPenerima: "PenerimaA1",
		Notelp:       "0812345",
		DetailAlamat: "Jl. contoh alamat A1",
	},
	{
		IdUser:       1,
		JudulAlamat:  "AlamatA2",
		NamaPenerima: "PenerimaA2",
		Notelp:       "0812345678",
		DetailAlamat: "Jl. contoh alamat A2",
	},
	{
		IdUser:       2,
		JudulAlamat:  "AlamatB1",
		NamaPenerima: "PenerimaB1",
		Notelp:       "081234",
		DetailAlamat: "Jl. contoh alamat B1",
	},
	{
		IdUser:       3,
		JudulAlamat:  "AlamatC1",
		NamaPenerima: "PenerimaC1",
		Notelp:       "0812",
		DetailAlamat: "Jl. contoh alamat C1",
	},
	{
		IdUser:       4,
		JudulAlamat:  "AlamatD1",
		NamaPenerima: "PenerimaD1",
		Notelp:       "0812388888",
		DetailAlamat: "Jl. contoh alamat D1",
	},
	{
		IdUser:       5,
		JudulAlamat:  "AlamatE1",
		NamaPenerima: "PenerimaE1",
		Notelp:       "0812388878",
		DetailAlamat: "Jl. contoh alamat E1",
	},
}
