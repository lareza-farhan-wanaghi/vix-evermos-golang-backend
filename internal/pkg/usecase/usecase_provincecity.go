package usecase

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type ProvinceCityUseCase interface {
	GetAllProvinces(ctx context.Context, filter *dto.ProvinceFilter) (res []*dto.ProvinceResp, err *helper.ErrorStruct)
	GetAllCities(ctx context.Context, provId string) (res []*dto.CityResp, err *helper.ErrorStruct)
	GetProvinceById(ctx context.Context, provId string) (res *dto.ProvinceResp, err *helper.ErrorStruct)
	GetCityById(ctx context.Context, cityId string) (res *dto.CityResp, err *helper.ErrorStruct)
}

type ProvinceCityUseCaseImpl struct {
	provinceCityRepository repository.ProvinceCityRepository
}

// NewProvinceCityUseCase returns the usecase for the provincecity group path
func NewProvinceCityUseCase(provinceCityRepository repository.ProvinceCityRepository) ProvinceCityUseCase {
	return &ProvinceCityUseCaseImpl{
		provinceCityRepository: provinceCityRepository,
	}
}

// GetAllProvinces handles the business logic to retrieve all province data
func (alc *ProvinceCityUseCaseImpl) GetAllProvinces(ctx context.Context, filter *dto.ProvinceFilter) (res []*dto.ProvinceResp, err *helper.ErrorStruct) {
	if filter.Limit < 1 {
		filter.Limit = 10
	}

	if filter.Page < 1 {
		filter.Page = 1
	}

	res, errRepo := alc.provinceCityRepository.GetAllProvinces(ctx, filter.Limit, (filter.Page-1)*filter.Limit, filter.Search)

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return res, nil
}

// GetAllCities handles the business logic to retrieve all city data
func (alc *ProvinceCityUseCaseImpl) GetAllCities(ctx context.Context, provId string) (res []*dto.CityResp, err *helper.ErrorStruct) {
	res, errRepo := alc.provinceCityRepository.GetAllCities(ctx, provId)

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return res, nil
}

// GetProvinceById handles the business logic to retrieve province data having the id
func (alc *ProvinceCityUseCaseImpl) GetProvinceById(ctx context.Context, provId string) (res *dto.ProvinceResp, err *helper.ErrorStruct) {
	res, errRepo := alc.provinceCityRepository.GetProvinceById(ctx, provId)

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return res, nil
}

// GetCityById handles the business logic to retrieve city data having the id
func (alc *ProvinceCityUseCaseImpl) GetCityById(ctx context.Context, cityId string) (res *dto.CityResp, err *helper.ErrorStruct) {
	res, errRepo := alc.provinceCityRepository.GetCityById(ctx, cityId)

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return res, nil
}
