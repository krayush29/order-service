package controllers

import (
	"encoding/json"
	"net/http"
	"order-service/dto/request"
	"order-service/services"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderRequest request.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderResponse, err := services.CreateOrder(orderRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderResponse)
}
