package main

import (
	"log"
	"net/http"
	"uart-mp3-player/internal/handlers"

	rice "github.com/GeertJohan/go.rice"
)

func main() {
	initRoutes()

	log.Println("Server starting at port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func initRoutes() {
	appBox, err := rice.FindBox("./client-dist")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/ping", handlers.PingHandler)
	http.HandleFunc("/api/reset", handlers.ResetHandler)

	http.Handle("/static/", http.FileServer(appBox.HTTPBox()))
	http.HandleFunc("/manifest.json", handlers.ServeResourceHandler(appBox, "manifest.json"))
	http.HandleFunc("/favicon.ico", handlers.ServeResourceHandler(appBox, "favicon.ico"))
	http.HandleFunc("/logo.png", handlers.ServeResourceHandler(appBox, "logo.png"))
	http.HandleFunc("/", handlers.ServeResourceHandler(appBox, "index.html"))
}
