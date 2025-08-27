package db

import (
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Topic struct {
	Nome string `bson:"nome"`
}

type Sensor struct {
    MqttID     string    `bson:"mqttID"`
    TopicName   string `bson:"topicName"`
    Descricao  string    `bson:"descricao"`
    LastUpdate time.Time `bson:"lastUpdate"`
	ShowOnScreen bool `bson:"showOnScreen"`
}

type SensorData struct {
	IDSensor string            `bson:"id_sensor"`
	Timestamp time.Time        `bson:"timestamp"`
	Dados     map[string]any   `bson:"dados"`
}

var client *mongo.Client
var TopicsCollection *mongo.Collection
var SensorsCollection *mongo.Collection
var DataCollection *mongo.Collection

