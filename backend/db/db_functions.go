package db

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
)

type Topic struct {
	Nome string `bson:"nome"`
}

type Sensor struct {
	ID               string `bson:"id"`
	IDTopic          string `bson:"id_topic"`
	Descricao        string `bson:"descricao"`
	UltimaAtualizacao string `bson:"ultimaAtualizacao"`
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
