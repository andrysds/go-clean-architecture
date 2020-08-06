package main

import (
	"log"
	"net/http"

	"github.com/andrysds/go-clean-architecture/handler"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	healthzHandler := &handler.HealthzHandler{}
	router.GET("/healthz", healthzHandler.Index)

	port := ":8080"
	log.Println("Listening at port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
