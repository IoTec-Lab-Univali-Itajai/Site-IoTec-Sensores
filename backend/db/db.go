package db

import (
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Topic struct {
    Nome string `bson:"nome" json:"nome"`
}

type Sensor struct {
    MqttID       string    `bson:"mqttID" json:"mqttID"`
    TopicName    string    `bson:"topicName" json:"topicName"`
    Descricao    string    `bson:"descricao" json:"descricao"`
    LastUpdate   time.Time `bson:"lastUpdate" json:"lastUpdate"`
    ShowOnScreen bool      `bson:"showOnScreen" json:"showOnScreen"`
}

type SensorData struct {
    IDSensor  string          `bson:"id_sensor" json:"id_sensor"`
    Timestamp time.Time       `bson:"timestamp" json:"timestamp"`
    Dados     map[string]any  `bson:"dados" json:"dados"`
}

var client *mongo.Client
var TopicsCollection *mongo.Collection
var SensorsCollection *mongo.Collection
var DataCollection *mongo.Collection

