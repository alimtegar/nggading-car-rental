package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/alimtegar/nggading-car-rental-system/handlers"
	"github.com/alimtegar/nggading-car-rental-system/middlewares"
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
	a.Router.HandleFunc("/users", middlewares.ValidateUser(a.GetUsers)).Methods("GET")
	a.Router.HandleFunc("/users/{id}", middlewares.ValidateUser(a.GetUser)).Methods("GET")
	a.Router.HandleFunc("/users", middlewares.ValidateUser(a.AddUser)).Methods("POST")
	a.Router.HandleFunc("/users/{id}", middlewares.ValidateUser(a.UpdateUser)).Methods("PUT")
	a.Router.HandleFunc("/users/{id}", middlewares.ValidateUser(a.DeleteUser)).Methods("DELETE")

	// Cars Routes
	a.Router.HandleFunc("/cars", middlewares.ValidateUser(a.GetCars)).Methods("GET")
	a.Router.HandleFunc("/cars/{id}", middlewares.ValidateUser(a.GetCar)).Methods("GET")
	a.Router.HandleFunc("/cars", middlewares.ValidateUser(a.AddCar)).Methods("POST")
	a.Router.HandleFunc("/cars/{id}", middlewares.ValidateUser(a.UpdateCar)).Methods("PUT")
	a.Router.HandleFunc("/cars/{id}", middlewares.ValidateUser(a.DeleteCar)).Methods("DELETE")

	// Customers Routes
	a.Router.HandleFunc("/customers", middlewares.ValidateUser(a.GetCustomers)).Methods("GET")
	a.Router.HandleFunc("/customers/{id}", middlewares.ValidateUser(a.GetCustomer)).Methods("GET")
	a.Router.HandleFunc("/customers", middlewares.ValidateUser(a.AddCustomer)).Methods("POST")
	a.Router.HandleFunc("/customers/{id}", middlewares.ValidateUser(a.UpdateCustomer)).Methods("PUT")
	a.Router.HandleFunc("/customers/{id}", middlewares.ValidateUser(a.DeleteCustomer)).Methods("DELETE")

	// Orders Routes
	a.Router.HandleFunc("/orders", middlewares.ValidateUser(a.GetOrders)).Methods("GET")
	a.Router.HandleFunc("/orders", middlewares.ValidateUser(a.AddOrder)).Methods("POST")
	a.Router.HandleFunc("/orders/{id}", middlewares.ValidateUser(a.DeleteOrder)).Methods("DELETE")

	// Auth Routes
	a.Router.HandleFunc("/login", a.Login).Methods("POST")
	a.Router.HandleFunc("/register", a.Register).Methods("POST")

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

func (a *App) GetCustomers(w http.ResponseWriter, r *http.Request) {
	handlers.GetCustomers(a.Client, w, r)
}
func (a *App) GetCustomer(w http.ResponseWriter, r *http.Request) {
	handlers.GetCustomer(a.Client, w, r)
}
func (a *App) AddCustomer(w http.ResponseWriter, r *http.Request) {
	handlers.AddCustomer(a.Client, w, r)
}
func (a *App) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateCustomer(a.Client, w, r)
}
func (a *App) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteCustomer(a.Client, w, r)
}

func (a *App) GetOrders(w http.ResponseWriter, r *http.Request) { handlers.GetOrders(a.Client, w, r) }
func (a *App) AddOrder(w http.ResponseWriter, r *http.Request)  { handlers.AddOrder(a.Client, w, r) }
func (a *App) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteOrder(a.Client, w, r)
}

func (a *App) Login(w http.ResponseWriter, r *http.Request)    { handlers.Login(a.Client, w, r) }
func (a *App) Register(w http.ResponseWriter, r *http.Request) { handlers.Register(a.Client, w, r) }
