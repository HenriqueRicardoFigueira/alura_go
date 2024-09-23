package routes

import (
	"alura_go/controllers"
	"alura_go/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/api/personalities", controllers.GetAllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("Put")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("Post")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
