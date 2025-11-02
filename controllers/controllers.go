package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GuilhermeLatansa/Go-Api-Rest/database"
	"github.com/GuilhermeLatansa/Go-Api-Rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ExibePaises(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p []models.Pais
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmPais(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var pais models.Pais
	database.DB.First(&pais, id)
	json.NewEncoder(w).Encode(pais)
}

func EnviaNovoPais(w http.ResponseWriter, r *http.Request) {
	var novoPais models.Pais
	json.NewDecoder(r.Body).Decode(&novoPais)
	database.DB.Create(&novoPais)
	json.NewEncoder(w).Encode(novoPais)
}

func DeletaPais(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var pais models.Pais
	database.DB.Delete(&pais, id)
	json.NewEncoder(w).Encode(pais)
}

func EditaPais(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var pais models.Pais
	database.DB.First(&pais, id)
	json.NewDecoder(r.Body).Decode(&pais)
	database.DB.Save(&pais)
	json.NewEncoder(w).Encode(pais)
}
