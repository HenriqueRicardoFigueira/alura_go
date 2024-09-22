package controllers

import (
	"alura_go/database"
	"alura_go/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var allPersonalities []models.Personality

	database.DB.Find(&allPersonalities)
	json.NewEncoder(w).Encode(allPersonalities)
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality

	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}
