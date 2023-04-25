package usecase

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProdukUseCase interface {
	GetAllProduks(ctx context.Context, filter *dto.ProdukFilter) (res *dto.AllProdukResp, customErr *helper.ErrorStruct)
	GetProdukById(ctx context.Context, param string) (res *dto.ProdukResp, customErr *helper.ErrorStruct)
	CreateProduk(ctx context.Context, data *dto.ProdukCreateReq, token string, photos []*multipart.FileHeader) (res uint, customErr *helper.ErrorStruct)
	UpdateProdukByID(ctx context.Context, data *dto.ProdukUpdateReq, id string, photos []*multipart.FileHeader) (customErr *helper.ErrorStruct)
	DeleteProdukByID(ctx context.Context, id string) (customErr *helper.ErrorStruct)
}

type ProdukUseCaseImpl struct {
	produkRepository repository.ProdukRepository
}

// NewProdukUseCase returns the usecase for the produk group path
func NewProdukUseCase(produkRepository repository.ProdukRepository) ProdukUseCase {
	return &ProdukUseCaseImpl{
		produkRepository: produkRepository,
	}
}

// GetAllProduks handles the business logic to retrieve all produk data
func (alc *ProdukUseCaseImpl) GetAllProduks(ctx context.Context, filter *dto.ProdukFilter) (res *dto.AllProdukResp, customErr *helper.ErrorStruct) {
	if filter.Limit < 1 {
		filter.Limit = 10
	}

	if filter.Page < 1 {
		filter.Page = 1
	}

	resRepo, err := alc.produkRepository.GetAllProduks(ctx, &daos.FilterProduk{
		Limit:      filter.Limit,
		Offset:     (filter.Page - 1) * filter.Limit,
		CategoryId: filter.CategoryId,
		TokoId:     filter.TokoId,
		NamaProduk: filter.NamaProduk,
		MinHarga:   filter.MinHarga,
		MaxHarga:   filter.MaxHarga,
	})
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	res, err = utils.ProdukArrayToAllProdukResp(resRepo)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}
	res.Page = filter.Page
	res.Limit = filter.Limit

	return res, nil
}

// GetProdukById handles the business logic to retrieve produk data having the id
func (alc *ProdukUseCaseImpl) GetProdukById(ctx context.Context, param string) (res *dto.ProdukResp, customErr *helper.ErrorStruct) {
	resRepo, err := alc.produkRepository.GetProdukById(ctx, param)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.New("no data produk")
		}

		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	produkResp, err := utils.ProdukToProdukResp(resRepo)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusInternalServerError,
			Err:  err,
		}
	}
	return produkResp, nil
}

// CreateProduk handles the business logic to insert the produk data
func (alc *ProdukUseCaseImpl) CreateProduk(ctx context.Context, data *dto.ProdukCreateReq, token string, photos []*multipart.FileHeader) (res uint, customErr *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errValidate.Error()))
		return 0, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	userId, err := utils.GetJWTUserIdString(token)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return 0, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	resRepo, err := alc.produkRepository.GetUserById(ctx, userId)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return 0, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	stok, err := strconv.Atoi(data.Stok)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return 0, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	idCategory, err := strconv.Atoi(data.CategoryId)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return 0, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	idProduk, err := alc.produkRepository.CreateProduk(ctx, &daos.Produk{
		NamaProduk:    data.NamaProduk,
		Slug:          strings.Replace(strings.ToLower(data.NamaProduk), " ", "-", -1),
		HargaKonsumen: data.HargaKonsumen,
		HargaReseller: data.HargaReseller,
		Stok:          stok,
		Deskripsi:     data.Deskripsi,
		IdCategory:    uint(idCategory),
		IdToko:        resRepo.Toko.ID,
	})
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return 0, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	if len(photos) > 0 {
		for _, fileHeader := range photos {
			internalFilepath := fmt.Sprintf("%s%d%s", utils.ProdukImagesPath, time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
			err = utils.SaveMultiFormImage(fileHeader, internalFilepath, 1000000, map[string]struct{}{
				"image/jpg":  {},
				"image/png":  {},
				"image/jpeg": {},
			})
			if err != nil {
				helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, err.Error())
				return 0, &helper.ErrorStruct{
					Code: fiber.StatusBadRequest,
					Err:  err,
				}
			}

			alc.produkRepository.CreateFotoProduk(ctx, &daos.FotoProduk{
				IdProduk: idProduk,
				Url:      internalFilepath[1:],
			})
		}
	}

	return idProduk, nil
}

// UpdateProdukByID handles the business logic to update produk data having the id
func (alc *ProdukUseCaseImpl) UpdateProdukByID(ctx context.Context, data *dto.ProdukUpdateReq, id string, photos []*multipart.FileHeader) (customErr *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errValidate.Error()))
		return &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, err := alc.produkRepository.GetProdukById(ctx, id)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	produkData := &daos.Produk{
		NamaProduk:    data.NamaProduk,
		Slug:          strings.ReplaceAll(strings.ToLower(data.NamaProduk), " ", "-"),
		HargaReseller: data.HargaReseller,
		HargaKonsumen: data.HargaKonsumen,
		Deskripsi:     data.Deskripsi,
	}

	if data.Stok != "" {
		stok, err := strconv.Atoi(data.Stok)
		if err != nil {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
			return &helper.ErrorStruct{
				Code: fiber.StatusBadRequest,
				Err:  err,
			}
		}
		produkData.Stok = stok
	}

	if data.CategoryId != "" {
		categoryId, err := strconv.Atoi(data.CategoryId)
		if err != nil {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
			return &helper.ErrorStruct{
				Code: fiber.StatusBadRequest,
				Err:  err,
			}
		}
		produkData.IdCategory = uint(categoryId)
	}

	err = alc.produkRepository.UpdateProduk(ctx, resRepo, produkData)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	if len(photos) > 0 {
		for _, v := range resRepo.FotoProduks {
			alc.produkRepository.DeleteFotoProduk(ctx, v)
			os.Remove(fmt.Sprintf(".%s", v.Url))
		}

		for _, fileHeader := range photos {
			internalFilepath := fmt.Sprintf("%s%d%s", utils.ProdukImagesPath, time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
			err = utils.SaveMultiFormImage(fileHeader, internalFilepath, 1000000, map[string]struct{}{
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

			alc.produkRepository.CreateFotoProduk(ctx, &daos.FotoProduk{
				IdProduk: resRepo.ID,
				Url:      internalFilepath[1:],
			})
		}
	}

	return nil
}

// DeleteProdukByID handles the business logic to delete produk data having the id
func (alc *ProdukUseCaseImpl) DeleteProdukByID(ctx context.Context, id string) (customErr *helper.ErrorStruct) {
	resRepo, err := alc.produkRepository.GetProdukById(ctx, id)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}
	err = alc.produkRepository.DeleteProduk(ctx, resRepo)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	for _, v := range resRepo.FotoProduks {
		alc.produkRepository.DeleteFotoProduk(ctx, v)
		os.Remove(fmt.Sprintf(".%s", v.Url))
	}

	return nil
}
