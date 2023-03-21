package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/AyoOluwa-Israel/bookstore/pkg/models"
	"github.com/AyoOluwa-Israel/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)


var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := map[string]interface{}{
		"message": "Retrieved Successfully",
		"dateValue": time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		"payload": newBooks,
		"status": http.StatusOK,
	}


	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	w.Write(jsonData)

}


func GetBooksById (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error in parsing")
	}
	bookDetails, _ := models.GetBookById(Id)
	resp := map[string]interface{}{
		"message": "Retrieved Successfully",
		"dateValue": time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		"payload": bookDetails,
		"status": http.StatusOK,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	newBook := CreateBook.CreateBook()
	resp := map[string]interface{}{
		"message": "Retrieved Successfully",
		"dateValue": time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		"payload": newBook,
		"status": http.StatusOK,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}


func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("Error in parsing")
	}
	book := models.DeleteBook(Id)
	resp := map[string]interface{}{
		"message": "Retrieved Successfully",
		"dateValue": time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		"payload": book,
		"status": http.StatusOK,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}