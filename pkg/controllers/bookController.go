package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AyoOluwa-Israel/bookstore/pkg/models"
	"github.com/AyoOluwa-Israel/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

type ArrayResponse struct {
	Message    string        `json:"message"`
	Status     string        `json:"status"`
	StatusCode int           `json:"statusCode"`
	Data       []models.Book `json:"payload"`
}

type Response struct {
	Message    string      `json:"message"`
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"payload"`
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(&ArrayResponse{"All Books Retrieved.", "success", 200, newBooks})
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	w.Write(jsonData)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error in parsing")
	}
	bookDetails, _ := models.GetBookById(Id)
	jsonData, err := json.Marshal(&Response{"Book Retrieved.", "success", 200, bookDetails})

	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	newBook := CreateBook.CreateBook()

	jsonData, err := json.Marshal(&Response{"Book Created Successfully.", "success", 200, newBook})

	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error in parsing")
	}
	book := models.DeleteBook(Id)
	jsonData, err := json.Marshal(&Response{"Book Deleted Successfully.", "success", 200, book})

	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}


func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	w.Header().Set("Content-Type", "application/json")

	bookDetails, db  := models.GetBookById(Id)

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
	jsonData, err := json.Marshal(&Response{"Book Deleted Successfully.", "success", 200, bookDetails})

	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	w.Write(jsonData)
}