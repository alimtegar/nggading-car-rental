package handlers

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