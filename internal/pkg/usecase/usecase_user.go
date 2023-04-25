package usecase

import (
	"context"
	"fmt"
	"strconv"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type UserUseCase interface {
	GetMyAlamats(ctx context.Context, token string, filter *dto.AlamatFilter) (res []*dto.AlamatResp, customErr *helper.ErrorStruct)
	GetAlamatById(ctx context.Context, id string) (res *dto.AlamatResp, customErr *helper.ErrorStruct)
	GetMyProfile(ctx context.Context, token string) (res *dto.UserResp, customErr *helper.ErrorStruct)
	CreateAlamat(ctx context.Context, data *dto.AlamatCreateReq, token string) (res uint, customErr *helper.ErrorStruct)
	UpdateAlamatById(ctx context.Context, id string, data *dto.AlamatUpdateReq) (customErr *helper.ErrorStruct)
	UpdateProfile(ctx context.Context, id string, data *dto.UserUpdateReq) (customErr *helper.ErrorStruct)
	DeleteAlamatByID(ctx context.Context, id string) (customErr *helper.ErrorStruct)
}

type UserUseCaseImpl struct {
	userRepository repository.UserRepository
}

// NewUserUseCase returns the usecase for the user group data
func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{
		userRepository: userRepository,
	}
}

// GetMyAlamats handles the business logic to retrieve alamat data of the current user
func (alc *UserUseCaseImpl) GetMyAlamats(ctx context.Context, token string, filter *dto.AlamatFilter) (res []*dto.AlamatResp, customErr *helper.ErrorStruct) {
	claims, err := utils.GetJWTClaims(token)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	resRepo, err := alc.userRepository.GetAlamatsByUserId(ctx, claims.UserId, &daos.FilterAlamat{JudulAlamat: filter.JudulAlamat})
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	for _, v := range resRepo {
		res = append(res, utils.AlamatToAlamatResp(v))
	}
	return res, nil
}

// GetAlamatById handles the business logic to retrieve alamat data having the id
func (alc *UserUseCaseImpl) GetAlamatById(ctx context.Context, id string) (res *dto.AlamatResp, customErr *helper.ErrorStruct) {
	resRepo, err := alc.userRepository.GetAlamatById(ctx, id)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	res = utils.AlamatToAlamatResp(resRepo)
	return res, nil
}

// GetMyProfile handles the business logic to retrieve user data of the current user
func (alc *UserUseCaseImpl) GetMyProfile(ctx context.Context, token string) (res *dto.UserResp, customErr *helper.ErrorStruct) {
	claims, err := utils.GetJWTClaims(token)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	resRepo, err := alc.userRepository.GetUserById(ctx, claims.UserId)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	provinceData, err := alc.userRepository.GetProvinceById(ctx, resRepo.IdProvinsi)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	cityData, err := alc.userRepository.GetCityById(ctx, resRepo.IdKota)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	res = utils.UserToUserResp(resRepo)
	res.IdProvinsi = provinceData
	res.IdKota = cityData

	return res, nil
}

// CreateAlamat handles the business logic to insert the user data
func (alc *UserUseCaseImpl) CreateAlamat(ctx context.Context, data *dto.AlamatCreateReq, token string) (res uint, customErr *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errValidate.Error()))
		return 0, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	claims, err := utils.GetJWTClaims(token)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return 0, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	userId, err := strconv.Atoi(claims.UserId)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return 0, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	alamatId, err := alc.userRepository.CreateAlamat(ctx, &daos.Alamat{
		IdUser:       uint(userId),
		JudulAlamat:  data.JudulAlamat,
		NamaPenerima: data.NamaPenerima,
		Notelp:       data.Notelp,
		DetailAlamat: data.DetailAlamat,
	})
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return 0, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	return alamatId, nil
}

// UpdateAlamatById handles the business logic to update alamat data having the
func (alc *UserUseCaseImpl) UpdateAlamatById(ctx context.Context, id string, data *dto.AlamatUpdateReq) (customErr *helper.ErrorStruct) {
	err := alc.userRepository.UpdateAlamatByID(ctx, id, &daos.Alamat{
		JudulAlamat:  data.JudulAlamat,
		NamaPenerima: data.NamaPenerima,
		Notelp:       data.Notelp,
		DetailAlamat: data.DetailAlamat,
	})
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	return nil
}

// UpdateAlamatById handles the business logic to update user data of the current user
func (alc *UserUseCaseImpl) UpdateProfile(ctx context.Context, token string, data *dto.UserUpdateReq) (customErr *helper.ErrorStruct) {
	claims, err := utils.GetJWTClaims(token)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	katasandi, err := utils.HashPassword(data.KataSandi)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	tanggalLahir, err := utils.StringToDate(data.TanggalLahir)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	err = alc.userRepository.UpdateUserById(ctx, claims.UserId, &daos.User{
		Nama:         data.Nama,
		KataSandi:    katasandi,
		Notelp:       data.Notelp,
		TanggalLahir: tanggalLahir,
		JenisKelamin: data.JenisKelamin,
		Tentang:      data.Tentang,
		Pekerjaan:    data.Pekerjaan,
		Email:        data.Email,
		IdProvinsi:   data.IdProvinsi,
		IdKota:       data.IdKota,
	})
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	return nil
}

// DeleteAlamatByID handles the business logic to delete alamat data having the id
func (alc *UserUseCaseImpl) DeleteAlamatByID(ctx context.Context, id string) (customErr *helper.ErrorStruct) {
	err := alc.userRepository.DeleteAlamatById(ctx, id)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	return nil
}
