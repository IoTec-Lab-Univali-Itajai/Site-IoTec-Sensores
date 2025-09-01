package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type SensorValue struct {
	infoType string  `json:"infoType"`
	Valor   float64 `json:"valor"`
	Unidade string  `json:"unidade"`
}

// Struct completa para representar a mensagem TTN com dados decodificados
type TTNMessage struct {
	EndDeviceIDs struct {
		DeviceID string `json:"device_id"`
	} `json:"end_device_ids"`
	ReceivedAt    string `json:"received_at"`
	UplinkMessage struct {
		FrmPayload     string `json:"frm_payload"`
		DecodedPayload struct {
			Message string `json:"message"`
		} `json:"decoded_payload"`
		// Você pode adicionar outros campos do uplink se necessário
	} `json:"uplink_message"`
	
	// Campo adicional para armazenar os dados dos sensores já parseados
	SensorData []SensorValue `json:"sensor_data,omitempty"`
}


// Função para parsear a mensagem TTN completa incluindo os sensores
func ParseCompleteTTNMessage(payload []byte) (*TTNMessage, error) {
	var ttnMsg TTNMessage
	
	// 1. Decodificar JSON bruto que veio do TTN
	err := json.Unmarshal(payload, &ttnMsg)
	if err != nil {
		return nil, fmt.Errorf("client.go diz: erro ao decodificar JSON TTN: %w", err)
	}
	
	// 2. Parsear o payload dos sensores e adicionar ao objeto
	sensorData, err := ParseSensorPayload(ttnMsg.UplinkMessage.DecodedPayload.Message)
	if err != nil {
		return nil, fmt.Errorf("client.go diz: erro ao parsear payload dos sensores: %w", err)
	}
	
	ttnMsg.SensorData = sensorData
	return &ttnMsg, nil
}

func ParseSensorPayload(payload string) ([]SensorValue, error) {
	parts := strings.Split(payload, ",")

	if len(parts)%3 != 0 {
		return nil, fmt.Errorf("client.go diz: payload inválido: número de elementos não é múltiplo de 3")
	}

	var results []SensorValue
	for i := 0; i < len(parts); i += 3 {
		val, err := strconv.ParseFloat(parts[i+1], 64)
		if err != nil {
			return nil, fmt.Errorf("client.go diz: erro ao converter valor '%s': %w", parts[i+1], err)
		}

		results = append(results, SensorValue{
			infoType:    parts[i],
			Valor:   val,
			Unidade: parts[i+2],
		})
	}
	return results, nil
}

// Exemplo de uso no messageHandler
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("[MQTT] Mensagem recebida em %s\n", msg.Topic())

	// Parsear a mensagem completa incluindo dados dos sensores
	ttnMsg, err := ParseCompleteTTNMessage(msg.Payload())
	if err != nil {
		fmt.Println("Erro ao processar mensagem:", err)
		return
	}

	// Agora ttnMsg já contém tudo em um único objeto
	fmt.Printf("Dispositivo: %s\n", ttnMsg.EndDeviceIDs.DeviceID)
	fmt.Printf("Recebido em: %s\n", ttnMsg.ReceivedAt)
	fmt.Printf("Dados dos sensores: %+v\n", ttnMsg.SensorData)

	// Converter para JSON completo
	jsonData, err := json.MarshalIndent(ttnMsg, "", "  ")
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}
	
	fmt.Println("Mensagem completa em JSON:")
	fmt.Println(string(jsonData))
	
	// Agora você pode enviar jsonData para o frontend ou MongoDB
}