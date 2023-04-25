package repository

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBooks(ctx context.Context, params daos.FilterBook) (res []daos.Book, err error)
	GetBookByID(ctx context.Context, bookid string) (res daos.Book, err error)
	CreateBook(ctx context.Context, data daos.Book) (res uint, err error)
	UpdateBookByID(ctx context.Context, bookid string, data daos.Book) (res string, err error)
	DeleteBookByID(ctx context.Context, bookid string) (res string, err error)
}

type BookRepositoryImpl struct {
	db *gorm.DB
}

// NewBookRepository returns the repository for the book group path
func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{
		db: db,
	}
}

// GetAllBooks returns all book data from the book table
func (alr *BookRepositoryImpl) GetAllBooks(ctx context.Context, params daos.FilterBook) (res []daos.Book, err error) {
	db := alr.db

	filter := map[string][]any{
		"title like ? or description like ? or author like ?": {fmt.Sprint("%" + params.Title), "%ab ", "%ab"},
	}

	for key, val := range filter {
		db = db.Where(key, val...)
	}

	if err := db.Debug().WithContext(ctx).Limit(params.Limit).Offset(params.Offset).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// GetBookByID returns book data having the id from the book table
func (alr *BookRepositoryImpl) GetBookByID(ctx context.Context, bookid string) (res daos.Book, err error) {
	if err := alr.db.First(&res, bookid).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

// CreateBook inserts the book data to the book table
func (alr *BookRepositoryImpl) CreateBook(ctx context.Context, data daos.Book) (res uint, err error) {
	result := alr.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

// UpdateBookByID updates book data having the id on the book table
func (alr *BookRepositoryImpl) UpdateBookByID(ctx context.Context, bookid string, data daos.Book) (res string, err error) {
	var dataBook daos.Book
	if err = alr.db.Where("id = ? ", bookid).First(&dataBook).WithContext(ctx).Error; err != nil {
		return "Update book failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataBook).Updates(&data).Where("id = ? ", bookid).Error; err != nil {
		return "Update book failed", err
	}

	return res, nil
}

// DeleteBookByID deletes book data having the id on the book table
func (alr *BookRepositoryImpl) DeleteBookByID(ctx context.Context, bookid string) (res string, err error) {
	var dataBook daos.Book
	if err = alr.db.Where("id = ?", bookid).First(&dataBook).WithContext(ctx).Error; err != nil {
		return "Delete book failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataBook).Delete(&dataBook).Error; err != nil {
		return "Delete book failed", err
	}

	return res, nil
}
