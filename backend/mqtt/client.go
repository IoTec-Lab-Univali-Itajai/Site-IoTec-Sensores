package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/data"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Variável global do cliente MQTT
var Client mqtt.Client

// Handler MQTT (cada mensagem é processada em goroutine separada)
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	go func() {
		fmt.Printf("[MQTT] Mensagem recebida em %s\n", msg.Topic())

		ttnMsg, err := ParseCompleteTTNMessage(msg.Payload())
		if err != nil {
			fmt.Println("Erro ao processar mensagem:", err)
			return
		}

		fmt.Printf("Dispositivo: %s\n", ttnMsg.EndDeviceIDs.DeviceID)
		fmt.Printf("Recebido em: %s\n", ttnMsg.ReceivedAt)
		fmt.Printf("Dados dos sensores: %+v\n", ttnMsg.SensorData)

		jsonData, err := json.MarshalIndent(ttnMsg, "", "  ")
		if err != nil {
			fmt.Println("Erro ao converter para JSON:", err)
			return
		}

		fmt.Println("Mensagem completa em JSON:")
		fmt.Println(string(jsonData))

		// Aqui você poderia salvar no MongoDB, enviar pro frontend, etc.
	}()
}

// Função para parsear payload dos sensores
func ParseSensorPayload(payload string) ([]data.SensorValue, error) {
	parts := strings.Split(payload, ",")
	if len(parts)%3 != 0 {
		return nil, fmt.Errorf("client.go diz: payload inválido: número de elementos não é múltiplo de 3")
	}

	var results []data.SensorValue
	for i := 0; i < len(parts); i += 3 {
		val, err := strconv.ParseFloat(parts[i+1], 64)
		if err != nil {
			return nil, fmt.Errorf("client.go diz: erro ao converter valor '%s': %w", parts[i+1], err)
		}
		results = append(results, data.SensorValue{
			InfoType: parts[i],
			Valor:    val,
			Unidade:  parts[i+2],
		})
	}
	return results, nil
}

// Função para parsear a mensagem TTN completa
func ParseCompleteTTNMessage(payload []byte) (*data.TTNMessage, error) {
	var ttnMsg data.TTNMessage

	// Decodifica JSON bruto
	err := json.Unmarshal(payload, &ttnMsg)
	if err != nil {
		return nil, fmt.Errorf("client.go diz: erro ao decodificar JSON TTN: %w", err)
	}

	// Parseia payload dos sensores
	sensorData, err := ParseSensorPayload(ttnMsg.UplinkMessage.DecodedPayload.Message)
	if err != nil {
		return nil, fmt.Errorf("client.go diz: erro ao parsear payload dos sensores: %w", err)
	}

	ttnMsg.SensorData = sensorData
	return &ttnMsg, nil
}


// Função para conectar ao broker MQTT e se inscrever nos tópicos
func ConnectMQTT(topicosString string) {
	opts := mqtt.NewClientOptions()

	// Configurações do broker via env
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

		topicos := strings.Split(topicosString, " ")
		for _, topico := range topicos {
			topico = strings.TrimSpace(topico)
			if topico == "" {
				continue
			}
			if token := c.Subscribe(topico, 1, nil); token.Wait() && token.Error() != nil {
				log.Printf("[MQTT] Erro ao se inscrever em %s: %v\n", topico, token.Error())
			} else {
				log.Printf("[MQTT] Inscrito no tópico: %s\n", topico)
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
