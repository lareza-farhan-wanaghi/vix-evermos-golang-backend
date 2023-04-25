package usecase

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TokoUseCase interface {
	GetAllTokos(ctx context.Context, queries *dto.TokoFilter) (res *dto.AllTokoResp, err *helper.ErrorStruct)
	GetTokoById(ctx context.Context, param, header string) (res *dto.TokoResp, err *helper.ErrorStruct)
	GetMyToko(ctx context.Context, header string) (res *dto.TokoResp, err *helper.ErrorStruct)
	UpdateTokoByID(ctx context.Context, id string, photo *multipart.FileHeader, data *dto.TokoUpdateReq) (customErr *helper.ErrorStruct)
}

type TokoUseCaseImpl struct {
	tokoRepository repository.TokoRepository
	jwtSecret      string
}

// NewTokoUseCase returns the usecase for the toko group path
func NewTokoUseCase(tokoRepository repository.TokoRepository, jwtSecret string) TokoUseCase {
	return &TokoUseCaseImpl{
		tokoRepository: tokoRepository,
		jwtSecret:      jwtSecret,
	}
}

// GetAllTokos handles the business logic to retrieve all toko data
func (alc *TokoUseCaseImpl) GetAllTokos(ctx context.Context, queries *dto.TokoFilter) (res *dto.AllTokoResp, err *helper.ErrorStruct) {
	if queries.Limit < 1 {
		queries.Limit = 10
	}

	if queries.Page < 1 {
		queries.Page = 1
	}

	resRepo, errRepo := alc.tokoRepository.GetAllTokos(ctx, daos.FilterToko{
		Limit:    queries.Limit,
		Offset:   (queries.Page - 1) * queries.Limit,
		NamaToko: queries.NamaToko,
	})
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("no data toko"),
		}
	}

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = utils.TokoArrayToAllTokoResp(resRepo)
	res.Page = queries.Page
	res.Limit = queries.Limit
	return res, nil
}

// GetTokoById handles the business logic to retrieve toko data having the id
func (alc *TokoUseCaseImpl) GetTokoById(ctx context.Context, param, header string) (res *dto.TokoResp, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.tokoRepository.GetTokoById(ctx, param)
	if errRepo != nil {
		if errRepo == gorm.ErrRecordNotFound {
			errRepo = errors.New("toko tidak ditemukan")
		}

		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = utils.TokoToTokoResp(resRepo)
	return res, nil
}

// GetMyToko handles the business logic to retrieve toko data of the user specified on the token
func (alc *TokoUseCaseImpl) GetMyToko(ctx context.Context, token string) (res *dto.TokoResp, err *helper.ErrorStruct) {
	userId, errGetClaims := utils.GetJWTUserIdString(token)
	if errGetClaims != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errGetClaims.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errGetClaims,
		}
	}

	resRepo, errRepo := alc.tokoRepository.GetTokoByUserID(ctx, userId)
	if errRepo != nil {
		if errRepo == gorm.ErrRecordNotFound {
			errRepo = errors.New("toko tidak ditemukan")
		}

		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = utils.TokoToTokoResp(resRepo)
	return res, nil
}

// UpdateTokoByID handles the business logic to update toko data having the id
func (alc *TokoUseCaseImpl) UpdateTokoByID(ctx context.Context, id string, photo *multipart.FileHeader, data *dto.TokoUpdateReq) (customErr *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errValidate.Error()))
		return &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	res, err := alc.tokoRepository.GetTokoById(ctx, id)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	updatedToko := &daos.Toko{
		NamaToko: data.NamaToko,
	}

	if photo != nil {
		os.Remove(fmt.Sprintf(".%s", res.UrlFoto))

		internalFilepath := fmt.Sprintf("%s%d%s", utils.TokoImagesPath, time.Now().UnixNano(), filepath.Ext(photo.Filename))
		err := utils.SaveMultiFormImage(photo, internalFilepath, 1000000, map[string]struct{}{
			"image/jpg":  {},
			"image/png":  {},
			"image/jpeg": {},
		})
		if err != nil {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, err.Error())
			return &helper.ErrorStruct{
				Code: fiber.StatusBadRequest,
				Err:  err,
			}
		}

		updatedToko.UrlFoto = internalFilepath[1:]
	}

	if errRepo := alc.tokoRepository.UpdateToko(ctx, res, updatedToko); errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return nil
}
