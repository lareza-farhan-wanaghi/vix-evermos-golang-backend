package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type TokoController interface {
	GetAllToko(ctx *fiber.Ctx) error
	GetTokoById(ctx *fiber.Ctx) error
	GetMyToko(ctx *fiber.Ctx) error
	UpdateTokoByID(ctx *fiber.Ctx) error
}

type TokoControllerImpl struct {
	tokousecase usecase.TokoUseCase
}

// NewTokoController returns the controller for the toko group path
func NewTokoController(tokousecase usecase.TokoUseCase) TokoController {
	return &TokoControllerImpl{
		tokousecase: tokousecase,
	}
}

// GetAllToko handles the delivery logic to retrieve all toko data
func (uc *TokoControllerImpl) GetAllToko(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := new(dto.TokoFilter)
	if err := ctx.QueryParser(filter); err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	res, err := uc.tokousecase.GetAllTokos(c, filter)

	if err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: err.Code,
			Errors:     []string{err.Err.Error()},
		})
	}

	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// GetTokoById handles the delivery logic to retrieve toko data having the id
func (uc *TokoControllerImpl) GetTokoById(ctx *fiber.Ctx) error {
	c := ctx.Context()

	toko, err := uc.tokousecase.GetTokoById(c, ctx.Params("id_toko"), ctx.Get("token"))
	if err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: err.Code,
			Errors:     []string{err.Err.Error()},
		})
	}

	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusOK,
		Data:       toko,
	})
}

// GetMyToko handles the delivery logic to retrieve toko data of the current user
func (uc *TokoControllerImpl) GetMyToko(ctx *fiber.Ctx) error {
	c := ctx.Context()

	toko, err := uc.tokousecase.GetMyToko(c, ctx.Get("token"))
	if err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: err.Code,
			Errors:     []string{err.Err.Error()},
		})
	}

	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusOK,
		Data:       toko,
	})
}

// UpdateTokoByID handles the delivery logic to update toko data having the id
func (uc *TokoControllerImpl) UpdateTokoByID(ctx *fiber.Ctx) error {
	c := ctx.Context()

	form, err := c.MultipartForm()
	if err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadGateway,
			Errors:     []string{err.Error()},
		})
	}

	customErr := uc.tokousecase.UpdateTokoByID(c, ctx.Params("id_toko"), utils.GetMultiFormFirstFile(form, "photo"), &dto.TokoUpdateReq{
		NamaToko: utils.GetMultiFormFirstValue(form, "nama_toko"),
	})
	if customErr != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, customErr.Err.Error())
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{customErr.Err.Error()},
		})
	}

	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusOK,
		Data:       "Update toko succeed",
	})
}
