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

func GetUsers(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []models.ProtectedUser

	collection := client.Database("nggadingCarRentalSystem").Collection("users")

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
		var user models.ProtectedUser

		cursor.Decode(&user)

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(users)
}

func GetUser(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var user models.ProtectedUser

	collection := client.Database("nggadingCarRentalSystem").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	// err := collection.FindOne(ctx, models.User{ID: id}).Decode(&user)
	err := collection.FindOne(ctx, models.User{ID: id}).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(user)
}

func AddUser(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	user.CreatedAt = time.Now()

	// Validation
	if err := validator.Validate(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	collection := client.Database("nggadingCarRentalSystem").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateUser(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	user.UpdatedAt = time.Now()

	collection := client.Database("nggadingCarRentalSystem").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	result, err := collection.UpdateOne(ctx, models.User{ID: id}, bson.M{"$set": user})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(result)
}

func DeleteUser(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := client.Database("nggadingCarRentalSystem").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	result, err := collection.DeleteOne(ctx, models.User{ID: id})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(result)
}
