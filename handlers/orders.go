package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/alimtegar/nggading-car-rental-system/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/validator.v2"
)

func GetOrders(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// var orders []Order
	var orders []map[string]interface{}

	collection := client.Database("nggadingCarRentalSystem").Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	// cursor, err := collection.Find(ctx, bson.M{})

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

	// log.Fatal(cursor)

	for cursor.Next(ctx) {
		// var order Order
		var order map[string]interface{}

		cursor.Decode(&order)

		orders = append(orders, order)
	}

	// if err := cursor.Err(); err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(`{"message": "` + err.Error() + `"}`))

	// 	return
	// }

	json.NewEncoder(w).Encode(orders)
}

func AddOrder(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order models.Order

	_ = json.NewDecoder(r.Body).Decode(&order)
	order.CreatedAt = time.Now()

	// Validation
	if err := validator.Validate(order); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	collection := client.Database("nggadingCarRentalSystem").Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	result, err := collection.InsertOne(ctx, order)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(result)
}
