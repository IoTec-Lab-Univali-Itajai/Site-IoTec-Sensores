package mqtt

import(
	"fmt"
	"log"
	"strings"
	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/db"
)

func BuscarConectarTopicos(){
	// Busca tópicos do banco (como slice de Topic)
	topicos, err := db.BuscarTopicos()
	if err != nil {
		log.Fatal("topicos.go diz: Erro ao buscar tópicos:", err)
	}
	fmt.Println("topicos.go diz: Tópicos encontrados: ", topicos)

	// Converte slice de Topic para string separada por espaços
	var topicosNomes []string
	for _, topico := range topicos {
		topicosNomes = append(topicosNomes, topico.Nome)
	}
	topicosString := strings.Join(topicosNomes, " ")
	
	fmt.Println("topicos.go diz: String de tópicos:", topicosString)

	DisconnectMQTT();

	// Chama ConnectMQTT com a string de tópicos
	ConnectMQTT(topicosString)
}
