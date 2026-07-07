package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("go-chi up on :" + port + "\n"))
	})
	println("Listening on :" + port)
	http.ListenAndServe(":"+port, r)
}
