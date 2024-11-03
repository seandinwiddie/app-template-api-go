package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// AppState holds the application state with thread-safe access
type AppState struct {
	InitialState map[string]interface{}
	mutex        sync.RWMutex
}

func home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome to the API",
	})
}

func status(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status": "OK",
	})
}

func getData(state *AppState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state.mutex.RLock()
		defer state.mutex.RUnlock()
		json.NewEncoder(w).Encode(state.InitialState)
	}
}

func getByKey(state *AppState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]

		state.mutex.RLock()
		value, exists := state.InitialState[key]
		state.mutex.RUnlock()

		if !exists {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Key not found",
			})
			return
		}

		response := map[string]interface{}{
			key: value,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	// Read initial state
	data, err := ioutil.ReadFile("data/initialState.json")
	if err != nil {
		log.Fatal("Failed to read initialState.json:", err)
	}

	var initialState map[string]interface{}
	if err := json.Unmarshal(data, &initialState); err != nil {
		log.Fatal("Failed to parse JSON:", err)
	}

	// Create app state
	state := &AppState{
		InitialState: initialState,
	}

	// Create router
	r := mux.NewRouter()

	// Setup routes
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/status", status).Methods("GET")
	r.HandleFunc("/data", getData(state)).Methods("GET")
	r.HandleFunc("/{key}", getByKey(state)).Methods("GET")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Configure appropriately for production
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	// Create server
	handler := c.Handler(r)
	server := &http.Server{
		Addr:    ":3000",
		Handler: handler,
	}

	log.Println("Server starting on http://localhost:3000")
	log.Fatal(server.ListenAndServe())
}
