package seed

import (
	"time"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/utils"
)

var UserSeed = []daos.User{
	{
		Nama:         "UserA",
		KataSandi:    utils.UnsafeHashPassword("123456"),
		Notelp:       "08961231235",
		TanggalLahir: time.Now(),
		JenisKelamin: "L",
		Tentang:      "Merupakan seorang pelajar",
		Pekerjaan:    "Belajar",
		Email:        "reza@example.com",
		IdProvinsi:   "11",
		IdKota:       "1101",
		IsAdmin:      true,
	},
	{
		Nama:         "UserB",
		KataSandi:    utils.UnsafeHashPassword("123456"),
		Notelp:       "08922261231235",
		TanggalLahir: time.Now(),
		JenisKelamin: "L",
		Tentang:      "Merupakan seorang pelajar",
		Pekerjaan:    "Belajar",
		Email:        "testa@example.com",
		IdProvinsi:   "11",
		IdKota:       "1101",
		IsAdmin:      false,
	},
	{
		Nama:         "UserC",
		KataSandi:    utils.UnsafeHashPassword("123456"),
		Notelp:       "1234",
		TanggalLahir: time.Now(),
		JenisKelamin: "L",
		Tentang:      "Merupakan seorang pelajar",
		Pekerjaan:    "Belajar",
		Email:        "test2a@example.com",
		IdProvinsi:   "11",
		IdKota:       "1101",
		IsAdmin:      false,
	},
	{
		Nama:         "UserD",
		KataSandi:    utils.UnsafeHashPassword("123456"),
		Notelp:       "12342",
		TanggalLahir: time.Now(),
		JenisKelamin: "L",
		Tentang:      "Merupakan seorang pelajar",
		Pekerjaan:    "Belajar",
		Email:        "testca@example.com",
		IdProvinsi:   "11",
		IdKota:       "1101",
		IsAdmin:      false,
	},
	{
		Nama:         "UserE",
		KataSandi:    utils.UnsafeHashPassword("123456"),
		Notelp:       "123422",
		TanggalLahir: time.Now(),
		JenisKelamin: "L",
		Tentang:      "Merupakan seorang pelajar",
		Pekerjaan:    "Belajar",
		Email:        "testea@example.com",
		IdProvinsi:   "11",
		IdKota:       "1101",
		IsAdmin:      false,
	},
}
