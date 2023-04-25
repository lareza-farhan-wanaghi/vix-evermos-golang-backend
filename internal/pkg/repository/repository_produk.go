package repository

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type ProdukRepository interface {
	GetAllProduks(ctx context.Context, filter *daos.FilterProduk) (res []*daos.Produk, err error)
	GetProdukById(ctx context.Context, id string) (res *daos.Produk, err error)
	GetUserById(ctx context.Context, id string) (res *daos.User, err error)
	CreateProduk(ctx context.Context, data *daos.Produk) (res uint, err error)
	CreateFotoProduk(ctx context.Context, data *daos.FotoProduk) (res uint, err error)
	UpdateProduk(ctx context.Context, prevData *daos.Produk, data *daos.Produk) (err error)
	DeleteProduk(ctx context.Context, data *daos.Produk) (err error)
	DeleteFotoProduk(ctx context.Context, data *daos.FotoProduk) (err error)
}

type ProdukRepositoryImpl struct {
	db *gorm.DB
}

// NewProdukRepository returns the repository for the produk group path
func NewProdukRepository(db *gorm.DB) ProdukRepository {
	return &ProdukRepositoryImpl{
		db: db,
	}
}

// GetAllProduks returns all produk data from the produk table
func (alr *ProdukRepositoryImpl) GetAllProduks(ctx context.Context, filter *daos.FilterProduk) (res []*daos.Produk, err error) {
	tx := alr.db.Where("nama_produk like ?", fmt.Sprintf("%%%s%%", filter.NamaProduk))
	if filter.MaxHarga != 0 && filter.MaxHarga >= filter.MinHarga {
		tx = tx.Where("harga_konsumen BETWEEN ? AND ?", filter.MinHarga, filter.MaxHarga)
		tx = tx.Where("harga_reseller BETWEEN ? AND ?", filter.MinHarga, filter.MaxHarga)
	}

	if filter.CategoryId > 0 {
		tx = tx.Where("id_category = ?", filter.CategoryId)
	}

	if filter.TokoId > 0 {
		tx = tx.Where("id_Toko = ?", filter.TokoId)
	}
	tx = tx.WithContext(ctx).Limit(filter.Limit).Offset(filter.Offset)
	tx = tx.Model(daos.Produk{}).Preload("FotoProduks").Preload("Toko").Preload("Category")
	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetProdukById returns produk data having the id from the produk table
func (alr *ProdukRepositoryImpl) GetProdukById(ctx context.Context, id string) (res *daos.Produk, err error) {
	res = &daos.Produk{}
	if err := alr.db.WithContext(ctx).Model(daos.Produk{}).Preload("FotoProduks").Preload("Toko").Preload("Category").Where("id = ? ", id).First(res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// GetUserById returns user data having the id from the user table
func (alr *ProdukRepositoryImpl) GetUserById(ctx context.Context, id string) (res *daos.User, err error) {
	res = &daos.User{}
	if err := alr.db.WithContext(ctx).Where("id = ? ", id).Model(&daos.User{}).Preload("Toko").First(res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// CreateProduk inserts the produk data to the produk table
func (alr *ProdukRepositoryImpl) CreateProduk(ctx context.Context, data *daos.Produk) (res uint, err error) {
	result := alr.db.Create(data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

// CreateFotoProduk inserts the fotoproduk data to the fotoproduk table
func (alr *ProdukRepositoryImpl) CreateFotoProduk(ctx context.Context, data *daos.FotoProduk) (res uint, err error) {
	result := alr.db.WithContext(ctx).Create(data)
	if result.Error != nil {
		return 0, result.Error
	}

	return data.ID, nil
}

// UpdateProduk updates produk data on the produk table
func (alr *ProdukRepositoryImpl) UpdateProduk(ctx context.Context, prevData *daos.Produk, data *daos.Produk) (err error) {
	if err := alr.db.WithContext(ctx).Where("id = ?", prevData.ID).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// DeleteProduk deletes produk data having the id on the produk table
func (alr *ProdukRepositoryImpl) DeleteProduk(ctx context.Context, data *daos.Produk) (err error) {
	if err := alr.db.WithContext(ctx).Delete(data).Error; err != nil {
		return err
	}

	return nil
}

// DeleteFotoProduk deletes fotoproduk data having the id on the fotoproduk table
func (alr *ProdukRepositoryImpl) DeleteFotoProduk(ctx context.Context, data *daos.FotoProduk) (err error) {
	if err := alr.db.WithContext(ctx).Delete(data).Error; err != nil {
		return err
	}

	return nil
}
