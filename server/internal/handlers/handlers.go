package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"uart-mp3-player/internal/repository/sd"
	"uart-mp3-player/internal/uart"

	rice "github.com/GeertJohan/go.rice"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Pong"))
}

func SDContentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := sd.GetAll()
	if err != nil {
		log.Fatalf("Error retrieve SD content: %v\n", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	uart.Reset()
}

func ServeResourceHandler(app *rice.Box, resourceName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexFile, err := app.Open(resourceName)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		http.ServeContent(w, r, resourceName, time.Time{}, indexFile)
	}
}
