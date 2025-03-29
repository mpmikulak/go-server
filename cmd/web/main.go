package main

import (
	"log"
	"net/http"
)

const serverPort = ":4000"

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on %s", serverPort)
	err := http.ListenAndServe(serverPort, mux)
	log.Fatal(err)
}
