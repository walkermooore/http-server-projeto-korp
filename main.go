package main

import (
	"http-server-projeto-korp/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc(
		"/projeto-korp",
		handlers.ProjetoKorp,
	)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
