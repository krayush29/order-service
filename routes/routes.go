package routes

import (
	"net/http"
	"order-service/controllers"
)

func InitRoutes() {
	http.HandleFunc("/orders", controllers.CreateOrder)
}
