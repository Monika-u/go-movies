package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Monika-u/go-movies/config"
	"github.com/gorilla/mux"
)

func main() {
	// r := mux.NewRouter()
	config.InitializeDB()

	// r.HandleFunc("/movies", getMovies).Methods("GET")
	fmt.Println("starting server : 8000 port \n")
	log.Fatal(http.ListenAndServe(":8000", r))
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
