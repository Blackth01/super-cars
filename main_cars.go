package main

import (
	"supercars/cars"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/cars", cars.CreateCar).Methods(http.MethodPost)
	router.HandleFunc("/cars", cars.GetCars).Methods(http.MethodGet)
	router.HandleFunc("/cars/{id}", cars.GetCar).Methods(http.MethodGet)
	router.HandleFunc("/cars/{id}", cars.UpdateCar).Methods(http.MethodPut)
	router.HandleFunc("/cars/{id}", cars.DeleteCar).Methods(http.MethodDelete)

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
