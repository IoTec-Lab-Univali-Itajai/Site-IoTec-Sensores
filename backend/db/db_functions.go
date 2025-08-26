package db

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Topic struct {
	Nome string `bson:"nome"`
}

type Sensor struct {
    MqttID     string    `bson:"mqttID"`
    TopicID    primitive.ObjectID `bson:"topicID"`
    Descricao  string    `bson:"descricao"`
    LastUpdate time.Time `bson:"lastUpdate"`
	ShowOnScreen bool `bson:"showOnScreen"`
}

type SensorData struct {
	IDSensor string            `bson:"id_sensor"`
	Timestamp time.Time        `bson:"timestamp"`
	Dados     map[string]any   `bson:"dados"`
}

func InserirTopico(topico Topic) error {
	_, err := TopicsCollection.InsertOne(context.TODO(), topico)
	return err
}

func InserirSensor(sensor Sensor) error {
	_, err := SensorsCollection.InsertOne(context.TODO(), sensor)
	return err
}

func InserirDado(dado SensorData) error {
	_, err := DataCollection.InsertOne(context.TODO(), dado)
	return err
}

func BuscarTopicos() ([]Topic, error) {
	cursor, err := TopicsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var topicos []Topic
	if err = cursor.All(context.TODO(), &topicos); err != nil {
		return nil, err
	}
	return topicos, nil
}

func BuscarSensoresPorTopico(idTopic string) ([]Sensor, error) {
	filter := bson.M{"id_topic": idTopic}
	cursor, err := SensorsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var sensores []Sensor
	if err = cursor.All(context.TODO(), &sensores); err != nil {
		return nil, err
	}
	return sensores, nil
}

func BuscarSensoresDisplay() ([]Sensor, error) {
    filter := bson.M{"showOnScreen": true}
    cursor, err := SensorsCollection.Find(context.TODO(), filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    var sensores []Sensor
    if err = cursor.All(context.TODO(), &sensores); err != nil {
        return nil, err
    }
    return sensores, nil
}
