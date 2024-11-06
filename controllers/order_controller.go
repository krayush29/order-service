package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"order-service/dto/request"
	"order-service/services"
	"strconv"
	"strings"
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

func GetOrders(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("user_id")
	statusStr := r.URL.Query().Get("status")

	var userIds []uint
	if userIdStr != "" {
		userIdParts := strings.Split(userIdStr, ",")
		for _, idStr := range userIdParts {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusBadRequest)
				return
			}
			userIds = append(userIds, uint(id))
		}
	}

	var statuses []string
	if statusStr != "" {
		statuses = strings.Split(statusStr, ",")
		for _, status := range statuses {
			if status != "DONE" && status != "PENDING" {
				http.Error(w, "Invalid status, Please enter valid status. Status can be DONE or PENDING", http.StatusBadRequest)
				return
			}
		}
	}

	orderResponses, err := services.GetOrders(userIds, statuses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderResponses)
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
