package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Actor struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	DOB    string `json:"dob"`
}

var actors = []Actor{}

func main() {
	http.HandleFunc("/actors", handleConn)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleConn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getActors(w, r)
	case http.MethodPost:
		newActor(w, r)
	case http.MethodDelete:
		deleteActor(w, r)
	case http.MethodPut:
		updateActor(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func getActors(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(actors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func newActor(w http.ResponseWriter, r *http.Request) {
	var newActor Actor
	err := json.NewDecoder(r.Body).Decode(&newActor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newActor.ID = len(actors) + 1
	actors = append(actors, newActor)
	w.WriteHeader(http.StatusCreated)
	_, err = fmt.Fprintf(w, "Actor created succesfully\n")
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
}

func deleteActor(w http.ResponseWriter, r *http.Request) {
	id, _ := getActorID(r)
	for i, actor := range actors {
		if actor.ID == id {
			actors = append(actors[:i], actors[i+1:]...)
			fmt.Fprintf(w, "Actor deleted successfully")
			return
		}
	}
	http.Error(w, "Actor not found", http.StatusNotFound)
}

func updateActor(w http.ResponseWriter, r *http.Request) {
	var updatedActor Actor
	err := json.NewDecoder(r.Body).Decode(&updatedActor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, _ := getActorID(r)
	for i, actor := range actors {
		if actor.ID == id {
			actors[i] = updatedActor
			fmt.Fprintf(w, "Actor updated successfully")
			return
		}
	}
	http.Error(w, "Actor not found", http.StatusNotFound)
}
func getActorID(r *http.Request) (int, error) {
	id := r.URL.Path[len("/actors/"):]
	return strconv.Atoi(id)
}
