package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"order-service/dto/request"
	"order-service/services"
	"strconv"
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

func GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, err := strconv.Atoi(vars["orderId"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	orderResponse, err := services.GetOrder(uint(orderId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderResponse)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var orderUpdateRequest request.OrderUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&orderUpdateRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if orderUpdateRequest.Status != "DONE" {
		http.Error(w, "Invalid status, Please enter valid status. Status can be DONE", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	orderId, err := strconv.Atoi(vars["orderId"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	orderResponse, err := services.UpdateOrder(uint(orderId), orderUpdateRequest.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderResponse)
}
