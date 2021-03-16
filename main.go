package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/go-chi/chi"
)

var db sync.Map

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("port is not set")
	}

	router := chi.NewRouter()
	initRoutes(router)

	log.Print("Starting server...")
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatal(err)
	}
}

func initRoutes(router *chi.Mux) {
	router.Get("/list", List)
	router.Get("/get/{key}", Get)
	router.Post("/upsert", Upsert)
	router.Delete("/delete/{key}", DeleteByKey)
}
