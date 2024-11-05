package main

import (
	"log"
	"net/http"
	"order-service/routes"
	"order-service/utils"
)

func main() {
	utils.InitDB()
	routes.InitRoutes()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
