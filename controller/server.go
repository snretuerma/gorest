package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initHandlers() {
	router.Get("/api/posts", controller.GetAllPosts)
}

func Start() {
	router := chi.NewRouter()

	initHandlers()
	fmt.Printf("router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
