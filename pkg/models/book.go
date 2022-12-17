package models

import (
	"github.com/jinzhu/gorm"
	"github.com/y0gesh02/go-bookstore/pkg/config"
)

var db *gorm.DB


//book struct 
type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

//connecting with databse
func init(){
	config.Connect()      //calling config connect func
	db = config.GetDB()   //getting db from config file
	db.AutoMigrate(&Book{})  //automigrated with a empty book
}

//to interact with database .each controllers has its own model

func (b *Book) CreateBook() *Book{
	db.NewRecord(b) //gorm will write query for inserting new record less coding
	db.Create(&b)  //creating new book
	return b
}

func GetAllBooks() []Book{
	var Books []Book  //slice for all books
	db.Find(&Books)   //finding books
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book  
	db:=db.Where("ID=?", Id).Find(&getBook) //finding bookbyid
	return &getBook, db  //book and db
	//return db for update book controller 
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

//dont have update func bcz we are getting id from db then updateid and save to d