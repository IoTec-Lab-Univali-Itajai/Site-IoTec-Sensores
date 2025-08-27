package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/api"
	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/db"
	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/mqtt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Main.go diz: ⚠️  Arquivo .env não encontrado, prosseguindo sem ele")
	}

	// Conecta no banco de dados
	db.ConnectMongoDB()

	// Busca tópicos do banco (como slice de Topic)
	topicos, err := db.BuscarTopicos()
	if err != nil {
		log.Fatal("Main.go diz: Erro ao buscar tópicos:", err)
	}
	fmt.Println("Main.go diz: Tópicos encontrados: ", topicos)

	// Converte slice de Topic para string separada por espaços
	var topicosNomes []string
	for _, topico := range topicos {
		topicosNomes = append(topicosNomes, topico.Nome)
	}
	topicosString := strings.Join(topicosNomes, " ")
	
	fmt.Println("Main.go diz: String de tópicos:", topicosString)
	
	// Chama ConnectMQTT com a string de tópicos
	mqtt.ConnectMQTT(topicosString)

	// Inicia o servidor de API
	api.StartAPI() // servidor em :8080
}