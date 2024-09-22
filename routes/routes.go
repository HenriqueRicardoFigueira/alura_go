package routes

import (
	"alura_go/controllers"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.GetAllPersonalities)
}
