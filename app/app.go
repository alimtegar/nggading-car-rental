package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/alimtegar/nggading-car-rental-system/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// App Model
type App struct {
	Router *mux.Router
	Client *mongo.Client
}

// Initialize App
func (a *App) Initialize() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	a.Client, _ = mongo.Connect(ctx, clientOptions)
	a.Router = mux.NewRouter()

	a.setRoutes()
}

func (a *App) setRoutes() {
	// Users Routes
	a.Router.HandleFunc("/users", a.GetUsers).Methods("GET")
	a.Router.HandleFunc("/users/{id}", a.GetUser).Methods("GET")
	a.Router.HandleFunc("/users", a.AddUser).Methods("POST")
	a.Router.HandleFunc("/users/{id}", a.UpdateUser).Methods("PUT")
	a.Router.HandleFunc("/users/{id}", a.DeleteUser).Methods("DELETE")

	// Cars Routes
	a.Router.HandleFunc("/cars", a.GetCars).Methods("GET")
	a.Router.HandleFunc("/cars/{id}", a.GetCar).Methods("GET")
	a.Router.HandleFunc("/cars", a.AddCar).Methods("POST")
	a.Router.HandleFunc("/cars/{id}", a.UpdateCar).Methods("PUT")
	a.Router.HandleFunc("/cars/{id}", a.DeleteCar).Methods("DELETE")

	// Orders Routes
	a.Router.HandleFunc("/orders", a.GetOrders).Methods("GET")
	a.Router.HandleFunc("/orders", a.AddOrder).Methods("POST")

	// Hello World Routes
	a.Router.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "Hello World"}`))
	}).Methods("GET")
}

// Run App
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func (a *App) GetUsers(w http.ResponseWriter, r *http.Request)   { handlers.GetUsers(a.Client, w, r) }
func (a *App) GetUser(w http.ResponseWriter, r *http.Request)    { handlers.GetUser(a.Client, w, r) }
func (a *App) AddUser(w http.ResponseWriter, r *http.Request)    { handlers.AddUser(a.Client, w, r) }
func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) { handlers.UpdateUser(a.Client, w, r) }
func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) { handlers.DeleteUser(a.Client, w, r) }

func (a *App) GetCars(w http.ResponseWriter, r *http.Request)   { handlers.GetCars(a.Client, w, r) }
func (a *App) GetCar(w http.ResponseWriter, r *http.Request)    { handlers.GetCar(a.Client, w, r) }
func (a *App) AddCar(w http.ResponseWriter, r *http.Request)    { handlers.AddCar(a.Client, w, r) }
func (a *App) UpdateCar(w http.ResponseWriter, r *http.Request) { handlers.UpdateCar(a.Client, w, r) }
func (a *App) DeleteCar(w http.ResponseWriter, r *http.Request) { handlers.DeleteCar(a.Client, w, r) }

func (a *App) GetOrders(w http.ResponseWriter, r *http.Request) { handlers.GetOrders(a.Client, w, r) }
func (a *App) AddOrder(w http.ResponseWriter, r *http.Request)  { handlers.AddOrder(a.Client, w, r) }
