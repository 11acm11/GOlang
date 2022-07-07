package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Standing struct {
	Pos    int
	Name   string
	Team   string
	Points int
}

var standings = []Standing{
	{1, "Max Verstappen", "Red Bull Racing", 181},
	{2, "Sergio Perez", "Red Bull Racing", 147},
	{3, "Charles Leclerc", "Ferrari", 138},
	{4, "Carlos Sainz", "Ferrari", 127},
	{5, "Geroge Russel", "Mercedes", 111},
	{6, "Lewis Hamilton", "Mercedes", 93},
	{7, "Lando Norris", "Mclaren", 58},
	{8, "Valtteri Bottas", "Alpha Romeo", 46},
	{9, "Estaban Ocon", "Alpine", 39},
	{10, "Fernando Alonso", "Alpine", 28},
}

func returnAllStandings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(standings)
}
func returnStandingsByTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	vars := mux.Vars(r)
	standT := vars["team"]
	stands := &[]Standing{}
	for _, stand := range standings {
		if stand.Team == standT {
			*stands = append(*stands, stand)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stands)
}

func returnStandingsByPos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	vars := mux.Vars(r)
	standP, err := strconv.Atoi(vars["pos"])
	if err != nil {
		fmt.Println("Unable to convert to string")
	}
	for _, stand := range standings {
		if stand.Pos == standP {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(stand)
		}
	}
}

func createStandings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	var newStand Standing
	json.NewDecoder(r.Body).Decode(&newStand)
	standings = append(standings, newStand)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(standings)
}

func removeStandings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	vars := mux.Vars(r)
	standP, err := strconv.Atoi(vars["pos"])
	if err != nil {
		fmt.Println("Unable to convert to string")
	}
	for k, stand := range standings {
		if stand.Pos == standP {
			standings = append(standings[:k], standings[k+1:]...)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(standings)
}

func updateStandings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	vars := mux.Vars(r)
	standP, err := strconv.Atoi(vars["pos"])
	if err != nil {
		fmt.Println("Unable to convert to string")
	}
	var updatedstand Standing
	json.NewDecoder(r.Body).Decode(&updatedstand)
	for k, stand := range standings {
		if stand.Pos == standP {
			standings = append(standings[:k], standings[k+1:]...)
			standings[k] = updatedstand
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(standings)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/standings", returnAllStandings).Methods("GET")
	router.HandleFunc("/standings/{team}", returnStandingsByTeam).Methods("GET")
	router.HandleFunc("/standings/pos/{pos}", returnStandingsByPos).Methods("GET")
	router.HandleFunc("/standings/pos/{pos}", updateStandings).Methods("PUT")
	router.HandleFunc("/standings", createStandings).Methods("POST")
	router.HandleFunc("/standings/pos/{pos}", removeStandings).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
