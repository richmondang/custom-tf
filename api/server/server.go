package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)


// Service holds the map of items and provides methods CRUD operations on the map
type Service struct {
	connectionString string
	items            map[string]Item
	volumes			map[string]Volume
	sync.RWMutex
}

// NewService returns a Service with a connectionString configured and can be a map of items setup. The items map can be empty,
// or can contain items
func NewService(connectionString string, items map[string]Item) *Service {
	return &Service{
		connectionString: connectionString,
		items:            items,
	}
}


// func runServer(connection string) error {

// 	router := mux.NewRouter()
// 	router.HandleFunc("/volume", logs(auth(getVolume)))

// 	// start server
// 	log.Println("Starting API server at ", connection)
// 	return http.ListenAndServe(connection, router)
// }




// ListenAndServe registers the routes to the server and starts the server on the host:port configured in Service
func (s *Service) ListenAndServe() error {
	r := mux.NewRouter()

	// Each handler is wrapped in logs() and auth() to log out the method and path and to
	// ensure that a non-empty Authorization header is present
	r.HandleFunc("/item", logs(auth(s.PostItem))).Methods("POST")
	r.HandleFunc("/item", logs(auth(s.GetItems))).Methods("GET")
	r.HandleFunc("/item/{name}", logs(auth(s.GetItem))).Methods("GET")
	r.HandleFunc("/item/{name}", logs(auth(s.PutItem))).Methods("PUT")
	r.HandleFunc("/item/{name}", logs(auth(s.DeleteItem))).Methods("DELETE")


	//Volume Routes
	//Update Volume
	r.HandleFunc("/volume", logs(auth(s.CreateVolume))).Methods("POST")
	//Get Storage volume information
	r.HandleFunc("/volume", logs(auth(s.GetVolumeInfo))).Methods("GET")
	//Get Volume information by ID
	// r.HandleFunc("/volume/{vol_id}", logs(auth(s.GetItem))).Methods("GET")
	// //Update volume by ID
	// r.HandleFunc("/volume/{vol_id}", logs(auth(s.PutItem))).Methods("PUT")
	// //Delete storage volume by ID
	// r.HandleFunc("/volume/{vol_id}", logs(auth(s.DeleteItem))).Methods("DELETE")


	log.Printf("Starting server on %s", s.connectionString)
	err := http.ListenAndServe(s.connectionString, r)
	if err != nil {
		return err
	}
	return nil
}

// logs prints the Method and Path to stdout
func logs(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path
		log.Printf("%s %s", method, path)
		handlerFunc(w, r)
		return
	}
}

// auth checks that a non-empty authorization header has been sent with the request
func auth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Authorization token required", http.StatusUnauthorized)
			return
		}
		handlerFunc(w, r)
		return
	}
}