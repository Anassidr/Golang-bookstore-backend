package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anassidr/go-bookstore/pkg/models"
	"github.com/anassidr/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0) //convert string to value (string to conv, base int, bitSize)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook) //extract data from the request body and store it in the CreateBook object
	b := CreateBook.CreateBook()   //The CreateBook object's method "CreateBook" is called to persist the book data in the database
	res, _ := json.Marshal(b)      // book object is serialized
	w.WriteHeader(http.StatusOK)   // http response is written
	w.Write(res)                   // response body will contain a JSON of the created book
}

func DeleteBook(w h)
