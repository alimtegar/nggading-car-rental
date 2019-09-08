package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/alimtegar/nggading-car-rental-system/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/validator.v2"
)

// GetOrders Handler
func GetOrders(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// var orders []Order
	var orders []map[string]interface{}

	collection := client.Database("nggadingCarRentalSystem").Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	pipeline := mongo.Pipeline{
		{{"$lookup", bson.D{
			{"from", "customers"},
			{"localField", "customer_id"},
			{"foreignField", "_id"},
			{"as", "customer"},
		}}},
		{{"$lookup", bson.D{
			{"from", "cars"},
			{"localField", "car_id"},
			{"foreignField", "_id"},
			{"as", "car"},
		}}},
		{{"$unwind", "$customer"}},
		{{"$unwind", "$car"}},
		// {{"$unset", "customer_id"}},
		// {{"$unset", "car_id"}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		// var order Order
		var order map[string]interface{}

		cursor.Decode(&order)

		orders = append(orders, order)
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(orders)
}

// AddOrder Handler
func AddOrder(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestedOrder models.RequestedOrder
	// var order models.Order

	_ = json.NewDecoder(r.Body).Decode(&requestedOrder)

	// order.CreatedAt = time.Now()

	// Validation
	if err := validator.Validate(requestedOrder); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	customer := models.Customer{
		Name:       requestedOrder.Name,
		Email:      requestedOrder.Email,
		Phone:      requestedOrder.Phone,
		NIK:        requestedOrder.NIK,
		STNKNumber: requestedOrder.STNKNumber,
		Address:    requestedOrder.Address,
		CreatedAt:  time.Now(),
	}

	collection := client.Database("nggadingCarRentalSystem").Collection("customers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	result, err := collection.InsertOne(ctx, customer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	// I know this code below is ugly -_-)
	insertedCustomerID := strings.TrimLeft(strings.TrimRight(fmt.Sprint(result.InsertedID), `")`), `ObjectID("`)

	customerID, err := primitive.ObjectIDFromHex(insertedCustomerID)

	if err := validator.Validate(requestedOrder); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	carID, err := primitive.ObjectIDFromHex(requestedOrder.CarID)

	if err := validator.Validate(requestedOrder); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	order := models.Order{
		CustomerID: customerID,
		CarID:      carID,
		CreatedAt:  time.Now(),
	}

	collection = client.Database("nggadingCarRentalSystem").Collection("orders")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	result, err = collection.InsertOne(ctx, order)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(result)
}

// DeleteOrder Handler
func DeleteOrder(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := client.Database("nggadingCarRentalSystem").Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	result, err := collection.DeleteOne(ctx, models.Order{ID: id})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(result)
}
