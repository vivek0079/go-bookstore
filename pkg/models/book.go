package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-bookstore/pkg/configs"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name" json:"name,omitempty"`
	Author      string `json:"author,omitempty"`
	Publication string `json:"publication,omitempty"`
}

func init() {
	configs.Connect()
	db = configs.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	result := db.Find(&books)
	fmt.Println(result.RowsAffected)
	return books
}

func GetBookById(id int64) *Book {
	var book Book
	db.Where("id = ?", id).Find(&book)
	return &book
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("id = ?", id).Delete(book)
	return book
}