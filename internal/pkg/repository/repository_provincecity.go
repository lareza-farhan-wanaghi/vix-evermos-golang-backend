package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"tugas_akhir_example/internal/pkg/dto"
)

const provinceCityListProvinceAPI = "https://emsifa.github.io/api-wilayah-indonesia/api/provinces.json"
const provinceCityListCityAPI = "https://emsifa.github.io/api-wilayah-indonesia/api/regencies/%s.json"
const provinceCityDetailProvinceAPI = "https://emsifa.github.io/api-wilayah-indonesia/api/province/%s.json"
const provinceCityDetailCityAPI = "https://emsifa.github.io/api-wilayah-indonesia/api/regency/%s.json"

type ProvinceCityRepository interface {
	GetAllProvinces(ctx context.Context, limit, offset int, search string) (res []*dto.ProvinceResp, err error)
	GetAllCities(ctx context.Context, provId string) (res []*dto.CityResp, err error)
	GetProvinceById(ctx context.Context, provId string) (res *dto.ProvinceResp, err error)
	GetCityById(ctx context.Context, cityId string) (res *dto.CityResp, err error)
}

type ProvinceCityRepositoryImpl struct {
}

// NewProvinceCityRepository returns the repository for the provincecity group path
func NewProvinceCityRepository() ProvinceCityRepository {
	return &ProvinceCityRepositoryImpl{}
}

// GetAllProvinces returns all province data from the external API
func (alr *ProvinceCityRepositoryImpl) GetAllProvinces(ctx context.Context, limit, offset int, search string) (res []*dto.ProvinceResp, err error) {
	resp, err := http.Get(provinceCityListProvinceAPI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	println(len(res))
	search = strings.ToLower(search)
	tmpres := []*dto.ProvinceResp{}
	for _, v := range res {
		if strings.Contains(strings.ToLower(v.Name), search) {
			tmpres = append(tmpres, v)
		}
	}
	res = tmpres
	println(len(res))

	if offset >= len(res) {
		res = nil
	} else {
		endIndex := offset + limit
		if endIndex > len(res) {
			endIndex = len(res)
		}
		res = res[offset:endIndex]
	}

	return res, nil
}

// GetAllCities returns all city data from the external API
func (alr *ProvinceCityRepositoryImpl) GetAllCities(ctx context.Context, provId string) (res []*dto.CityResp, err error) {
	resp, err := http.Get(fmt.Sprintf(provinceCityListCityAPI, provId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetProvinceById returns province data having the id from the external API
func (alr *ProvinceCityRepositoryImpl) GetProvinceById(ctx context.Context, provId string) (res *dto.ProvinceResp, err error) {
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

// GetCityById returns city data having the id from the external API
func (alr *ProvinceCityRepositoryImpl) GetCityById(ctx context.Context, cityId string) (res *dto.CityResp, err error) {
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
