package repository

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type TokoRepository interface {
	GetAllTokos(ctx context.Context, queries daos.FilterToko) (res []*daos.Toko, err error)
	GetTokoById(ctx context.Context, id string) (res *daos.Toko, err error)
	GetTokoByUserID(ctx context.Context, userId string) (res *daos.Toko, err error)
	UpdateToko(ctx context.Context, prevData *daos.Toko, data *daos.Toko) (err error)
}

type TokoRepositoryImpl struct {
	db *gorm.DB
}

// NewTokoRepository returns the repository for the toko group path
func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &TokoRepositoryImpl{
		db: db,
	}
}

// GetAllTokos returns all toko data from the toko table
func (alr *TokoRepositoryImpl) GetAllTokos(ctx context.Context, queries daos.FilterToko) (res []*daos.Toko, err error) {
	if err := alr.db.Where("nama_toko like ?", fmt.Sprintf("%%%s%%", queries.NamaToko)).WithContext(ctx).Limit(queries.Limit).Offset(queries.Offset).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// GetTokoById returns toko data having the id from the toko table
func (alr *TokoRepositoryImpl) GetTokoById(ctx context.Context, id string) (res *daos.Toko, err error) {
	if err := alr.db.Where("id = ? ", id).First(&res).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

// GetTokoByUserID returns toko data having the userid from the toko table
func (alr *TokoRepositoryImpl) GetTokoByUserID(ctx context.Context, userId string) (res *daos.Toko, err error) {
	if err := alr.db.Where("id_user = ? ", userId).First(&res).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

// UpdateToko updates toko data on the toko table
func (alr *TokoRepositoryImpl) UpdateToko(ctx context.Context, prevData *daos.Toko, data *daos.Toko) (err error) {
	if err := alr.db.WithContext(ctx).Where("id = ?", prevData.ID).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
