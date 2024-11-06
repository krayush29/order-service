package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"order-service/controllers"
)

func InitRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{orderId}", controllers.GetOrder).Methods("GET")
	r.HandleFunc("/orders", controllers.GetOrders).Methods("GET")
	r.HandleFunc("/orders/{orderId}", controllers.UpdateOrder).Methods("PUT")
	http.Handle("/", r)
}
