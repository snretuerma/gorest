package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	blog "github.com/snretuerma/gorest/controller/blog"
)

func Start() {
	router := chi.NewRouter()

	router.Get("/api/posts", blog.GetAllPosts)
	router.Get("/api/post/{id}", blog.GetPost)
	router.Post("/api/post", blog.CreatePost)
	router.Put("/api/post/{id}", blog.UpdatePost)
	router.Delete("/api/post/{id}", blog.DeletePost)
	fmt.Printf("Router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
