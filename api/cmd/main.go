package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"api/pkg/handlers"
	"api/pkg/mocks"
	"api/pkg/models"
)

func main() {
	router := mux.NewRouter()

	// router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
	// 	json.NewEncoder(w).Encode("Hello World")
	// })

	router.HandleFunc("/subscription/volume", handlers.GetAllVolumes).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
