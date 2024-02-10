package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshtrigger/api/models"
	"gorm.io/gorm"
)

type BookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type bookhandler struct {
	service models.BookService
}

type BookHandler interface {
	GetBooks(c *gin.Context)
	GetBook(c *gin.Context)
	UpdateBook(c *gin.Context)
	CreateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
}

func InitBookHandler(db *gorm.DB) BookHandler {
	s := models.InitBookService(db)

	return &bookhandler{service: s}
}

func (handler *bookhandler) GetBooks(c *gin.Context) {
	books := handler.service.GetBooks()

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func (handler *bookhandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	book, err := handler.service.GetBook(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (handler *bookhandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	_, err := handler.service.DeleteBook(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func (handler *bookhandler) UpdateBook(c *gin.Context) {
	var input BookInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	book, err := handler.service.UpdateBook(id, models.Book{Title: input.Title, Description: input.Description})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "record not found here lol"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (handler *bookhandler) CreateBook(c *gin.Context) {
	var input BookInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := models.Book{Title: input.Title, Description: input.Description}
	book := handler.service.CreateBook(data)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
