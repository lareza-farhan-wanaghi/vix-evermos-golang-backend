package usecase

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type CategoryUseCase interface {
	GetAllCategories(ctx context.Context) (res []*dto.CategoryResp, customErr *helper.ErrorStruct)
	GetCategoryById(ctx context.Context, id string) (res *dto.CategoryResp, customErr *helper.ErrorStruct)
	CreateCategory(ctx context.Context, data *dto.CategoryCreateReq) (res uint, customErr *helper.ErrorStruct)
	UpdateCategoryById(ctx context.Context, id string, data *dto.CategoryUpdateReq) (customErr *helper.ErrorStruct)
	DeleteCategoryByID(ctx context.Context, id string) (customErr *helper.ErrorStruct)
}

type CategoryUseCaseImpl struct {
	categoryRepository repository.CategoryRepository
}

// NewCategoryUseCase returns the usecase for the category group path
func NewCategoryUseCase(categoryRepository repository.CategoryRepository) CategoryUseCase {
	return &CategoryUseCaseImpl{
		categoryRepository: categoryRepository,
	}
}

// GetAllCategories handles the business logic to retrieve all stored category data
func (alc *CategoryUseCaseImpl) GetAllCategories(ctx context.Context) (res []*dto.CategoryResp, customErr *helper.ErrorStruct) {
	resRepo, err := alc.categoryRepository.GetAllCategory(ctx)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	for _, v := range resRepo {
		res = append(res, utils.CatergoryToCategoryResp(v))
	}
	return res, nil
}

// GetCategoryById handles the business logic to retrieve stored category data having the id
func (alc *CategoryUseCaseImpl) GetCategoryById(ctx context.Context, id string) (res *dto.CategoryResp, customErr *helper.ErrorStruct) {
	resRepo, err := alc.categoryRepository.GetCategoryById(ctx, id)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	res = utils.CatergoryToCategoryResp(resRepo)

	return res, nil
}

// CreateCategory handles the business logic to insert the category data
func (alc *CategoryUseCaseImpl) CreateCategory(ctx context.Context, data *dto.CategoryCreateReq) (res uint, customErr *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errValidate.Error()))
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, err := alc.categoryRepository.CreateCategory(ctx, &daos.Category{
		NamaCategory: data.NamaCategory,
	})
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	return resRepo, nil
}

// UpdateCategoryById handles the business logic to update category data having the id
func (alc *CategoryUseCaseImpl) UpdateCategoryById(ctx context.Context, id string, data *dto.CategoryUpdateReq) (customErr *helper.ErrorStruct) {
	err := alc.categoryRepository.UpdateCategoryById(ctx, id, &daos.Category{
		NamaCategory: data.NamaCategory,
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

// DeleteCategoryByID handles the business logic to delete category data having the id
func (alc *CategoryUseCaseImpl) DeleteCategoryByID(ctx context.Context, id string) (customErr *helper.ErrorStruct) {
	err := alc.categoryRepository.DeleteCategoryById(ctx, id)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	return nil
}
