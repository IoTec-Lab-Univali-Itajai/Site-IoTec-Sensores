package main

import (
	"fmt"
	"log"

	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/api"
	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/db"
	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/mqtt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("⚠️  Arquivo .env não encontrado, prosseguindo sem ele")
	}

	// Conecta no banco de dados
	db.ConnectMongoDB();

	// Conecta ao broker MQTT e escuta os tópicos configurados no .env
	topicos, err := db.BuscarTopicos() // ← CORREÇÃO AQUI: duas variáveis
	if err != nil {
		log.Fatal("Erro ao buscar tópicos:", err)
	}
	fmt.Println("topicos encontrados: ", topicos);
	
	mqtt.ConnectMQTT(topicos);

	// Inicia o servidor de API
	api.StartAPI(); // servidor em :8080
}
