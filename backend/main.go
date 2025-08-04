package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/db"
	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/mqtt"
	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/api"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("⚠️  Arquivo .env não encontrado, prosseguindo sem ele")
	}

	// Conecta no banco de dados
	db.ConnectMongoDB()

	// Exemplo: escutar tópicos MQTT (em produção seria mais dinâmico)
	topics := []string{"iot/lab/sala1", "iot/sala2"}
	mqtt.ConnectMQTT(topics)

	// Inicia o servidor de API
	api.StartAPI() // ← AQUI! Isso roda o servidor em :8080
}
