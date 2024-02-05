package main

import (
	"github.com/gorilla/mux"
	"github.com/scnguenha/learn-go/book-store/pkg/routes"
	"log"
	"net/http"
)

func main() {
	log.Println("Server starting on port 9010")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))
}
