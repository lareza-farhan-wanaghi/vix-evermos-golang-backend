package controller

import (
	"tugas_akhir_example/internal/helper"
	bookdto "tugas_akhir_example/internal/pkg/dto"
	bookusecase "tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type BookController interface {
	GetAllBook(ctx *fiber.Ctx) error
	GetBookByID(ctx *fiber.Ctx) error
	CreateBook(ctx *fiber.Ctx) error
	UpdateBookByID(ctx *fiber.Ctx) error
	DeleteBookByID(ctx *fiber.Ctx) error
}

type BookControllerImpl struct {
	bookusecase bookusecase.BookUseCase
}

// NewBookController returns the controller for the book group path
func NewBookController(bookusecase bookusecase.BookUseCase) BookController {
	return &BookControllerImpl{
		bookusecase: bookusecase,
	}
}

// GetAllBook handles the delivery logic to retrieve all book data
func (uc *BookControllerImpl) GetAllBook(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := new(bookdto.BookFilter)
	if err := ctx.QueryParser(filter); err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	res, customErr := uc.bookusecase.GetAllBooks(c, bookdto.BookFilter{
		Title: filter.Title,
		Limit: filter.Limit,
		Page:  filter.Page,
	})

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
		Data:       res,
	})
}

// GetBookByID handles the delivery logic to retrieve book data having the id
func (uc *BookControllerImpl) GetBookByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	bookid := ctx.Params("id_book")
	if bookid == "" {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{"Bad Request"},
		})
	}

	res, customErr := uc.bookusecase.GetBookByID(c, bookid)
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
		Data:       res,
	})
}

// CreateBook handles the delivery logic to insert the book data
func (uc *BookControllerImpl) CreateBook(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := new(bookdto.BookReqCreate)
	if err := ctx.BodyParser(data); err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	res, customErr := uc.bookusecase.CreateBook(c, *data)
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
		Data:       res,
	})
}

// UpdateBookByID handles the delivery logic to update book data having the id
func (uc *BookControllerImpl) UpdateBookByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	bookid := ctx.Params("id_book")
	if bookid == "" {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{"Bad request"},
		})
	}

	data := new(bookdto.BookReqUpdate)
	if err := ctx.BodyParser(data); err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{"Bad request"},
		})
	}

	res, customErr := uc.bookusecase.UpdateBookByID(c, bookid, *data)
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
		Data:       res,
	})
}

// DeleteBookByID handles the delivery logic to delete book data having the id
func (uc *BookControllerImpl) DeleteBookByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	bookid := ctx.Params("id_book")
	if bookid == "" {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{"Bad request"},
		})
	}

	res, customErr := uc.bookusecase.DeleteBookByID(c, bookid)
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
		Data:       res,
	})
}
