package main

import (
	"github.com/srijan-raghavula/loresharer/internal/handlers"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))

	mux.Handle(
		"/static/",
		http.StripPrefix("/static", fs),
	)
	mux.HandleFunc(
		"GET /home/", handlers.RootHandler,
	)

	server := &http.Server{
		Handler: mux,
		Addr:    port,
	}
	log.Println("Starting server")
	log.Fatal(server.ListenAndServe())
}
