package main

import (
	"fmt"

	"github.com/GuilhermeLatansa/Go-Api-Rest/database"
	"github.com/GuilhermeLatansa/Go-Api-Rest/models"
	"github.com/GuilhermeLatansa/Go-Api-Rest/routes"
)

func main() {

	models.Paises = []models.Pais{
		{Id: 1, Nome: "Pais 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Pais 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor na porta 8000")
	routes.HandleRequests()
}
