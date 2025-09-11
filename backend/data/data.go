package data

import (
	"time"
	"sync"
)


//STRUCTS e VARIAVEIS GLOBAIS DE BANCO DE DADOS=========================================================================
type Topic struct {
    Nome string `bson:"nome" json:"nome"`
}

type Sensor struct {
    MqttID       string    `bson:"mqttID" json:"mqttID"`
    TopicName    string    `bson:"topicName" json:"topicName"`
    Descricao    string    `bson:"descricao" json:"descricao"`
    LastUpdate   time.Time `bson:"lastUpdate" json:"lastUpdate"`
	LastData	[]SensorValue `bson:"lastData" json:"lastData"` 
    ShowOnScreen bool      `bson:"showOnScreen" json:"showOnScreen"`
}

type SensorData struct {
    IDSensor  string          `bson:"id_sensor" json:"id_sensor"`
    Timestamp time.Time       `bson:"timestamp" json:"timestamp"`
    Dados     map[string]any  `bson:"dados" json:"dados"`
}

//=============================================================================================================
//STRUCTS E VARIAVEIS GLOBAIS DE MQTT ================================================================================================

type SensorValue struct {
	InfoType string  `json:"infoType"`
	Valor    float64 `json:"valor"`
	Unidade  string  `json:"unidade"`
}

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
	} `json:"uplink_message"`

	SensorData []SensorValue `json:"sensor_data,omitempty"`
}

//=============================================================================================================
//STRUCTS E VARIAVEIS GLOBAIS DE API

type InfoDisplay struct{
	SensorID string `json:"SensorID"`
	SensorData []SensorValue `json:"SensorData,omitempty"`
}

var DadosDisplay []InfoDisplay;
var DadosDisplayMutex sync.Mutex;

func removerDoDisplay(){}