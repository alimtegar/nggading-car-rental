package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/alimtegar/nggading-car-rental-system/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/validator.v2"
)

func GetCars(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cars []models.Car

	collection := client.Database("nggadingCarRentalSystem").Collection("cars")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var car models.Car

		cursor.Decode(&car)

		cars = append(cars, car)
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(cars)
}

func GetCar(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var car models.Car

	collection := client.Database("nggadingCarRentalSystem").Collection("cars")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	err := collection.FindOne(ctx, models.Car{ID: id}).Decode(&car)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(car)
}

func AddCar(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var car models.Car

	_ = json.NewDecoder(r.Body).Decode(&car)
	car.CreatedAt = time.Now()

	// Validation
	if err := validator.Validate(car); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	collection := client.Database("nggadingCarRentalSystem").Collection("cars")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	result, err := collection.InsertOne(ctx, car)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateCar(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var car models.Car

	_ = json.NewDecoder(r.Body).Decode(&car)
	car.UpdatedAt = time.Now()

	collection := client.Database("nggadingCarRentalSystem").Collection("cars")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	result, err := collection.UpdateOne(ctx, models.Car{ID: id}, bson.M{"$set": car})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(result)
}

func DeleteCar(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := client.Database("nggadingCarRentalSystem").Collection("cars")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	result, err := collection.DeleteOne(ctx, models.Car{ID: id})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(result)
}
