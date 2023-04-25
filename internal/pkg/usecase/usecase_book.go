package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	bookdto "tugas_akhir_example/internal/pkg/dto"
	bookrepository "tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BookUseCase interface {
	GetAllBooks(ctx context.Context, params bookdto.BookFilter) (res []bookdto.BookResp, err *helper.ErrorStruct)
	GetBookByID(ctx context.Context, bookid string) (res bookdto.BookResp, err *helper.ErrorStruct)
	CreateBook(ctx context.Context, data bookdto.BookReqCreate) (res uint, err *helper.ErrorStruct)
	UpdateBookByID(ctx context.Context, bookid string, data bookdto.BookReqUpdate) (res string, err *helper.ErrorStruct)
	DeleteBookByID(ctx context.Context, bookid string) (res string, err *helper.ErrorStruct)
}

type BookUseCaseImpl struct {
	bookrepository bookrepository.BookRepository
}

// NewBookUseCase returns the usecase for the book group path
func NewBookUseCase(bookrepository bookrepository.BookRepository) BookUseCase {
	return &BookUseCaseImpl{
		bookrepository: bookrepository,
	}

}

// GetAllBooks handles the business logic to retrieve all stored book data
func (alc *BookUseCaseImpl) GetAllBooks(ctx context.Context, params bookdto.BookFilter) (res []bookdto.BookResp, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	resRepo, errRepo := alc.bookrepository.GetAllBooks(ctx, daos.FilterBook{
		Limit:  params.Limit,
		Offset: params.Page,
		Title:  params.Title,
	})
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("no data book"),
		}
	}

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	for _, v := range resRepo {
		res = append(res, bookdto.BookResp{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Author:      v.Author,
		})
	}

	return res, nil
}

// GetBookByID handles the business logic to retrieve book data having the id
func (alc *BookUseCaseImpl) GetBookByID(ctx context.Context, bookid string) (res bookdto.BookResp, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.bookrepository.GetBookByID(ctx, bookid)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("no data book"),
		}
	}

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = bookdto.BookResp{
		ID:          resRepo.ID,
		Title:       resRepo.Title,
		Description: resRepo.Description,
		Author:      resRepo.Author,
	}

	return res, nil
}

// CreateBook handles the business logic to insert the book data
func (alc *BookUseCaseImpl) CreateBook(ctx context.Context, data bookdto.BookReqCreate) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.bookrepository.CreateBook(ctx, daos.Book{
		Title:       data.Title,
		Description: data.Description,
		Author:      data.Author,
	})
	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

// UpdateBookByID handles the business logic to update book data having the id
func (alc *BookUseCaseImpl) UpdateBookByID(ctx context.Context, bookid string, data bookdto.BookReqUpdate) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.bookrepository.UpdateBookByID(ctx, bookid, daos.Book{
		Title:       data.Title,
		Description: data.Description,
		Author:      data.Author,
	})

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

// DeleteBookByID handles the business logic to delete book data having the id
func (alc *BookUseCaseImpl) DeleteBookByID(ctx context.Context, bookid string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.bookrepository.DeleteBookByID(ctx, bookid)
	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
