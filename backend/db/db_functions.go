package db

import (
	"context"
	"log"
	"os"
	"time"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

func ConnectMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_HOST")
	if uri == "" {
		log.Fatal("db.functions diz: MONGO_HOST não encontrado no .env")
	}

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("iotec-lab-database")
	TopicsCollection = db.Collection("topicsCollection")
	SensorsCollection = db.Collection("sensorsCollection")
	DataCollection = db.Collection("dataCollection")

	log.Println("db.functions diz: ✅ Conectado ao MongoDB com sucesso.")
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

func BuscarSensoresPorTopico(topicName string) ([]Sensor, error) {
    filter := bson.M{"topicName": topicName}
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

func DeletarSensor(mqttID string) error {
    filter := bson.M{"mqttID": mqttID}
    
    result, err := SensorsCollection.DeleteOne(context.TODO(), filter)
    if err != nil {
        return err
    }
    
    if result.DeletedCount == 0 {
        return fmt.Errorf("")
    }
    
    return nil
}

func TopicoExiste(nome string) bool {
    err := TopicsCollection.FindOne(context.TODO(), bson.M{"nome": nome}).Err()
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return false
        }
        // Em caso de outro erro, logamos mas consideramos que não existe
        log.Printf("db_functions diz: Erro ao verificar tópico '%s': %v", nome, err)
        return false
    }
    return true
}

func AtualizarSensorVisibilidade(mqttID string, showOnScreen bool) error {
    filter := bson.M{"mqttID": mqttID}
    update := bson.M{"$set": bson.M{"showOnScreen": showOnScreen}}

    result, err := SensorsCollection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return err
    }

    if result.MatchedCount == 0 {
        return fmt.Errorf("db_functions diz: sensor com mqttID '%s' não encontrado", mqttID)
    }

    return nil
}
