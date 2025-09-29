package data

import (
	"time"
	"sync"
)


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

type SensorValue struct {
	InfoType string  `json:"infoType" bson:"infotype"`
	Valor    float64 `json:"valor" bson:"valor"`
	Unidade  string  `json:"unidade" bson:"unidade"`
}

type InfoDisplay struct{
	SensorID string `json:"SensorID"`
	SensorData []SensorValue `json:"SensorData,omitempty"`
}

var DadosDisplay []InfoDisplay;
var DadosDisplayMutex sync.Mutex;