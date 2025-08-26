package mqtt

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/db"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Variável global do cliente MQTT
var Client mqtt.Client

// Callback executado ao receber uma mensagem
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("[MQTT] Mensagem recebida em %s: %s\n", msg.Topic(), msg.Payload())
}

// Função para conectar ao broker MQTT
func ConnectMQTT(topicos []db.Topic) {
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
		log.Println("[MQTT] Conectado ao broker:", broker)

		// Inscrever em todos os tópicos vindos do banco
		for _, t := range topicos {
			if token := c.Subscribe(t.Nome, 1, nil); token.Wait() && token.Error() != nil {
				log.Printf("[MQTT] Erro ao se inscrever em %s: %v\n", t.Nome, token.Error())
			} else {
				log.Printf("[MQTT] Inscrito no tópico: %s\n", t.Nome)
			}
		}
	}

	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Println("[MQTT] Conexão perdida:", err)
	}

	Client = mqtt.NewClient(opts)
	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln("[MQTT] Erro ao conectar:", token.Error())
	}
}


