package models

import (
	"github.com/lvy/gobooks/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:""json:"Name"`
	Author string `json:"Author"`
	Publication string `json:"Publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()

	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book{
	db.Create(&book)
	return book
}

func GetAllBooks() []Book{
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func DeleteBook(Id int64) Book{
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}

func UpdateBook(Id int64, book Book) Book{
	db.Where("ID=?", Id).Updates(&book)
	return book
}