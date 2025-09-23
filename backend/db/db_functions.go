package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

var Client *mongo.Client
var TopicsCollection *mongo.Collection
var SensorsCollection *mongo.Collection
var DataCollection *mongo.Collection

func ConnectMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_HOST")
	if uri == "" {
		log.Fatal("db.functions diz: MONGO_HOST não encontrado no .env")
	}

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	db := Client.Database("iotec-lab-database")
	TopicsCollection = db.Collection("topicsCollection")
	SensorsCollection = db.Collection("sensorsCollection")
	DataCollection = db.Collection("dataCollection")

	log.Println("db.functions diz: ✅ Conectado ao MongoDB com sucesso.")
}

func InserirTopico(topico data.Topic) error {
	_, err := TopicsCollection.InsertOne(context.TODO(), topico)
	return err
}

func InserirSensor(sensor data.Sensor) error {
	_, err := SensorsCollection.InsertOne(context.TODO(), sensor)
	return err
}

func InserirDado(dado data.InfoDisplay) error {
	// Converte os dados para o formato exigido pelo schema
	dadosMap := make(map[string]interface{})
	for _, sensorValue := range dado.SensorData {
		dadosMap[sensorValue.InfoType] = map[string]interface{}{
			"valor":   sensorValue.Valor,
			"unidade": sensorValue.Unidade,
		}
	}

	// Cria o documento no formato correto
	document := bson.M{
		"sensorId":   dado.SensorID,
		"timestamp":  time.Now(), // Usa o timestamp atual
		"dados":      dadosMap,   // Mapa com os dados
	}

	_, err := DataCollection.InsertOne(context.TODO(), document)
	return err
}

func BuscarTopicos() ([]data.Topic, error) {
	cursor, err := TopicsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var topicos []data.Topic
	if err = cursor.All(context.TODO(), &topicos); err != nil {
		return nil, err
	}
	return topicos, nil
}

func BuscarSensoresPorTopico(topicName string) ([]data.Sensor, error) {
	filter := bson.M{"topicName": topicName}
	fmt.Println("db_funcitons diz: topico selecionado foi: ", topicName);
	cursor, err := SensorsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var sensores []data.Sensor
	if err = cursor.All(context.TODO(), &sensores); err != nil {
		return nil, err
	}
	return sensores, nil
}

func BuscarSensorPorNome(mqttID string)(data.Sensor, error){
	filter := bson.M{"mqttID": mqttID}
	var sensor data.Sensor
	err := SensorsCollection.FindOne(context.TODO(), filter).Decode(&sensor)
	if err != nil {
		log.Printf("api.go diz: Sensor '%s' não encontrado no BD", mqttID)
		return data.Sensor{}, err
	}
	return sensor, nil;
}

func BuscarSensoresDisplay() ([]data.Sensor, error) {
	filter := bson.M{"showOnScreen": true}
	cursor, err := SensorsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var sensores []data.Sensor
	if err = cursor.All(context.TODO(), &sensores); err != nil {
		return nil, err
	}
	return sensores, nil
}

func SensorExiste(mqttID string) bool {
	filter := bson.M{"mqttID": mqttID}
	
	var result struct { ID primitive.ObjectID `bson:"_id"` }
	err := SensorsCollection.FindOne(context.TODO(), filter).Decode(&result)
	
	return err == nil
}

func DeletarSensor(mqttID string) error {
	filter := bson.M{"mqttID": mqttID}

	result, err := SensorsCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("db_functions diz: nenhum sensor com mqttID '%s' encontrado", mqttID)
	}

	return nil
}

func TopicoExiste(nome string) bool {
	err := TopicsCollection.FindOne(context.TODO(), bson.M{"nome": nome}).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
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

func AtualizarDados(SensorID string, SensorData []data.SensorValue) error {
	filter := bson.M{"mqttID": SensorID}
	
	update := bson.M{
		"$set": bson.M{
			"lastData":     SensorData,    // Campo correto: lastData
			"lastUpdate":   time.Now(),    // Campo correto: lastUpdate
		},
	}
	
	result, err := SensorsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("db_functions diz: erro ao atualizar sensor %s: %w", SensorID, err)
	}
	
	if result.MatchedCount == 0 {
		return fmt.Errorf("db_functions diz: sensor não encontrado: %s", SensorID)
	}
	
	fmt.Printf("db_functions diz: Dados do sensor %s atualizados com sucesso. %d documento(s) modificado(s)\n", SensorID, result.ModifiedCount)
	return nil
}