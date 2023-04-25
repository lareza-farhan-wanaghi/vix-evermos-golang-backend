package seed

import (
	"tugas_akhir_example/internal/daos"
)

var ProdukSeed = []daos.Produk{
	{
		NamaProduk:    "ProdukA",
		Slug:          "produk-a",
		HargaReseller: "50000",
		HargaKonsumen: "75000",
		Stok:          10,
		Deskripsi:     "Suatu deskripsi yang menjelaskan produk",
		IdToko:        1,
		IdCategory:    1,
	},
	{
		NamaProduk:    "ProdukB",
		Slug:          "produk-b",
		HargaReseller: "75000",
		HargaKonsumen: "100000",
		Stok:          25,
		Deskripsi:     "Suatu deskripsi yang menjelaskan produk",
		IdToko:        2,
		IdCategory:    1,
	},
	{
		NamaProduk:    "ProdukC",
		Slug:          "produk-c",
		HargaReseller: "10000",
		HargaKonsumen: "20000",
		Stok:          5,
		Deskripsi:     "Suatu deskripsi yang menjelaskan produk",
		IdToko:        3,
		IdCategory:    2,
	},
	{
		NamaProduk:    "ProdukD",
		Slug:          "produk-d",
		HargaReseller: "5000",
		HargaKonsumen: "6000",
		Stok:          1,
		Deskripsi:     "Suatu deskripsi yang menjelaskan produk",
		IdToko:        4,
		IdCategory:    3,
	},
	{
		NamaProduk:    "ProdukE",
		Slug:          "produk-e",
		HargaReseller: "15000",
		HargaKonsumen: "17500",
		Stok:          1,
		Deskripsi:     "Suatu deskripsi yang menjelaskan produk",
		IdToko:        5,
		IdCategory:    5,
	},
	{
		NamaProduk:    "ProdukF",
		Slug:          "produk-f",
		HargaReseller: "25000",
		HargaKonsumen: "30000",
		Stok:          10,
		Deskripsi:     "Suatu deskripsi yang menjelaskan produk",
		IdToko:        2,
		IdCategory:    4,
	},
}
