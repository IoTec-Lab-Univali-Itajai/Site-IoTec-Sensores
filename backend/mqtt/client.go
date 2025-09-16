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
	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/db"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Variável global do cliente MQTT
var Client mqtt.Client

// Handler MQTT (cada mensagem é processada em goroutine separada)
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	go func() {
		fmt.Printf("[MQTT] Mensagem recebida em %s\n", msg.Topic())
		fmt.Printf("\n=====================mensagem: ============================================\n%s\n===============================================================\n", string(msg.Payload()))
		// Extrai apenas o necessário do JSON
		var message struct {
			DeviceInfo struct {
				DeviceName string `json:"deviceName"`
			} `json:"deviceInfo"`
			Object struct {
				Message string `json:"message"`
			} `json:"object"`
		}

		if err := json.Unmarshal(msg.Payload(), &message); err != nil {
			fmt.Println("client.go diz: Erro ao decodificar JSON:", err)
			return
		}

		// Parseia os dados dos sensores
		sensorData, err := ParseSensorPayload(message.Object.Message)
		if err != nil {
			fmt.Println("client.go diz: Erro ao parsear sensores:", err)
			return
		}

		fmt.Printf("client.go diz: Dispositivo: %s\n", message.DeviceInfo.DeviceName)
		fmt.Printf("client.go diz: Dados: %+v\n", sensorData)

		// Atualiza direto no data.DadosDisplay
		data.DadosDisplayMutex.Lock()
		
		// Procura se já existe
		for i, display := range data.DadosDisplay {
			if display.SensorID == message.DeviceInfo.DeviceName {
				data.DadosDisplay[i].SensorData = sensorData
				fmt.Printf("client.go diz: DadosDisplay atualizado para %s\n", message.DeviceInfo.DeviceName)
				break
			}
		}

		data.DadosDisplayMutex.Unlock()

		//busca se está cadastrado no BD e atualiza os dados
		if db.SensorExiste(message.DeviceInfo.DeviceName){
			// Atualiza os dados do sensor
			err := db.AtualizarDados(message.DeviceInfo.DeviceName, sensorData)
			if err != nil {
				fmt.Println("client.go diz: Erro ao atualizar dados:", err)
			}
			
			// Insere no histórico de dados (se necessário)
			dadoNovo := data.InfoDisplay{
				SensorID:   message.DeviceInfo.DeviceName,
				SensorData: sensorData,
			}
			err = db.InserirDado(dadoNovo) // ← Agora correto!
			if err != nil {
				fmt.Println("client.go diz: Erro ao inserir dado:", err)
			}
		}

		
		
	}()
}

// Função para parsear payload dos sensores
func ParseSensorPayload(payload string) ([]data.SensorValue, error) {
	parts := strings.Split(payload, ",")
	if len(parts)%3 != 0 {
		return nil, fmt.Errorf("client.go diz: payload inválido")
	}

	var results []data.SensorValue
	for i := 0; i < len(parts); i += 3 {
		val, err := strconv.ParseFloat(parts[i+1], 64)
		if err != nil {
			return nil, err
		}
		results = append(results, data.SensorValue{
			InfoType: parts[i],
			Valor:    val,
			Unidade:  parts[i+2],
		})
	}
	return results, nil
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
