package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	GetAllCategories(ctx *fiber.Ctx) error
	GetCategoryById(ctx *fiber.Ctx) error
	CreateCategory(ctx *fiber.Ctx) error
	UpdateCategoryById(ctx *fiber.Ctx) error
	DeleteCategoryById(ctx *fiber.Ctx) error
}

type CategoryControllerImpl struct {
	categoryusecase usecase.CategoryUseCase
}

// NewCategoryController returns the controller for the category group path
func NewCategoryController(categoryusecase usecase.CategoryUseCase) CategoryController {
	return &CategoryControllerImpl{
		categoryusecase: categoryusecase,
	}
}

// GetAllCategories handles the delivery logic to retrieve all category data
func (uc *CategoryControllerImpl) GetAllCategories(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, customErr := uc.categoryusecase.GetAllCategories(c)
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

// GetCategoryById handles the delivery logic to retrieve category data having the id
func (uc *CategoryControllerImpl) GetCategoryById(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, customErr := uc.categoryusecase.GetCategoryById(c, ctx.Params("id"))
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

// CreateCategory handles the delivery logic to insert the category data
func (uc *CategoryControllerImpl) CreateCategory(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := &dto.CategoryCreateReq{}
	err := ctx.BodyParser(data)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, err.Error())
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	res, customErr := uc.categoryusecase.CreateCategory(c, data)
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

// UpdateCategoryById handles the delivery logic to update category data having the id
func (uc *CategoryControllerImpl) UpdateCategoryById(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := &dto.CategoryUpdateReq{}
	err := ctx.BodyParser(data)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, err.Error())
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	customErr := uc.categoryusecase.UpdateCategoryById(c, ctx.Params("id"), data)
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
		Data:       "Update succeed",
	})
}

// DeleteCategoryById handles the delivery logic to delete category data having the id
func (uc *CategoryControllerImpl) DeleteCategoryById(ctx *fiber.Ctx) error {
	c := ctx.Context()

	customErr := uc.categoryusecase.DeleteCategoryByID(c, ctx.Params("id"))
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
		Data:       "Delete succeed",
	})
}
