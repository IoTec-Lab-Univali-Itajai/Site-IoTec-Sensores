// api.go
package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"fmt"

	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/db" // IMPORTA O PACOTE DB
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StartAPI() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // cache preflight for 5 minutes
	}))

	// ROTAS
	r.Get("/api/topics", GetTopicsWithSensors)
	r.Get("/api/sensorView", GetSensorView)
	r.Delete("/api/sensor/{id}", DeleteSensor)
	r.Post("/api/sensor", CreateSensor)
	r.Post("/api/topic", CreateTopic)
	r.Put("/api/sensor/{id}", UpdateSensorVisibility)


	http.ListenAndServe(":8080", r)
}

// Handlers
func GetTopicsWithSensors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	cursor, err := db.TopicsCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Erro ao buscar tópicos", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var topicos []bson.M
	if err = cursor.All(ctx, &topicos); err != nil {
		http.Error(w, "Erro ao decodificar tópicos", http.StatusInternalServerError)
		return
	}

	var resposta []map[string]interface{}
	for _, topico := range topicos {
		nomeTopico, ok := topico["nome"].(string)
		if !ok {
			continue
		}

		filter := bson.M{"topicID": topico["_id"]}
		cursorSensores, err := db.SensorsCollection.Find(ctx, filter)

		if err != nil {
			continue
		}
		var sensores []bson.M
		cursorSensores.All(ctx, &sensores)

		for i := range sensores {
			if val, ok := sensores[i]["lastUpdate"]; ok {
				sensores[i]["ultimaAtualizacao"] = val
			}
		}

		resposta = append(resposta, map[string]interface{}{
			"nome":     nomeTopico,
			"sensores": sensores,
		})
	}

	json.NewEncoder(w).Encode(resposta)
}

func DeleteSensor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idSensor := chi.URLParam(r, "id")
	if idSensor == "" {
		http.Error(w, "ID do sensor não fornecido", http.StatusBadRequest)
		return
	}

	ctx := context.TODO()
	result, err := db.SensorsCollection.DeleteOne(ctx, bson.M{"mqttID": idSensor})
	if err != nil {
		http.Error(w, "Erro ao remover sensor", http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "Sensor não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Sensor removido com sucesso"})
}

// Adicione esta rota junto com as outras rotas

// Handler para criar um novo sensor
func CreateSensor(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var payload struct {
        MqttID    string `json:"mqttID"`
        TopicName string `json:"topicName"` // Agora recebemos o nome do tópico
        Descricao string `json:"descricao"`
		ShowOnScreen string `json:"showOnScreen"`
    }


    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "Dados do sensor inválidos", http.StatusBadRequest)
        return
    }

    ctx := context.TODO()

    // Busca o tópico pelo nome para obter seu ObjectID
    var topico bson.M
    err := db.TopicsCollection.FindOne(ctx, bson.M{"nome": payload.TopicName}).Decode(&topico)
    if err != nil {
        http.Error(w, "Tópico não encontrado", http.StatusNotFound)
        return
    }

    topicID := topico["_id"].(primitive.ObjectID)

	// Converte ShowOnScreen de string para bool
    showOnScreenBool, err := strconv.ParseBool(payload.ShowOnScreen)
    if err != nil {
        fmt.Println("Erro na conversão:", err)
        showOnScreenBool = false // valor padrão em caso de erro
    }

    // Cria o sensor com todos os campos obrigatórios
    sensor := bson.M{
        "mqttID":     payload.MqttID,
        "topicID":    topicID,
        "descricao": payload.Descricao,
        "lastUpdate": time.Date(1977, 11, 18, 12, 0, 0, 0, time.UTC),
		"showOnScreen":  showOnScreenBool,
    }
	
    _, err = db.SensorsCollection.InsertOne(ctx, sensor)
    if err != nil {
        http.Error(w, "Erro ao criar sensor: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Sensor criado com sucesso"})
}

// UpdateSensorVisibility atualiza o campo showOnScreen de um sensor
func UpdateSensorVisibility(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    mqttID := chi.URLParam(r, "id")
    if mqttID == "" {
        http.Error(w, "mqttID do sensor não fornecido", http.StatusBadRequest)
        return
    }

    var payload struct {
        ShowOnScreen bool `json:"showOnScreen"`
    }
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "Payload inválido", http.StatusBadRequest)
        return
    }

    ctx := context.TODO()
    filter := bson.M{"mqttID": mqttID}
    update := bson.M{"$set": bson.M{"showOnScreen": payload.ShowOnScreen}}

    result, err := db.SensorsCollection.UpdateOne(ctx, filter, update)
    if err != nil {
        http.Error(w, "Erro ao atualizar sensor: "+err.Error(), http.StatusInternalServerError)
        return
    }

    if result.MatchedCount == 0 {
        http.Error(w, "Sensor não encontrado", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "message": "Sensor atualizado com sucesso",
    })
}

func CreateTopic(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var payload struct {
        Nome string `json:"nome"`
    }

    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "Payload inválido", http.StatusBadRequest)
        return
    }

    if payload.Nome == "" {
        http.Error(w, "Nome do tópico é obrigatório", http.StatusBadRequest)
        return
    }

    ctx := context.TODO()

    // Verifica se já existe um tópico com esse nome
    var existente bson.M
    err := db.TopicsCollection.FindOne(ctx, bson.M{"nome": payload.Nome}).Decode(&existente)
    if err == nil {
        http.Error(w, "Já existe um tópico com esse nome", http.StatusConflict)
        return
    }

    // Cria o tópico
    novoTopico := db.Topic{
        Nome: payload.Nome,
    }

    _, err = db.TopicsCollection.InsertOne(ctx, novoTopico)
    if err != nil {
        http.Error(w, "Erro ao criar tópico: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Tópico criado com sucesso",
    })
}

func GetSensorView(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    sensores, err := db.BuscarSensoresDisplay()
    if err != nil {
        http.Error(w, "Erro ao buscar sensores para exibição: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Prepara resposta convertendo para JSON
    var resposta []map[string]interface{}
    for _, s := range sensores {
        resposta = append(resposta, map[string]interface{}{
            "mqttID":        s.MqttID,
            "descricao":     s.Descricao,
            "lastUpdate":    s.LastUpdate,
            "showOnScreen":  s.ShowOnScreen,
            "topicID":       s.TopicID, // mantém o vínculo com o tópico
        })
    }

    json.NewEncoder(w).Encode(resposta)
}