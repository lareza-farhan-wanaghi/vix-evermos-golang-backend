package repository

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategory(ctx context.Context) (res []*daos.Category, err error)
	GetCategoryById(ctx context.Context, id string) (res *daos.Category, err error)
	GetUserById(ctx context.Context, id string) (res *daos.User, err error)
	CreateCategory(ctx context.Context, data *daos.Category) (res uint, err error)
	UpdateCategoryById(ctx context.Context, id string, data *daos.Category) (err error)
	DeleteCategoryById(ctx context.Context, id string) (err error)
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

// NewCategoryRepository returns the repository for the category group path
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

// GetAllCategory returns all cateory data from the category table
func (alr *CategoryRepositoryImpl) GetAllCategory(ctx context.Context) (res []*daos.Category, err error) {
	if err := alr.db.WithContext(ctx).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetCategoryById returns cateory data having the id from the category table
func (alr *CategoryRepositoryImpl) GetCategoryById(ctx context.Context, id string) (res *daos.Category, err error) {
	res = &daos.Category{}
	if err := alr.db.WithContext(ctx).Where("id = ?", id).First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetUserById returns user data having the id from the user table
func (alr *CategoryRepositoryImpl) GetUserById(ctx context.Context, id string) (res *daos.User, err error) {
	if err := alr.db.WithContext(ctx).Model(&daos.User{}).Where("id = ?", id).First(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// CreateCategory inserts the category data to the category table
func (alr *CategoryRepositoryImpl) CreateCategory(ctx context.Context, data *daos.Category) (res uint, err error) {
	result := alr.db.WithContext(ctx).Create(data)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

// UpdateCategoryById updates category data having the id on the category table
func (alr *CategoryRepositoryImpl) UpdateCategoryById(ctx context.Context, id string, data *daos.Category) (err error) {
	if err = alr.db.WithContext(ctx).Where("id = ? ", id).First(&daos.Category{}).Error; err != nil {
		return gorm.ErrRecordNotFound
	}
	fmt.Printf("id: %v data:%v\n", data.ID, data)
	if err := alr.db.WithContext(ctx).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// DeleteCategoryById deletes category data having the id on the category table
func (alr *CategoryRepositoryImpl) DeleteCategoryById(ctx context.Context, id string) (err error) {
	if err = alr.db.WithContext(ctx).Where("id = ? ", id).First(&daos.Category{}).Error; err != nil {
		return gorm.ErrRecordNotFound
	}

	if err := alr.db.WithContext(ctx).Where("id = ?", id).Delete(&daos.Category{}).Error; err != nil {
		return err
	}

	return nil
}
