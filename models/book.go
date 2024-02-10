package models

import (
	"gorm.io/gorm"
)

type BookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

type bookService struct {
	db *gorm.DB
}

type BookService interface {
	GetBooks() *[]Book
	GetBook(id string) (*Book, error)
	UpdateBook(id string, input Book) (*Book, error)
	CreateBook(input Book) *Book
	DeleteBook(id string) (bool, error)
}

func InitBookService(db *gorm.DB) BookService {
	return &bookService{db}
}

func (service *bookService) GetBooks() *[]Book {
	var books []Book

	service.db.Find(&books)

	return &books
}

func (service *bookService) GetBook(id string) (*Book, error) {
	var book Book

	if err := service.db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (service *bookService) DeleteBook(id string) (bool, error) {
	var book Book

	if err := service.db.Where("id = ?", id).First(&book).Error; err != nil {
		return false, err
	}
	service.db.Delete(&book)

	return true, nil
}

func (service *bookService) UpdateBook(id string, input Book) (*Book, error) {
	var book Book

	if err := service.db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}

	service.db.Model(&book).Updates(input)

	return &book, nil
}

func (service *bookService) CreateBook(input Book) *Book {
	book := Book{Title: input.Title, Description: input.Description}

	service.db.Create(&book)

	return &book
}

// func (service *bookService) GetBookById(id string) (*Book, error) {
// 	var book Book

// 	if err := service.db.Where("id = ?", id).First(&book).Error; err != nil {
// 		return &Book{}, err
// 	}
// 	return &book, nil
// }
