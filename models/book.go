package models

import (
	"github.com/gin-gonic/gin"
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
	GetBooks(c *gin.Context) (*[]Book)
	GetBook(c *gin.Context) (*Book, error)
	UpdateBook(c *gin.Context, input Book) (*Book, error)
	CreateBook(c *gin.Context, input Book) (*Book)
	DeleteBook(c *gin.Context) (bool, error)
}

func InitBookService(db *gorm.DB ) BookService {
	return &bookService{db}
}

func (service *bookService) GetBooks(c *gin.Context) (*[]Book) {
	var books []Book

	service.db.Find(&books)

	return &books
}

func (service *bookService) GetBook(c *gin.Context) (*Book, error) {
	var book Book

	if err := DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return &Book{}, err
	}
	return &Book{}, nil
}

func (service *bookService) DeleteBook(c *gin.Context) (bool, error) {
	var book Book

	if err := DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return false, err
	}
	DB.Delete(&book)

	return true, nil
}

func (service *bookService) UpdateBook(c *gin.Context, input Book) (*Book, error) {
	var book Book

	if err := DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return &Book{}, err
	}

	DB.Model(&book).Updates(input)

	return &book, nil
}

func (service *bookService) CreateBook(c *gin.Context, input Book) (*Book) {
	book := Book{Title: input.Title, Description: input.Description}
  
	DB.Create(&book)

	return &book
}
