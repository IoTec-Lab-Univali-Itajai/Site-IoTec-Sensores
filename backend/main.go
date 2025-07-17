package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/mqtt"
)

func main() {
	// Carregar .env (se estiver usando)
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("⚠️  Arquivo .env não encontrado, prosseguindo sem ele")
	}

	// Exemplo de tópicos
	topics := []string{"iot/lab/sala1", "iot/sala2"}

	// Conectar e escutar MQTT
	mqtt.ConnectMQTT(topics)

	// Bloquear a main para manter o programa rodando
	select {}
}
