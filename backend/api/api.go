// api.go
package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores/backend/db" // IMPORTA O PACOTE DB
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
	r.Get("/api/topicsJSON", GetTopicsJSON)
	r.Get("/api/sensorJSON", GetSensorsJSON)
	r.Get("/api/sensorView", GetSensorView)
	r.Delete("/api/sensorDelete/{id}", DeleteSensor)
	r.Post("/api/sensorCreate", CreateSensor)
	r.Post("/api/topicCreate", CreateTopic)
	r.Put("/api/sensorUpdate/{id}", UpdateSensorVisibility)

	http.ListenAndServe(":8080", r)
}

// Handlers
func GetTopicsJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	topicos, err := db.BuscarTopicos()
	if err != nil {
		log.Printf("api.go diz: Erro ao buscar tópicos: %v", err)
		http.Error(w, `{"error": "Erro ao buscar tópicos"}`, http.StatusInternalServerError)
		return
	}

	log.Printf("api.go diz: %d tópicos encontrados", len(topicos))

	if err := json.NewEncoder(w).Encode(topicos); err != nil {
		log.Printf("api.go diz: Erro ao codificar JSON: %v", err)
		http.Error(w, `{"error": "Erro ao gerar resposta"}`, http.StatusInternalServerError)
		return
	}
}

func GetSensorsJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Obter o parâmetro 'topic' da query string
	topicName := r.URL.Query().Get("topic")
	if topicName == "" {
		log.Printf("api.go diz: Parâmetro 'topic' não fornecido")
		http.Error(w, `{"error": "Parâmetro 'topic' é obrigatório"}`, http.StatusBadRequest)
		return
	}

	// Buscar sensores pelo tópico
	sensores, err := db.BuscarSensoresPorTopico(topicName)
	if err != nil {
		log.Printf("api.go diz: Erro ao buscar sensores para o tópico '%s': %v", topicName, err)
		http.Error(w, `{"error": "Erro ao buscar sensores"}`, http.StatusInternalServerError)
		return
	}

	log.Printf("api.go diz: %d sensores encontrados para o tópico '%s'", len(sensores), topicName)

	// Retornar os sensores em formato JSON
	if err := json.NewEncoder(w).Encode(sensores); err != nil {
		log.Printf("api.go diz: Erro ao codificar JSON para sensores: %v", err)
		http.Error(w, `{"error": "Erro ao gerar resposta"}`, http.StatusInternalServerError)
		return
	}
}

func DeleteSensor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idSensor := chi.URLParam(r, "id")
	if idSensor == "" {
		http.Error(w, `{"error": "ID do sensor não fornecido"}`, http.StatusBadRequest)
		return
	}

	// Usando a nova função DeletarSensor
	err := db.DeletarSensor(idSensor)
	if err != nil {
		if err.Error() == "api.go diz: sensor não encontrado" {
			http.Error(w, `{"error": "Sensor não encontrado"}`, http.StatusNotFound)
		} else {
			log.Printf("api.go diz: Erro ao deletar sensor %s: %v", idSensor, err)
			http.Error(w, `{"error": "Erro ao remover sensor"}`, http.StatusInternalServerError)
		}
		return
	}

	// Retorna resposta de sucesso
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sensor removido com sucesso"})
}

func CreateSensor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var payload struct {
		MqttID       string `json:"mqttID"`
		TopicName    string `json:"topicName"`
		Descricao    string `json:"descricao"`
		ShowOnScreen bool `json:"showOnScreen"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("api.go diz: Dados do sensor inválidos - %v", err)
		http.Error(w, `{"error": "Dados do sensor inválidos"}`, http.StatusBadRequest)
		return
	}

	if !db.TopicoExiste(payload.TopicName) {
		log.Printf("api.go diz: topico inexistente")
		http.Error(w, `{"error": topico inexistente"}`, http.StatusBadRequest)
		return
	}

	// Cria o sensor com a struct Sensor
	sensor := db.Sensor{
		MqttID:       payload.MqttID,
		TopicName:    payload.TopicName, // Usando TopicName em vez de topicID
		Descricao:    payload.Descricao,
		LastUpdate:   time.Date(1977, 11, 18, 12, 0, 0, 0, time.UTC),
		ShowOnScreen: payload.ShowOnScreen,
	}

	err := db.InserirSensor(sensor)
	if err != nil {
		log.Printf("api.go diz: Erro ao criar sensor '%s' - %v", payload.MqttID, err)
		http.Error(w, `{"error": "Erro ao criar sensor"}`, http.StatusInternalServerError)
		return
	}

	log.Printf("api.go diz: Sensor '%s' criado com sucesso para o tópico '%s'", payload.MqttID, payload.TopicName)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sensor criado com sucesso"})
}

func UpdateSensorVisibility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mqttID := chi.URLParam(r, "id")
	if mqttID == "" {
		log.Printf("api.go diz: mqttID do sensor não fornecido")
		http.Error(w, `{"error": "mqttID do sensor não fornecido"}`, http.StatusBadRequest)
		return
	}

	var payload struct {
		ShowOnScreen bool `json:"showOnScreen"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("api.go diz: Payload inválido - %v", err)
		http.Error(w, `{"error": "Payload inválido"}`, http.StatusBadRequest)
		return
	}

	err := db.AtualizarSensorVisibilidade(mqttID, payload.ShowOnScreen)
	if err != nil {
		if strings.Contains(err.Error(), "não encontrado") {
			log.Printf("api.go diz: Sensor '%s' não encontrado", mqttID)
			http.Error(w, `{"error": "Sensor não encontrado"}`, http.StatusNotFound)
		} else {
			log.Printf("api.go diz: Erro ao atualizar sensor '%s' - %v", mqttID, err)
			http.Error(w, `{"error": "Erro ao atualizar sensor"}`, http.StatusInternalServerError)
		}
		return
	}

	log.Printf("api.go diz: Sensor '%s' atualizado com sucesso - showOnScreen: %t", mqttID, payload.ShowOnScreen)
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
		log.Printf("api.go diz: Payload inválido - %v", err)
		http.Error(w, `{"error": "Payload inválido"}`, http.StatusBadRequest)
		return
	}

	if payload.Nome == "" {
		log.Printf("api.go diz: Nome do tópico é obrigatório")
		http.Error(w, `{"error": "Nome do tópico é obrigatório"}`, http.StatusBadRequest)
		return
	}

	if db.TopicoExiste(payload.Nome) {
		log.Printf("api.go diz: Já existe um tópico com o nome '%s'", payload.Nome)
		http.Error(w, `{"error": "Já existe um tópico com esse nome"}`, http.StatusConflict)
		return
	}

	// Cria o tópico
	novoTopico := db.Topic{
		Nome: payload.Nome,
	}

	err := db.InserirTopico(novoTopico)
	if err != nil {
		log.Printf("api.go diz: Erro ao criar tópico '%s' - %v", payload.Nome, err)
		http.Error(w, `{"error": "Erro ao criar tópico"}`, http.StatusInternalServerError)
		return
	}

	log.Printf("api.go diz: Tópico '%s' criado com sucesso", payload.Nome)
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
			"mqttID":       s.MqttID,
			"descricao":    s.Descricao,
			"lastUpdate":   s.LastUpdate,
			"showOnScreen": s.ShowOnScreen,
			"topicID":      s.TopicName, // mantém o vínculo com o tópico
		})
	}

	json.NewEncoder(w).Encode(resposta)
}
