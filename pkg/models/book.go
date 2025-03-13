package models

import (
	"github.com/hamza-s47/book-store/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `gorm:"default:'Unknown'" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}

// func UpdateBook(Id int64, updatedBook Book) (*Book, *gorm.DB) {
// 	var book Book

// 	db := db.Where("ID=?", Id).First(&book)
// 	if db.RecordNotFound() {
// 		return nil, db
// 	}
// 	book.Name = updatedBook.Name
// 	book.Author = updatedBook.Author
// 	book.Publication = updatedBook.Publication

// 	db.Save(&book)
// 	return &book, db
// }
