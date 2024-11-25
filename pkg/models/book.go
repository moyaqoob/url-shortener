package models

import (
	"bookstore-go/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB
type Book struct{
	gorm.Model
	Name string `gorm:""json:"id"`
	Author string `json:"email"`
	Publication int `json:"Publication"`
}



func init(){
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	result:= db.Create(&b)
	if result.Error != nil{
		return nil
	}
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(ID int64) (*Book, *gorm.DB){
	var getBook Book
	db.Where("ID:?",ID).Delete(&getBook)
	return &getBook,db
}

func DeleteBook(ID int64) (Book){
	var book Book
	db.Where("ID=?",ID).Delete(book)
	return book
}