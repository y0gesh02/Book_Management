package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/y0gesh02/go-bookstore/pkg/models"
	"github.com/y0gesh02/go-bookstore/pkg/utils"
)

//utils for json unmarshling **create and update**
//models for dbconnection and struct and db func

var NewBook models.Book  //creating var newbook of type struct book 

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:=models.GetAllBooks()  //geeting list of all books form db
	res, _ :=json.Marshal(newBooks)  //converting data to json
	w.Header().Set("Content-Type","pkglication/json")  //
	w.WriteHeader(http.StatusOK) //200ok
	w.Write(res) //returning response
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)  //to get access to r
	bookId := vars["bookId"]  //getting id
	ID, err:= strconv.ParseInt(bookId,0,0) //conveting to int
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _:= models.GetBookById(ID) //geeting a book 
	res, _ := json.Marshal(bookDetails) 
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{} //getting struct func ,, pointing to book struct
	utils.ParseBody(r, CreateBook) //receiving json. josn to db lang
	b:= CreateBook.CreateBook() //calling func
	res, _ := json.Marshal(b) //data to json
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
//same as getbookbyid
func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//get bookbyid match details&updte save to db
func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}  //pointing to book struct
	utils.ParseBody(r, updateBook)  //receiving json. josn to db lang
	vars := mux.Vars(r)
	bookId := vars["bookId"] 
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db:=models.GetBookById(ID) //getting update book
	//updating details
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}