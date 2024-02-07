package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joshtrigger/api/controllers"
	"github.com/joshtrigger/api/models"
)

func main() {
	r := gin.Default()
	
	models.ConnectDatabase()

	bookhandler := controllers.InitBookHandler()

	r.GET("/books", bookhandler.GetBooks)
	r.GET("/books/:id", bookhandler.GetBook)
	r.DELETE("/books/:id", bookhandler.DeleteBook)
	r.PATCH("/books/:id", bookhandler.UpdateBook)
	r.POST("/books", bookhandler.CreateBook)

	r.Run()
}
