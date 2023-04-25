package usecase

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthUseCase interface {
	LoginUser(ctx context.Context, data dto.AuthReqLogin) (res *dto.LoginResp, err *helper.ErrorStruct)
	RegisterUser(ctx context.Context, data dto.AuthReqRegister) (err *helper.ErrorStruct)
}

type AuthUseCaseImpl struct {
	authRepository repository.AuthRepository
	jwtSecret      string
}

// NewAuthUseCase returns the usecase for the auth group path
func NewAuthUseCase(authRepository repository.AuthRepository, jwtSecret string) AuthUseCase {
	return &AuthUseCaseImpl{
		authRepository: authRepository,
		jwtSecret:      jwtSecret,
	}
}

// LoginUser handles the business logic to log in the user
func (alc *AuthUseCaseImpl) LoginUser(ctx context.Context, data dto.AuthReqLogin) (res *dto.LoginResp, customErr *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errValidate.Error()))
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, err := alc.authRepository.GetUserByNotelp(ctx, data.Notelp)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("no data user")
		}
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	err = utils.ValidatePassword(resRepo.KataSandi, data.KataSandi)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			err = errors.New("no telp atau kata sandi salah")
		}
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	token, err := utils.GenerateNewJWT(&utils.Claims{
		UserId: strconv.Itoa(int(resRepo.ID)),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	})
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	provinceData, err := alc.authRepository.GetProvinceById(ctx, resRepo.IdProvinsi)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	cityData, err := alc.authRepository.GetCityById(ctx, resRepo.IdKota)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	res = utils.UserToLoginResp(resRepo)
	res.IdProvinsi = provinceData
	res.IdKota = cityData
	res.Token = token

	return res, nil
}

// RegisterUser handles the business logic to register the user
func (alc *AuthUseCaseImpl) RegisterUser(ctx context.Context, data dto.AuthReqRegister) (err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errValidate.Error()))
		return &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	tanggalLahir, parsingErr := utils.StringToDate(data.TanggalLahir)
	if parsingErr != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", parsingErr.Error()))
		return &helper.ErrorStruct{
			Err:  parsingErr,
			Code: fiber.StatusBadRequest,
		}
	}

	katasandiHash, hashingErr := utils.HashPassword(data.KataSandi)
	if hashingErr != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", hashingErr.Error()))
		return &helper.ErrorStruct{
			Err:  hashingErr,
			Code: fiber.StatusBadRequest,
		}
	}

	userID, errRepo := alc.authRepository.CreateUser(ctx, &daos.User{
		Nama:         data.Nama,
		KataSandi:    katasandiHash,
		Notelp:       data.Notelp,
		TanggalLahir: tanggalLahir,
		JenisKelamin: data.JenisKelamin,
		Tentang:      data.Tentang,
		Pekerjaan:    data.Pekerjaan,
		Email:        data.Email,
		IdProvinsi:   data.IdProvinsi,
		IdKota:       data.IdKota,
		IsAdmin:      false,
	})

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	_, errRepo = alc.authRepository.CreateToko(ctx, &daos.Toko{
		IdUser:   userID,
		NamaToko: fmt.Sprintf("TokoUser%d", userID),
	})

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return nil
}
