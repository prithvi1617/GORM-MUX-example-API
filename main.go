package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func main() {
	fmt.Println("GORM, MUX example API")

	initialMigration()
	// Handle Subsequent requests
	handleRequests()
}
