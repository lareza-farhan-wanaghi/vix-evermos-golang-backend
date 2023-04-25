package repository

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type TrxRepository interface {
	GetAllTrxs(ctx context.Context, filter *daos.FilterTrx) (res []*daos.Trx, err error)
	GetTrxById(ctx context.Context, id string) (res *daos.Trx, err error)
	GetProdukById(ctx context.Context, id string) (res *daos.Produk, err error)
	GetAlamatById(ctx context.Context, id string) (res *daos.Alamat, err error)
	CreateTrx(ctx context.Context, data *daos.Trx) (res uint, err error)
}

type TrxRepositoryImpl struct {
	db *gorm.DB
}

// NewTrxRepository returns the repository for the trx group path
func NewTrxRepository(db *gorm.DB) TrxRepository {
	return &TrxRepositoryImpl{
		db: db,
	}
}

// GetAllTrxs returns all trx data from the trx table
func (alr *TrxRepositoryImpl) GetAllTrxs(ctx context.Context, filter *daos.FilterTrx) (res []*daos.Trx, err error) {
	tx := alr.db.WithContext(ctx).Model(&res).Limit(filter.Limit).Offset(filter.Offset)
	tx = tx.Preload("DetailTrxs").Preload("Alamat")
	tx = tx.Preload("DetailTrxs.LogProduk")
	tx = tx.Preload("DetailTrxs.LogProduk.Produk").Preload("DetailTrxs.LogProduk.Toko").Preload("DetailTrxs.LogProduk.Category")
	tx = tx.Preload("DetailTrxs.LogProduk.Produk.FotoProduks")
	tx = tx.Where("kode_invoice like ?", fmt.Sprintf("%%%s%%", filter.KodeInvoice))
	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// GetTrxById returns trx data having the id from the trx table
func (alr *TrxRepositoryImpl) GetTrxById(ctx context.Context, id string) (res *daos.Trx, err error) {
	tx := alr.db.WithContext(ctx).Model(&res)
	tx = tx.Preload("DetailTrxs").Preload("Alamat")
	tx = tx.Preload("DetailTrxs.LogProduk")
	tx = tx.Preload("DetailTrxs.LogProduk.Produk").Preload("DetailTrxs.LogProduk.Toko").Preload("DetailTrxs.LogProduk.Category")
	tx = tx.Preload("DetailTrxs.LogProduk.Produk.FotoProduks")
	if err := tx.Where("id = ?", id).First(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// GetProdukById returns produk data having the id from the produk table
func (alr *TrxRepositoryImpl) GetProdukById(ctx context.Context, id string) (res *daos.Produk, err error) {
	if err := alr.db.WithContext(ctx).Model(&res).Preload("FotoProduks").Where("id = ?", id).First(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// GetAlamatById returns alamat data having the id from the alamat table
func (alr *TrxRepositoryImpl) GetAlamatById(ctx context.Context, id string) (res *daos.Alamat, err error) {
	if err := alr.db.WithContext(ctx).Model(&res).Where("id = ?", id).First(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// CreateTrx inserts the trx data to the trx table
func (alr *TrxRepositoryImpl) CreateTrx(ctx context.Context, data *daos.Trx) (res uint, err error) {
	result := alr.db.WithContext(ctx).Create(data)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}
