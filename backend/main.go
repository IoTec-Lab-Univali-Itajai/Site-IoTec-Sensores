package main

import (
	"log"

	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/api"
	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/mqtt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Main.go diz: ⚠️  Arquivo .env não encontrado, prosseguindo sem ele")
	}

	mqtt.BuscarConectarTopicos();

	// Inicia o servidor de API
	api.StartAPI() // servidor em :8080
}