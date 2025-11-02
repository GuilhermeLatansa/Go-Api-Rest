package routes

import (
	"log"
	"net/http"

	"github.com/GuilhermeLatansa/Go-Api-Rest/controllers"
	"github.com/GuilhermeLatansa/Go-Api-Rest/middleware"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentType)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/paises", controllers.ExibePaises).Methods("Get")
	r.HandleFunc("/api/paises/{id}", controllers.RetornaUmPais).Methods("Get")
	r.HandleFunc("/api/paises", controllers.EnviaNovoPais).Methods("Post")
	r.HandleFunc("/api/paises/{id}", controllers.DeletaPais).Methods("Delete")
	r.HandleFunc("/api/paises/{id}", controllers.EditaPais).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", r))
}
