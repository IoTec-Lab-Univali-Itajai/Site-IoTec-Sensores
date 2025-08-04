// api.go
package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/db" // IMPORTA O PACOTE DB
)

func StartAPI() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// ROTAS
	r.Get("/api/topics", GetTopicsWithSensors)
	r.Delete("/api/sensor/{id}", DeleteSensor)

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

		cursorSensores, err := db.SensorsCollection.Find(ctx, bson.M{"id_topic": nomeTopico})
		if err != nil {
			continue
		}
		var sensores []bson.M
		cursorSensores.All(ctx, &sensores)

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
	result, err := db.SensorsCollection.DeleteOne(ctx, bson.M{"id": idSensor})
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
