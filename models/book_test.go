package models

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type bookServiceSuite struct {
	suite.Suite
	service BookService
}

func (suite *bookServiceSuite) SetupSuite() {
	test_db := connectTestDatabase()
	s := InitBookService(test_db)

	suite.service = s

	// create new dummy data before each test
	book := Book{
		Title:       "sample title",
		Description: "sample description",
	}
	suite.service.CreateBook(book)
}

// func (suite *bookServiceSuite) SetupTest() {
// }

func (suite *bookServiceSuite) TearDownSuite() {
	// clear db after test suite is done
	// dropTable()
}

func (suite *bookServiceSuite) TestGetBooks() {
	books := suite.service.GetBooks()
	suite.True(len(*books) > 1)
}

func (suite *bookServiceSuite) TestGetBook() {
	book, _ := suite.service.GetBook("1")
	suite.NotNil(*book)

	_, err := suite.service.GetBook("100")
	suite.Error(err)
}

func (suite *bookServiceSuite) TestUpdateBook() {
	bookInput := Book{
		Title:       "sample title updated",
		Description: "sample description updated",
	}
	book, _ := suite.service.UpdateBook("1", bookInput)
	suite.Equal("sample title updated", book.Title)
	suite.Equal("sample description updated", book.Description)
}

func (suite *bookServiceSuite) TestCreateBook() {
	bookInput := Book{
		Title:       "new book",
		Description: "new sample",
	}
	book := suite.service.CreateBook(bookInput)
	suite.Equal("new book", book.Title)
	suite.Equal("new sample", book.Description)
}

// func (suite *bookServiceSuite) TestDeleteBook() {
// 	isDeleted, _ := suite.service.DeleteBook("1")
// 	suite.True(isDeleted)

// 	isNotDeleted, err := suite.service.DeleteBook("100")
// 	suite.Error(err)
// 	suite.False(isNotDeleted)
// }

func connectTestDatabase() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("../db/books_test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Book{})
	if err != nil {
		return nil
	}

	return database
}

func dropTable() {
	database, err := gorm.Open(sqlite.Open("../db/books_test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	_err := database.Migrator().DropTable(&Book{})
	if _err != nil {
		return
	}
}

func TestBookService(t *testing.T) {
	suite.Run(t, new(bookServiceSuite))
}
