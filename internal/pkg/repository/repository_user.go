package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/pkg/dto"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAlamatsByUserId(ctx context.Context, userId string, filter *daos.FilterAlamat) (res []*daos.Alamat, err error)
	GetAlamatById(ctx context.Context, id string) (res *daos.Alamat, err error)
	GetUserById(ctx context.Context, id string) (res *daos.User, err error)
	GetCityById(ctx context.Context, cityId string) (res *dto.CityResp, err error)
	GetProvinceById(ctx context.Context, provId string) (res *dto.ProvinceResp, err error)
	CreateAlamat(ctx context.Context, data *daos.Alamat) (res uint, err error)
	UpdateAlamatByID(ctx context.Context, id string, data *daos.Alamat) (err error)
	UpdateUserById(ctx context.Context, id string, data *daos.User) (err error)
	DeleteAlamatById(ctx context.Context, id string) (err error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository returns the repository for the user group path
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

// GetAlamatsByUserId returns alamat data having the userid from the alamat table
func (alr *UserRepositoryImpl) GetAlamatsByUserId(ctx context.Context, userId string, filter *daos.FilterAlamat) (res []*daos.Alamat, err error) {
	if err := alr.db.WithContext(ctx).Where("id_user = ?", userId).Where("judul_alamat like ?", fmt.Sprintf("%%%s%%", filter.JudulAlamat)).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// GetAlamatById returns alamat data having the id from the alamat table
func (alr *UserRepositoryImpl) GetAlamatById(ctx context.Context, id string) (res *daos.Alamat, err error) {
	res = &daos.Alamat{}
	if err := alr.db.WithContext(ctx).Where("id = ? ", id).First(res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (alr *UserRepositoryImpl) GetUserById(ctx context.Context, id string) (res *daos.User, err error) {
	res = &daos.User{}
	if err := alr.db.WithContext(ctx).Model(&daos.User{}).Where("id = ?", id).Preload("Toko").Preload("Alamats").First(res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (alr *UserRepositoryImpl) GetCityById(ctx context.Context, cityId string) (res *dto.CityResp, err error) {
	resp, err := http.Get(fmt.Sprintf(provinceCityDetailCityAPI, cityId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unable to retrieve city data")
	}

	res = &dto.CityResp{}
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (alr *UserRepositoryImpl) GetProvinceById(ctx context.Context, provId string) (res *dto.ProvinceResp, err error) {
	resp, err := http.Get(fmt.Sprintf(provinceCityDetailProvinceAPI, provId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unable to retrieve province data")
	}

	res = &dto.ProvinceResp{}
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateAlamat inserts the alamat data to the alamat table
func (alr *UserRepositoryImpl) CreateAlamat(ctx context.Context, data *daos.Alamat) (res uint, err error) {
	result := alr.db.WithContext(ctx).Create(data)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

// UpdateAlamatByID updates alamat data having the id on the alamat table
func (alr *UserRepositoryImpl) UpdateAlamatByID(ctx context.Context, id string, data *daos.Alamat) (err error) {
	if err = alr.db.WithContext(ctx).Where("id = ? ", id).First(&daos.Alamat{}).Error; err != nil {
		return gorm.ErrRecordNotFound
	}

	if err := alr.db.WithContext(ctx).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// UpdateUserById updates user data having the id on the user table
func (alr *UserRepositoryImpl) UpdateUserById(ctx context.Context, id string, data *daos.User) (err error) {
	if err = alr.db.WithContext(ctx).Where("id = ? ", id).First(&daos.User{}).Error; err != nil {
		return gorm.ErrRecordNotFound
	}

	if err := alr.db.WithContext(ctx).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// DeleteAlamatById deletes alamat data having the id on the alamat table
func (alr *UserRepositoryImpl) DeleteAlamatById(ctx context.Context, id string) (err error) {
	if err = alr.db.WithContext(ctx).Where("id = ? ", id).First(&daos.Alamat{}).Error; err != nil {
		return gorm.ErrRecordNotFound
	}

	if err := alr.db.WithContext(ctx).Where("id = ?", id).Delete(&daos.Alamat{}).Error; err != nil {
		return err
	}

	return nil
}
