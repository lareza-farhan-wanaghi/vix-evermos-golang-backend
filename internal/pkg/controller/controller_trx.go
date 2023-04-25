package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type TrxController interface {
	GetAllTrxs(ctx *fiber.Ctx) error
	GetTrxById(ctx *fiber.Ctx) error
	CreateTrx(ctx *fiber.Ctx) error
}

type TrxControllerImpl struct {
	trxusecase usecase.TrxUseCase
}

// NewTrxController returns the controller for the trx group path
func NewTrxController(trxusecase usecase.TrxUseCase) TrxController {
	return &TrxControllerImpl{
		trxusecase: trxusecase,
	}
}

// GetAllTrxs handles the delivery logic to retrieve all trx data
func (uc *TrxControllerImpl) GetAllTrxs(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := &dto.TrxFilter{}
	err := ctx.QueryParser(filter)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, err.Error())
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadGateway,
			Errors:     []string{err.Error()},
		})
	}

	res, customErr := uc.trxusecase.GetAllTrxs(c, filter)
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

// GetTrxById handles the delivery logic to retrieve trx data having the id
func (uc *TrxControllerImpl) GetTrxById(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, customErr := uc.trxusecase.GetTrxById(c, ctx.Params("id"))
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

// CreateTrx handles the delivery logic to insert the trx data
func (uc *TrxControllerImpl) CreateTrx(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := &dto.TrxCreateReq{}
	err := ctx.BodyParser(data)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, err.Error())
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	res, customErr := uc.trxusecase.CreateTrx(c, ctx.Get("token"), data)
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
