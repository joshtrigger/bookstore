package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"

	// "fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	// "github.com/joshtrigger/api/mocks"
	"github.com/joshtrigger/api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// "github.com/joshtrigger/api/models"
	"github.com/stretchr/testify/suite"
)

type bookhandlerSuite struct {
	suite.Suite
	// handlerMock *mocks.BookHandler
	// router      *gin.Engine
	// bookhandler *bookhandler
	// testingServer *httptest.Server
	// handler bookhandler
}

var router = gin.Default()
var db = connectTestDatabase()
var handler = InitBookHandler(db)

// var router *gin.Engine

func TestBookHandler(t *testing.T) {
	suite.Run(t, new(bookhandlerSuite))
}

func (suite *bookhandlerSuite) SetUpSuite() {
	// mock := new(mocks.BookHandler)

	// suite.router = router
	// suite.handlerMock = mock
}

func (suite *bookhandlerSuite) TestCreateBook() {
	book := models.Book{Title: "new book", Description: "brand new"}
	router.POST("/books", handler.CreateBook)
	// router.GET("/books/:id", handler.GetBook)

	json_data, _ := json.Marshal(book)
	post_req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(json_data))
	// req, _ := http.NewRequest("GET", "/books/2", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, post_req)
	// router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	suite.Contains(string(responseData), "new book")
	suite.Equal(http.StatusOK, w.Code)
}

func (suite *bookhandlerSuite) TestGetBook() {
	book := models.Book{Title: "sample", Description: "sample"}
	// router.POST("/books", handler.CreateBook)
	router.GET("/books/:id", handler.GetBook)

	json_data, _ := json.Marshal(book)
	post_req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(json_data))
	req, _ := http.NewRequest("GET", "/books/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, post_req)
	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	suite.Contains(string(responseData), "sample")
	suite.Equal(http.StatusOK, w.Code)
}

func (suite *bookhandlerSuite) TestGetBooks() {
	book := models.Book{Title: "sample", Description: "sample"}

	router.GET("/books", handler.GetBooks)
	json_data, _ := json.Marshal(book)
	post_req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(json_data))
	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, post_req)
	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	suite.Contains(string(responseData), "sample")
	suite.Equal(http.StatusOK, w.Code)
}

func (suite *bookhandlerSuite) TestUpdateBook() {
	book := models.Book{Title: "sample", Description: "sample"}
	book1 := models.Book{Title: "new sample", Description: "new sample"}

	router.PATCH("/books/:id", handler.UpdateBook)
	new_json_data, _ := json.Marshal(book)
	updated_json_data, _ := json.Marshal(book1)
	post_req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(new_json_data))
	// req, _ := http.NewRequest("GET", "/books/1", nil)
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w1, post_req)

	r := make(map[string] map[string] interface{})

	new_book_response, _ := io.ReadAll(w1.Body)
	err := json.Unmarshal(new_book_response, &r)
	if err != nil {
		panic(err.Error())
	}
	id := r["data"]["ID"]
	_id := fmt.Sprintf("%v", id)
	put_req, _ := http.NewRequest("PATCH", "/books/" + _id, bytes.NewBuffer(updated_json_data))
	// suite.Contains(id, "new sample")

	router.ServeHTTP(w2, put_req)
	// router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w2.Body)
	suite.Contains(string(responseData), "new sample")
	suite.Equal(http.StatusOK, w2.Code)
}

func connectTestDatabase() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("../db/books_test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&models.Book{})
	if err != nil {
		return nil
	}

	return database
}
