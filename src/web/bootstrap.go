package web

import (
	"net/http"
)

func StartupServer() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	http.ListenAndServe(":8080", mux)
}

func main() {
	StartupServer()
}
