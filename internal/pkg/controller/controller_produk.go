package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type ProdukController interface {
	GetAllProduks(ctx *fiber.Ctx) error
	GetProdukById(ctx *fiber.Ctx) error
	CreateProduk(ctx *fiber.Ctx) error
	UpdateProdukById(ctx *fiber.Ctx) error
	DeleteProdukById(ctx *fiber.Ctx) error
}

type ProdukControllerImpl struct {
	produkusecase usecase.ProdukUseCase
}

// NewProdukController return the controller for the produk group path
func NewProdukController(produkusecase usecase.ProdukUseCase) ProdukController {
	return &ProdukControllerImpl{
		produkusecase: produkusecase,
	}
}

// GetAllProduks handles the delivery logic to retrieve all produk data
func (uc *ProdukControllerImpl) GetAllProduks(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := new(dto.ProdukFilter)
	if err := ctx.QueryParser(filter); err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	res, customErr := uc.produkusecase.GetAllProduks(c, filter)
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

// GetProdukById handles the delivery logic to retrieve produk data having the id
func (uc *ProdukControllerImpl) GetProdukById(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, err := uc.produkusecase.GetProdukById(c, ctx.Params("id"))
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

// CreateProduk handles the delivery logic to insert the produk data
func (uc *ProdukControllerImpl) CreateProduk(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := &dto.ProdukCreateReq{}
	if err := ctx.BodyParser(data); err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadGateway,
			Errors:     []string{err.Error()},
		})
	}

	res, customErr := uc.produkusecase.CreateProduk(c, data, ctx.Get("token"), form.File["photos"])
	if customErr != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: customErr.Code,
			Errors:     []string{customErr.Err.Error()},
		})
	}

	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusCreated,
		Data:       res,
	})
}

// UpdateProdukById handles the delivery logic to update produk data having the id
func (uc *ProdukControllerImpl) UpdateProdukById(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := &dto.ProdukUpdateReq{}
	if err := ctx.BodyParser(data); err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadGateway,
			Errors:     []string{err.Error()},
		})
	}

	customErr := uc.produkusecase.UpdateProdukByID(c, data, ctx.Params("id"), form.File["photos"])
	if customErr != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: customErr.Code,
			Errors:     []string{customErr.Err.Error()},
		})
	}

	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusOK,
		Data:       "Update Succeed",
	})
}

// DeleteProdukById handles the delivery logic to delete produk data having the id
func (uc *ProdukControllerImpl) DeleteProdukById(ctx *fiber.Ctx) error {
	c := ctx.Context()

	customErr := uc.produkusecase.DeleteProdukByID(c, ctx.Params("id"))
	if customErr != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: customErr.Code,
			Errors:     []string{customErr.Err.Error()},
		})
	}

	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusOK,
		Data:       "Delete Succeed",
	})
}
