package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"order-service/controllers"
)

func InitRoutes() {
	http.HandleFunc("/orders", controllers.CreateOrder)

	r := mux.NewRouter()
	r.HandleFunc("/orders/{orderId}", controllers.GetOrder).Methods("GET")
	http.Handle("/", r)
}
