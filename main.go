package main

import (
	"alura_go/models"
	"alura_go/routes"
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Carl Jung", History: "Carl Gustav Jung was a Swiss psychiatrist and psychoanalyst who founded analytical psychology."},
		{Id: 2, Name: "Sigmund Freud", History: "Sigmund Freud was an Austrian neurologist and the founder of psychoanalysis."},
	}

	fmt.Println("Start server on localhost:8000")
	routes.HandleRequests()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
