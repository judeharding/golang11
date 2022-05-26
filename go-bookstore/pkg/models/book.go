package models

// GORM = A GOLANG Object-Relational Mapper (ORM) is a code
// library that automates the transfer of data stored in
// relational database tables into objects that are more
// commonly used in application code.

import (
	"github.com/akhil/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB


// a struct is a model to help us store stuff in a db
type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

//helps us initialize and connect to the db 
func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}


// model functions to talk to a db 
func (b *Book) CreateBook() *Book{
	db.NewRecord(b)  //NewRecord is a GORM func
	db.Create(&b)    //Create is a GORM func
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)  //Find is a GORM func
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.Where("ID=?", Id).Find(&getBook)  //Where is a GORM func
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book) //Delete is a GORM func
	return book
} 

// no update function because we are going to delete and post a new (same) book