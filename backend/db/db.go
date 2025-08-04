package db

import (
	"context"
	"log"
	"os"
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var TopicsCollection *mongo.Collection
var SensorsCollection *mongo.Collection
var DataCollection *mongo.Collection

func ConnectMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_HOST")
	if uri == "" {
		log.Fatal("MONGO_HOST não encontrado no .env")
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

	log.Println("✅ Conectado ao MongoDB com sucesso.")
}
