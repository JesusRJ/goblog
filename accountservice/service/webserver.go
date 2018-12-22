package service

import (
	"log"
	"net/http"
)

// StartWebServer Inicia o servidor web para o Account Service
func StartWebServer(port string) {
	log.Println("Starting HTTP service at " + port)

	// Adiciona as rotas
	r := NewRouter()
	http.Handle("/", r)

	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		log.Println("An error ocurred starting HTTP listener at port " + port)
		log.Printf("Error: " + err.Error())
	}
}
