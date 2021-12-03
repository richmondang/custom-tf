package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"regexp"
)

// represents a single Volume
type Volume struct {
	ID                       string               `json:"id,omitempty"`
	Name                     string               `json:"name,omitempty"`
	Description              string               `json:"description,omitempty"`
	ApplianceID              string               `json:"appliance_id,omitempty"`
	Size                     int                  `json:"size,omitempty"`
}

// Get all volumes
func (s *Service) GetVolumes(w http.ResponseWriter, r *http.Request) {
	s.RLock()
	defer s.RUnlock()

	err := json.NewEncoder(w).Encode(s.volumes)
	if err != nil {
		log.Println(err)
	}
}

// Create Volume
func (s *Service) CreateVolume(w http.ResponseWriter, r *http.Request) {
	var volume Volume
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&volume)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(volume.ID)) {
		http.Error(w, "Invalid volume ID", 400)
		return
	}

	s.Lock()
	defer s.Unlock()

	if s.volumeExists(volume.ID) {
		http.Error(w, fmt.Sprintf("volume %s already exists", volume.ID), http.StatusBadRequest)
		return
	}

	s.volumes[volume.ID] = volume
	log.Printf("added volume - ID: %s", volume.ID)
	err = json.NewEncoder(w).Encode(volume)
	if err != nil {
		log.Printf("error sending response - %s", err)
	}
}

// Update volume by ID
func (s *Service) UpdateVolume(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volumeName := vars["vol_id"]
	if volumeName == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	var volume Volume
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&volume)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	s.Lock()
	defer s.Unlock()

	if !s.volumeExists(volumeName) {
		log.Printf("volume %s does not exist", volumeName)
		http.Error(w, fmt.Sprintf("volume %v does not exist", volumeName), http.StatusBadRequest)
		return
	}

	s.volumes[volumeName] = volume
	log.Printf("updated volume: %s", volume.ID)
	err = json.NewEncoder(w).Encode(volume)
	if err != nil {
		log.Printf("error sending response - %s", err)
	}
}

// Delete volume by ID
func (s *Service) DeleteVolume(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volumeName := vars["vol_id"]
	if volumeName == "" {
		http.Error(w, "volume not found", http.StatusNotFound)
		return
	}
	s.Lock()
	defer s.Unlock()

	if !s.volumeExists(volumeName) {
		http.Error(w, fmt.Sprintf("volume %s does not exist", volumeName), http.StatusNotFound)
		return
	}

	delete(s.volumes, volumeName)

	_, err := fmt.Fprintf(w, "Deleted volume - ID: %s", volumeName)
	if err != nil {
		log.Println(err)
	}
}

// Get specific volume by ID
func (s *Service) GetVolume(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volumeName := vars["vol_id"]
	if volumeName == "" {
		http.Error(w, "volume not found", http.StatusNotFound)
		return
	}

	s.RLock()
	defer s.RUnlock()
	if !s.volumeExists(volumeName) {
		http.Error(w, "volume not found", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(s.volumes[volumeName])
	if err != nil {
		log.Println(err)
		return
	}
}

// Check if volume exists
func (s *Service) volumeExists(volumeName string) bool {
	if _, ok := s.volumes[volumeName]; ok {
		return true
	}
	return false
}


