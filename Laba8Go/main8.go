package main

import (
	"Laba8Go/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", routes.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", routes.GetUser).Methods("GET")
	router.HandleFunc("/users", routes.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", routes.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", routes.DeleteUser).Methods("DELETE")

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8081", router))
}
