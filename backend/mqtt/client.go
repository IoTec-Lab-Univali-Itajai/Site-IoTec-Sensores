package mqtt

import (
	"fmt"
	"log"
	"os"
	"time"
	"strings"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Variável global do cliente MQTT
var Client mqtt.Client

// Callback executado ao receber uma mensagem
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("[MQTT] Mensagem recebida em %s: %s\n", msg.Topic(), msg.Payload())
}

// Função para conectar ao broker MQTT
// Função para conectar ao broker MQTT
func ConnectMQTT(topicosString string) {
	opts := mqtt.NewClientOptions()

	// Configurações do broker
	broker := os.Getenv("MQTT_BROKER_URL")
	clientID := os.Getenv("MQTT_CLIENT_ID")
	username := os.Getenv("MQTT_USERNAME")
	password := os.Getenv("MQTT_PASSWORD")

	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetCleanSession(true)
	opts.SetProtocolVersion(4)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(5 * time.Second)
	opts.SetDefaultPublishHandler(messageHandler)

	opts.OnConnect = func(c mqtt.Client) {
		log.Println("client.go diz: [MQTT] Conectado ao broker:", broker)

		// Divide a string de tópicos separados por espaço
		topicos := strings.Split(topicosString, " ")
		
		// Inscrever em todos os tópicos
		for _, topico := range topicos {
			topico = strings.TrimSpace(topico)
			if topico == "" {
				continue // Pula tópicos vazios
			}
			
			if token := c.Subscribe(topico, 1, nil); token.Wait() && token.Error() != nil {
				log.Printf("client.go diz: [MQTT] Erro ao se inscrever em %s: %v\n", topico, token.Error())
			} else {
				log.Printf("client.go diz: [MQTT] Inscrito no tópico: %s\n", topico)
			}
		}
	}

	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Println("client.go diz: [MQTT] Conexão perdida:", err)
	}

	Client = mqtt.NewClient(opts)
	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln("client.go diz: [MQTT] Erro ao conectar:", token.Error())
	}
}


