package controllers

import (
	"alura_go/database"
	"alura_go/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality

	database.DB.Delete(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality

	database.DB.First(&personality, id)
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Save(&personality)

	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)

	database.DB.Create(&personality)

	json.NewEncoder(w).Encode(personality)
}
