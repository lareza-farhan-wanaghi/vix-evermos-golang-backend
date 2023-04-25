package utils

import (
	"fmt"
	"strconv"

	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
)

// @TODO : make middleware like Auth

// TokoAuthMiddleware auths the user by comparing the userid contained in the jwt token and the userid of the toko data
func TokoAuthMiddleware(tokoRepository repository.TokoRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		claims, err := GetJWTClaims(ctx.Get("token"))
		if err != nil {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{err.Error()},
			})
		}

		tokoId := ctx.Params("id_toko")
		if tokoId == "" {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, "id_toko params required ")
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{"Bad request"},
			})
		}

		resRepo, err := tokoRepository.GetTokoById(ctx.Context(), tokoId)
		if err != nil {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{err.Error()},
			})
		}

		if strconv.Itoa(int(resRepo.IdUser)) != claims.UserId {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", "unauthorized"))
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{"You are unauthorized"},
			})
		}
		return ctx.Next()
	}
}

// ProdukAuthMiddleware auths the user by comparing the userid contained in the jwt token and the userid of the toko data having the produk
func ProdukAuthMiddleware(produkRepository repository.ProdukRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		claims, err := GetJWTClaims(ctx.Get("token"))
		if err != nil {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{err.Error()},
			})
		}

		produkId := ctx.Params("id")
		if produkId == "" {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, "id params required ")
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{"Bad request"},
			})
		}

		resRepo, err := produkRepository.GetProdukById(ctx.Context(), produkId)
		if err != nil {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{err.Error()},
			})
		}

		if strconv.Itoa(int(resRepo.Toko.IdUser)) != claims.UserId {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", "unauthorized"))
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{"You are unauthorized"},
			})
		}
		return ctx.Next()
	}
}

// AlamatAuthMiddleware auths the user by comparing the userid contained in the jwt token and the userid of the alamat data
func AlamatAuthMiddleware(userRepository repository.UserRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		claims, err := GetJWTClaims(ctx.Get("token"))
		if err != nil {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{err.Error()},
			})
		}

		alamatId := ctx.Params("id")
		if alamatId == "" {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, "id params required ")
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{"Bad request"},
			})
		}

		resRepo, err := userRepository.GetAlamatById(ctx.Context(), alamatId)
		if err != nil {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{err.Error()},
			})
		}

		if strconv.Itoa(int(resRepo.IdUser)) != claims.UserId {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", "unauthorized"))
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{"You are unauthorized"},
			})
		}
		return ctx.Next()
	}
}

// CategoryAuthMiddleware auths the user by checking whether the user specified in the jwt token is an admin
func CategoryAuthMiddleware(categoryRepository repository.CategoryRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		claims, err := GetJWTClaims(ctx.Get("token"))
		if err != nil {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{err.Error()},
			})
		}

		resRepo, err := categoryRepository.GetUserById(ctx.Context(), claims.UserId)
		if err != nil {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{err.Error()},
			})
		}

		if !resRepo.IsAdmin {
			helper.Logger(GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", "unauthorized"))
			return helper.ResponseWithJSON(&helper.JSONRespArgs{
				Ctx:        ctx,
				StatusCode: fiber.StatusBadRequest,
				Errors:     []string{"You are unauthorized"},
			})
		}
		return ctx.Next()
	}
}
