package seed

import (
	"tugas_akhir_example/internal/daos"
)

var LogProdukSeed = []daos.LogProduk{
	{
		IdProduk:      1,
		NamaProduk:    "ProdukA",
		Slug:          "produk-a",
		HargaReseller: "50000",
		HargaKonsumen: "75000",
		Deskripsi:     "Suatu deskripsi yang menjelaskan produk",
		IdToko:        1,
		IdCategory:    1,
	},
	{
		IdProduk:      2,
		NamaProduk:    "ProdukB",
		Slug:          "produk-b",
		HargaReseller: "75000",
		HargaKonsumen: "100000",
		Deskripsi:     "Suatu deskripsi yang menjelaskan produk",
		IdToko:        2,
		IdCategory:    1,
	},
	{
		IdProduk:      6,
		NamaProduk:    "ProdukF",
		Slug:          "produk-f",
		HargaReseller: "25000",
		HargaKonsumen: "30000",
		Deskripsi:     "Suatu deskripsi yang menjelaskan produk",
		IdToko:        2,
		IdCategory:    4,
	},
}
