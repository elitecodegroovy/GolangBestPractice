package strcon

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development")
}

/**
* A simple way startup Server.
**/
func StartSimpleServer() {
	http.HandleFunc("/welcome", messageHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

/**
	type Server struct {
	Addr string
	Handler Handler
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	MaxHeaderBytes int
	TLSConfig *tls.Config
	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
	ConnState func(net.Conn, ConnState)
	ErrorLog *log.Logger
	}
**/
func StartServer() {
	http.HandleFunc("/welcome", messageHandler)
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
