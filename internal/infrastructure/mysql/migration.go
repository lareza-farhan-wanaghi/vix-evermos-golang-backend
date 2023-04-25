package mysql

import (
	"fmt"
	"reflect"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/infrastructure/mysql/seed"

	"gorm.io/gorm"
)

// RunMigration runs database migrations and seeds mock data to the database
func RunMigration(mysqlDB *gorm.DB) {
	err := mysqlDB.AutoMigrate(
		&daos.Category{},
		&daos.User{},
		&daos.Toko{},
		&daos.Alamat{},
		&daos.Produk{},
		&daos.LogProduk{},
		&daos.FotoProduk{},
		&daos.Trx{},
		&daos.DetailTrx{},
		&daos.Book{},
	)

	SeedData(mysqlDB,
		seed.CategorySeed,
		seed.UserSeed,
		seed.TokoSeed,
		seed.AlamatSeed,
		seed.ProdukSeed,
		seed.LogProdukSeed,
		seed.FotoProdukSeed,
		seed.TrxSeed,
		seed.DetailTrxSeed,
		seed.BookSeed,
	)

	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Failed Database Migrated : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Database Migrated")
}

// SeedData inserts the data to the database
func SeedData(mysqlDB *gorm.DB, seeds ...interface{}) {
	for _, seed := range seeds {
		var count int64
		firstData := reflect.ValueOf(seed).Index(0).Interface()
		if mysqlDB.Migrator().HasTable(firstData) {
			mysqlDB.Model(firstData).Count(&count)
			if count < 1 {
				mysqlDB.CreateInBatches(seed, reflect.ValueOf(seed).Len())
			}
		}
	}
}
