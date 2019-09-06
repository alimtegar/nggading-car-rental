package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/validator.v2"
)

var client *mongo.Client

// User Model
type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty" validate:"nonzero"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty" validate:"nonzero"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" validate:"nonzero"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" validate:"nonzero"`
}

// Car Model
type Car struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Brand       string             `json:"brand,omitempty" bson:"brand,omitempty" validate:"nonzero"`
	Model       string             `json:"model,omitempty" bson:"model,omitempty" validate:"nonzero"`
	Year        int32              `json:"year,omitempty" bson:"year,omitempty" validate:"nonzero"`
	Color       string             `json:"color,omitempty" bson:"color,omitempty" validate:"nonzero"`
	PlatNumber  string             `json:"platNumber,omitempty" bson:"platNumber,omitempty" validate:"nonzero"`
	Stock       int32              `json:"stock,omitempty" bson:"stock,omitempty" validate:"nonzero"`
	Price       int32              `json:"price,omitempty" bson:"price,omitempty" validate:"nonzero"`
	Description string             `json:"description,omitempty" bson:"description,omitempty" validate:"nonzero"`
	CreatedAt   time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" validate:"nonzero"`
	UpdatedAt   time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" validate:"nonzero"`
}

// Order Model
type Order struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CustomerID primitive.ObjectID `json:"customer,omitempty" bson:"customer,omitempty" validate:"nonzero"`
	Car        primitive.ObjectID `json:"car,omitempty" bson:"car,omitempty" validate:"nonzero"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" validate:"nonzero"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" validate:"nonzero"`
}

// User Controllers
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []User

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
		var user User

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

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var user User

	collection := client.Database("nggadingCarRentalSystem").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	err := collection.FindOne(ctx, User{ID: id}).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(user)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User

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

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var user User

	_ = json.NewDecoder(r.Body).Decode(&user)
	user.UpdatedAt = time.Now()

	collection := client.Database("nggadingCarRentalSystem").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	result, err := collection.UpdateOne(ctx, User{ID: id}, bson.M{"$set": user})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(result)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := client.Database("nggadingCarRentalSystem").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	result, err := collection.DeleteOne(ctx, User{ID: id})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(result)
}

// Car Controllers
func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cars []Car

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
		var car Car

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

func getCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var car Car

	collection := client.Database("nggadingCarRentalSystem").Collection("cars")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	err := collection.FindOne(ctx, Car{ID: id}).Decode(&car)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(car)
}

func addCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var car Car

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

func updateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var car Car

	_ = json.NewDecoder(r.Body).Decode(&car)
	car.UpdatedAt = time.Now()

	collection := client.Database("nggadingCarRentalSystem").Collection("cars")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	result, err := collection.UpdateOne(ctx, Car{ID: id}, bson.M{"$set": car})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(result)
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := client.Database("nggadingCarRentalSystem").Collection("cars")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	result, err := collection.DeleteOne(ctx, Car{ID: id})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	json.NewEncoder(w).Encode(result)
}

// Order Controllers
func getOrders(w http.ResponseWriter, r *http.Request) {
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
		{{"$unset", "customer_id"}},
		{{"$unset", "car_id"}},
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

func addOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order Order

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

func main() {
	log.Println("Starting the application...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()

	// Users Routes
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", addUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	// Cars Routes
	router.HandleFunc("/cars", getCars).Methods("GET")
	router.HandleFunc("/cars/{id}", getCar).Methods("GET")
	router.HandleFunc("/cars", addCar).Methods("POST")
	router.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", deleteCar).Methods("DELETE")

	// Orders Routes
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/orders", addOrder).Methods("POST")

	http.ListenAndServe(":3001", router)
}
