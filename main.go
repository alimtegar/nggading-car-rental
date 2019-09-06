package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/alimtegar/nggading-car-rental-system/models"
	"github.com/alimtegar/nggading-car-rental-system/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/validator.v2"
)

var client *mongo.Client

func main() {
	log.Println("Starting the application...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()

	// Users Routes
	router.HandleFunc("/users", handlers.getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.getUser).Methods("GET")
	router.HandleFunc("/users", handlers.addUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.deleteUser).Methods("DELETE")

	// Cars Routes
	router.HandleFunc("/cars", handlers.getCars).Methods("GET")
	router.HandleFunc("/cars/{id}", handlers.getCar).Methods("GET")
	router.HandleFunc("/cars", handlers.addCar).Methods("POST")
	router.HandleFunc("/cars/{id}", handlers.updateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", handlers.deleteCar).Methods("DELETE")

	// Orders Routes
	router.HandleFunc("/orders", handlers.getOrders).Methods("GET")
	router.HandleFunc("/orders", handlers.addOrder).Methods("POST")

	http.ListenAndServe(":3001", router)
}
