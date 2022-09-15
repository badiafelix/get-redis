package main

import (
	"api-get-redis/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users/profile", controllers.C_getProfile).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
