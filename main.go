package main

import (
	"log"
	"net/http"

	"github.com/andrysds/go-clean-architecture/connection"
	"github.com/andrysds/go-clean-architecture/handler"
	"github.com/andrysds/go-clean-architecture/repository"
	"github.com/andrysds/go-clean-architecture/service"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// connections
	db := connection.NewDB()

	// repositories
	friendRepository := &repository.FriendRepository{DB: db}

	// services
	friendService := &service.FriendService{FriendRepository: friendRepository}

	// router & handlers
	router := httprouter.New()

	healthzHandler := &handler.HealthzHandler{}
	router.GET("/healthz", healthzHandler.Index)

	friendHandler := &handler.FriendHandler{FriendService: friendService}
	router.GET("/friends", friendHandler.Index)
	router.POST("/friends", friendHandler.Create)
	router.PUT("/friends/:id", friendHandler.Update)
	router.DELETE("/friends/:id", friendHandler.Delete)

	port := ":8080"
	log.Println("Listening at port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
