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

type AuthRepository interface {
	GetUserByNotelp(ctx context.Context, nama string) (res *daos.User, err error)
	GetProvinceById(ctx context.Context, provId string) (res *dto.ProvinceResp, err error)
	GetCityById(ctx context.Context, cityId string) (res *dto.CityResp, err error)
	CreateUser(ctx context.Context, data *daos.User) (res uint, err error)
	CreateToko(ctx context.Context, data *daos.Toko) (res uint, err error)
}

type AuthRepositoryImpl struct {
	db *gorm.DB
}

// NewAuthRepository returns the repository for the auth group path
func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		db: db,
	}
}

// GetUserByNotelp returns user data having the notelp from the user table
func (alr *AuthRepositoryImpl) GetUserByNotelp(ctx context.Context, notelp string) (res *daos.User, err error) {
	if err := alr.db.Where("notelp = ? ", notelp).First(&res).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

// GetProvinceById returns province data having the id from the specified external API
func (alr *AuthRepositoryImpl) GetProvinceById(ctx context.Context, provId string) (res *dto.ProvinceResp, err error) {
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

// GetCityById returns city data having the id from the specified external API
func (alr *AuthRepositoryImpl) GetCityById(ctx context.Context, cityId string) (res *dto.CityResp, err error) {
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

// CreateUser inserts the user data to the user table
func (alr *AuthRepositoryImpl) CreateUser(ctx context.Context, data *daos.User) (res uint, err error) {
	result := alr.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

// CreateToko inserts the toko data to the toko table
func (alr *AuthRepositoryImpl) CreateToko(ctx context.Context, data *daos.Toko) (res uint, err error) {
	result := alr.db.Create(data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return res, nil
}
