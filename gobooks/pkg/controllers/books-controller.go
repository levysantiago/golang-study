package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lvy/gobooks/pkg/models"
	"github.com/lvy/gobooks/pkg/utils"
)

var NewBook models.Book

func GetBook(res http.ResponseWriter, req *http.Request){
	newBooks := models.GetAllBooks()
	data, _ := json.Marshal(newBooks)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func GetBookById(res http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	bookId := params["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if(err != nil){
		fmt.Println("Error while parsing")
	}

	bookDetails, _ := models.GetBookById(id)
	
	data, _ := json.Marshal(bookDetails)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func CreateBook(res http.ResponseWriter, req *http.Request){
	createBookData := &models.Book{}

	utils.ParseBody(req, createBookData)
	book := createBookData.CreateBook()
	
	data, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func DeleteBook(res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if(err != nil){
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(id)
	
	data, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func UpdateBook(res http.ResponseWriter, req *http.Request){
	var updateBook = &models.Book{}

	utils.ParseBody(req, updateBook)
	vars := mux.Vars(req)	
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if(err != nil){
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookById(id)
	if(updateBook.Name != ""){
		bookDetails.Name = updateBook.Name
	}
	if(updateBook.Author != ""){
		bookDetails.Author = updateBook.Author
	}
	if(updateBook.Publication != ""){
		bookDetails.Publication = updateBook.Publication
	}

	book := models.UpdateBook(id, *bookDetails)
	
	data, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}