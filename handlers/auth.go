package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/alimtegar/nggading-car-rental-system/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
)

// Login Handler
func Login(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credential models.Credential
	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&credential)

	collection := client.Database("nggadingCarRentalSystem").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	err := collection.FindOne(ctx, models.Credential{Username: credential.Username}).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credential.Password))

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":      user.ID,
		"username": user.Username,
	})
	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
	}

	json.NewEncoder(w).Encode(models.Token{Token: tokenString})
}

// Register Handler
func Register(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	user.Password = string(hash)
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
