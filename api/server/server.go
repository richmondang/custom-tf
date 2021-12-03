package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)



type Service struct {
	connectionString string
	volumes			map[string]Volume
	sync.RWMutex
}


func NewService(connectionString string, volumes map[string]Volume) *Service {
	return &Service{
		connectionString: connectionString,
		volumes:            volumes,
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

	//Create Volume
	r.HandleFunc("/volume", logs(auth(s.CreateVolume))).Methods("POST")
	//Get Volumes info
	r.HandleFunc("/volume", logs(auth(s.GetVolumes))).Methods("GET")
	//Get volume by ID
	r.HandleFunc("/volume/{vol_id}", logs(auth(s.GetVolume))).Methods("GET")
	//Update volume info by ID
	r.HandleFunc("/volume/{vol_id}", logs(auth(s.UpdateVolume))).Methods("PUT")
	//Delete volume by ID
	r.HandleFunc("/volume/{vol_id}", logs(auth(s.DeleteVolume))).Methods("DELETE")


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
