package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AyoOluwa-Israel/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/gorm"
)

func root(res http.ResponseWriter, req *http.Request) {
	resp := map[string]interface{}{
		"message": "Welcome to my GO Bookstore",
		"status":  http.StatusOK,
	}

	jsonData, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	res.Write(jsonData)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", root).Methods("GET")
	s := r.PathPrefix("/api/v1").Subrouter()

	routes.RegisterBookStore(s)
	r.HandleFunc("/", root).Methods("GET")
	http.Handle("/", r)

	fmt.Printf("Starting Server at port 8080 and checking.. \n")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
