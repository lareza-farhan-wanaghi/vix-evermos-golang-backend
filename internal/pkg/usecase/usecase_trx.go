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
)

type TrxUseCase interface {
	GetAllTrxs(ctx context.Context, filter *dto.TrxFilter) (res *dto.AllTrxResp, customErr *helper.ErrorStruct)
	GetTrxById(ctx context.Context, id string) (res *dto.TrxResp, customErr *helper.ErrorStruct)
	CreateTrx(ctx context.Context, token string, data *dto.TrxCreateReq) (res uint, customErr *helper.ErrorStruct)
}

type TrxUseCaseImpl struct {
	trxRepository repository.TrxRepository
}

// NewTrxUseCase returns the usecase for the trx group path
func NewTrxUseCase(trxRepository repository.TrxRepository) TrxUseCase {
	return &TrxUseCaseImpl{
		trxRepository: trxRepository,
	}
}

// GetAllTrxs handles the business logic to retrieve all trx data
func (alc *TrxUseCaseImpl) GetAllTrxs(ctx context.Context, filter *dto.TrxFilter) (res *dto.AllTrxResp, customErr *helper.ErrorStruct) {
	if filter.Limit < 1 {
		filter.Limit = 10
	}

	if filter.Page < 1 {
		filter.Page = 1
	}

	resRepo, err := alc.trxRepository.GetAllTrxs(ctx, &daos.FilterTrx{
		Limit:       filter.Limit,
		Offset:      (filter.Page - 1) * filter.Limit,
		KodeInvoice: filter.Search,
	})
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	res, err = utils.TrxArrayToAllTrxResp(resRepo)
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

// GetTrxById handles the business logic to retrieve trx data having the id
func (alc *TrxUseCaseImpl) GetTrxById(ctx context.Context, id string) (res *dto.TrxResp, customErr *helper.ErrorStruct) {
	resRepo, err := alc.trxRepository.GetTrxById(ctx, id)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	res, err = utils.TrxToTrxResp(resRepo)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	return res, nil
}

// CreateTrx handles the business logic to insert the trx data
func (alc *TrxUseCaseImpl) CreateTrx(ctx context.Context, token string, data *dto.TrxCreateReq) (res uint, customErr *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errValidate.Error()))
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	userId, err := utils.GetJWTUserId(token)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return 0, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	detailTrxes := []*daos.DetailTrx{}
	trxHargaTotal := 0
	for _, v := range data.DetailTrxes {
		resRepoProduk, err := alc.trxRepository.GetProdukById(ctx, strconv.Itoa(int(v.ProductId)))
		if err != nil {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
			return 0, &helper.ErrorStruct{
				Code: fiber.StatusBadRequest,
				Err:  err,
			}
		}

		logProduk := &daos.LogProduk{
			IdProduk:      resRepoProduk.ID,
			NamaProduk:    resRepoProduk.NamaProduk,
			Slug:          resRepoProduk.Slug,
			HargaReseller: resRepoProduk.HargaReseller,
			HargaKonsumen: resRepoProduk.HargaKonsumen,
			Deskripsi:     resRepoProduk.Deskripsi,
			IdToko:        resRepoProduk.IdToko,
			IdCategory:    resRepoProduk.IdCategory,
		}

		hargaKonsumen, err := strconv.Atoi(logProduk.HargaKonsumen)
		if err != nil {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
			return 0, &helper.ErrorStruct{
				Code: fiber.StatusBadRequest,
				Err:  err,
			}
		}

		detailHargaTotal := v.Kuantitas * hargaKonsumen
		trxHargaTotal += detailHargaTotal

		detailTrxes = append(detailTrxes, &daos.DetailTrx{
			LogProduk:  logProduk,
			Kuantitas:  int(v.Kuantitas),
			HargaTotal: detailHargaTotal,
			IdToko:     logProduk.IdToko,
		})
	}

	resRepoAlamat, err := alc.trxRepository.GetAlamatById(ctx, strconv.Itoa(int(data.AlamatKirim)))
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return 0, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	if resRepoAlamat.IdUser != userId {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return 0, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errors.New("unauthorized alamat kirim"),
		}
	}

	trx := &daos.Trx{
		IdUser:           userId,
		AlamatPengiriman: resRepoAlamat.ID,
		HargaTotal:       trxHargaTotal,
		KodeInvoice:      fmt.Sprintf("INV-%d", time.Now().UnixNano()),
		MethodBayar:      data.MethodBayar,
		DetailTrxs:       detailTrxes,
	}

	resRepo, err := alc.trxRepository.CreateTrx(ctx, trx)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	return resRepo, nil
}
