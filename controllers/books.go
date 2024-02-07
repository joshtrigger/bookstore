package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshtrigger/api/models"
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

func InitBookHandler() BookHandler {
	s := models.InitBookService(models.DB)

	return &bookhandler{service: s}
}

func (handler *bookhandler) GetBooks(c *gin.Context) {
	books := handler.service.GetBooks(c)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func (handler *bookhandler) GetBook(c *gin.Context) {
	book, err := handler.service.GetBook(c)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (handler *bookhandler) DeleteBook(c *gin.Context) {
	_, err := handler.service.DeleteBook(c)

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

	book, err := handler.service.UpdateBook(c, models.Book{Title: input.Title, Description: input.Description})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "record not found"})
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
	book := handler.service.CreateBook(c, data)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
