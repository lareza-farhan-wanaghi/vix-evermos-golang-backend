package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type ProvinceCityController interface {
	GetAllProvinces(ctx *fiber.Ctx) error
	GetAllCities(ctx *fiber.Ctx) error
	GetProvinceById(ctx *fiber.Ctx) error
	GetCityById(ctx *fiber.Ctx) error
}

type ProvinceCityControllerImpl struct {
	provincecityusecase usecase.ProvinceCityUseCase
}

// NewProvinceCityController returns the controller for the provincecity group path
func NewProvinceCityController(provincecityusecase usecase.ProvinceCityUseCase) ProvinceCityController {
	return &ProvinceCityControllerImpl{
		provincecityusecase: provincecityusecase,
	}
}

// GetAllProvinces handles the delivery logic to retrieve all province data
func (uc *ProvinceCityControllerImpl) GetAllProvinces(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := &dto.ProvinceFilter{}
	err := ctx.QueryParser(filter)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, err.Error())
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	res, customErr := uc.provincecityusecase.GetAllProvinces(c, filter)
	if customErr != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, customErr.Err.Error())
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: customErr.Code,
			Errors:     []string{customErr.Err.Error()},
		})
	}

	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// GetAllCities handles the delivery logic to retrieve all city data
func (uc *ProvinceCityControllerImpl) GetAllCities(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, customErr := uc.provincecityusecase.GetAllCities(c, ctx.Params("prov_id"))
	if customErr != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, customErr.Err.Error())
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: customErr.Code,
			Errors:     []string{customErr.Err.Error()},
		})
	}

	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// GetProvinceById handles the delivery logic to retrieve province data having the id
func (uc *ProvinceCityControllerImpl) GetProvinceById(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, customErr := uc.provincecityusecase.GetProvinceById(c, ctx.Params("prov_id"))
	if customErr != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, customErr.Err.Error())
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: customErr.Code,
			Errors:     []string{customErr.Err.Error()},
		})
	}

	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// GetCityById handles the delivery logic to retrieve city data having the id
func (uc *ProvinceCityControllerImpl) GetCityById(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, customErr := uc.provincecityusecase.GetCityById(c, ctx.Params("city_id"))
	if customErr != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, customErr.Err.Error())
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: customErr.Code,
			Errors:     []string{customErr.Err.Error()},
		})
	}

	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
