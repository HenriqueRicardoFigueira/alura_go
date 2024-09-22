package main

import (
	"alura_go/database"
	"alura_go/routes"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	database.ConnectDb()

	fmt.Println("Start server on localhost:8000")
	routes.HandleRequests()
}
