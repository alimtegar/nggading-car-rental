package handlers

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
