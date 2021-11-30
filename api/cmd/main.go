package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/richmondang/terraform-example-apex/api/pkg/handlers"
)

func main() {
	router := mux.NewRouter()

	// router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
	// 	json.NewEncoder(w).Encode("Hello World")
	// })

	router.HandleFunc("/subscription/volume", handlers.GetAllVolumes).Methods(http.MethodGet)

	log.Println("Webserver running")
	http.ListenAndServe(":4000", router)
}
