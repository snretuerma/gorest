package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	blog "github.com/snretuerma/gorest/controller/blog"
)

// func initHandlers() {
// 	router.Get("/api/posts", controller.GetAllPosts)
// }

func Start() {
	router := chi.NewRouter()

	router.Get("/api/posts", blog.GetAllPosts)
	fmt.Printf("Router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
