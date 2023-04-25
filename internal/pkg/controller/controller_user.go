package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetMyAlamats(ctx *fiber.Ctx) error
	GetAlamatById(ctx *fiber.Ctx) error
	GetMyProfile(ctx *fiber.Ctx) error
	CreateAlamat(ctx *fiber.Ctx) error
	UpdateAlamatById(ctx *fiber.Ctx) error
	UpdateProfile(ctx *fiber.Ctx) error
	DeleteAlamatById(ctx *fiber.Ctx) error
}

type UserControllerImpl struct {
	userusecase usecase.UserUseCase
}

// NewUserController returns the controller for the user group path
func NewUserController(userusecase usecase.UserUseCase) UserController {
	return &UserControllerImpl{
		userusecase: userusecase,
	}
}

// GetMyAlamats handles the delivery logic to retrieve alamat data of the current user
func (uc *UserControllerImpl) GetMyAlamats(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := &dto.AlamatFilter{}
	err := ctx.QueryParser(filter)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, err.Error())
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	res, customErr := uc.userusecase.GetMyAlamats(c, ctx.Get("token"), filter)
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

// GetAlamatById handles the delivery logic to retrieve alamat data having the id
func (uc *UserControllerImpl) GetAlamatById(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, customErr := uc.userusecase.GetAlamatById(c, ctx.Params("id"))
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

// GetMyProfile handles the delivery logic to retrieve user data of the current user
func (uc *UserControllerImpl) GetMyProfile(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, customErr := uc.userusecase.GetMyProfile(c, ctx.Get("token"))
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

// CreateAlamat handles the delivery logic to insert the user data
func (uc *UserControllerImpl) CreateAlamat(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := &dto.AlamatCreateReq{}
	err := ctx.BodyParser(data)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, err.Error())
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	res, customErr := uc.userusecase.CreateAlamat(c, data, ctx.Get("token"))
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
		StatusCode: fiber.StatusCreated,
		Data:       res,
	})
}

// UpdateAlamatById handles the delivery logic to update alamat data having the id
func (uc *UserControllerImpl) UpdateAlamatById(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := &dto.AlamatUpdateReq{}
	err := ctx.BodyParser(data)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, err.Error())
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	customErr := uc.userusecase.UpdateAlamatById(c, ctx.Params("id"), data)
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

// UpdateProfile handles the delivery logic to update user data of the current user
func (uc *UserControllerImpl) UpdateProfile(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := &dto.UserUpdateReq{}
	err := ctx.BodyParser(data)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, err.Error())
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	customErr := uc.userusecase.UpdateProfile(c, ctx.Get("token"), data)
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

// DeleteAlamatById handles the delivery logic to delete alamat data having the id
func (uc *UserControllerImpl) DeleteAlamatById(ctx *fiber.Ctx) error {
	c := ctx.Context()

	customErr := uc.userusecase.DeleteAlamatByID(c, ctx.Params("id"))
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
