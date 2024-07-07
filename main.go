package main

import (
	"immmodi/framework/handlers"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server http://127.0.0.1:8000")

	http.HandleFunc("/", handlers.MakeHandler(handlers.RootHandler))
	http.ListenAndServe(":8000", nil)
}
