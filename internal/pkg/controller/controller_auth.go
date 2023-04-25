package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	RegisterUsers(ctx *fiber.Ctx) error
	LoginUsers(ctx *fiber.Ctx) error
}

type AuthControllerImpl struct {
	authusecase usecase.AuthUseCase
}

// NewAuthController returns the controller for the auth group path
func NewAuthController(authusecase usecase.AuthUseCase) AuthController {
	return &AuthControllerImpl{
		authusecase: authusecase,
	}
}

// RegisterUsers handles the delivery logic to register the user
func (uc *AuthControllerImpl) RegisterUsers(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := new(dto.AuthReqRegister)
	if err := ctx.BodyParser(data); err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	customErr := uc.authusecase.RegisterUser(c, *data)
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
		Data:       "Register Succeed",
	})
}

// LoginUsers handles the delivery logic to login the user
func (uc *AuthControllerImpl) LoginUsers(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := new(dto.AuthReqLogin)
	if err := ctx.BodyParser(data); err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	loginResp, customErr := uc.authusecase.LoginUser(c, *data)
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
		Data:       loginResp,
	})
}
