package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func indexGorilla(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing indexGorilla handler")
	fmt.Fprintf(w, "Welcome indexGorilla!")
}
func aboutGorilla(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing aboutGorilla handler")
	fmt.Fprintf(w, "Go Middleware indexGorilla")
}
func iconHandlerGorilla(w http.ResponseWriter, r *http.Request) {
}

func StartGorillaHandlersServer() {
	http.HandleFunc("/favicon.ico", iconHandler)
	indexHandler := http.HandlerFunc(indexGorilla)
	aboutHandler := http.HandlerFunc(aboutGorilla)
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	http.Handle("/", handlers.LoggingHandler(logFile, handlers.CompressHandler(indexHandler)))
	http.Handle("/about", handlers.LoggingHandler(logFile, handlers.CompressHandler(
		aboutHandler)))
	server := &http.Server{
		Addr: ":8080",
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
