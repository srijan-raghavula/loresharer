package main

import (
	"github.com/srijan-raghavula/loresharer/internal/handlers"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("."))
	server := &http.Server{
		Handler: mux,
		Addr:    port,
	}

	mux.Handle("/static/*", http.StripPrefix("/static", fs))
	http.HandleFunc("GET /", handlers.RootHandler)

	log.Println("Starting server")
	log.Fatal(server.ListenAndServe())
}
