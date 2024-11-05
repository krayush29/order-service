package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"order-service/dto/request"
	"order-service/dto/response"
	"order-service/models"
	"order-service/repositories"
	"strconv"
)

func CreateOrder(orderRequest request.OrderRequest) (response.OrderResponse, error) {
	user, err := validateUser(orderRequest.Username, orderRequest.Password)
	if err != nil {
		return response.OrderResponse{}, err
	}

	if err := validateRestaurant(orderRequest.RestaurantID, orderRequest.MenuItemIDs); err != nil {
		return response.OrderResponse{}, err
	}

	order := models.Order{
		RestaurantID: orderRequest.RestaurantID,
		UserID:       user.ID,
		MenuItemIDs:  orderRequest.MenuItemIDs,
	}

	order, err = repositories.CreateOrder(order)
	if err != nil {
		return response.OrderResponse{}, err
	}

	return response.OrderResponse{
		OrderID:      order.ID,
		Username:     orderRequest.Username,
		RestaurantID: order.RestaurantID,
		MenuItemIDs:  order.MenuItemIDs,
	}, nil
}

func validateUser(username, password string) (response.UserResponse, error) {
	userCredentials := map[string]string{
		"username": username,
		"password": password,
	}
	jsonData, err := json.Marshal(userCredentials)
	if err != nil {
		return response.UserResponse{}, err
	}

	req, err := http.NewRequest("GET", "http://localhost:8080/users", bytes.NewBuffer(jsonData))
	if err != nil {
		return response.UserResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response.UserResponse{}, err
	}
	defer resp.Body.Close()

	var user response.UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return response.UserResponse{}, err
	}

	return user, nil
}

func validateRestaurant(restaurantID uint, menuItemIDs []int64) error {
	url := "http://localhost:8080/restaurants/" + strconv.FormatUint(uint64(restaurantID), 10)
	resp, err := http.Get(url)
	if err != nil {
		return errors.New("failed to fetch restaurant data")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid restaurant ID")
	}

	var restaurant response.RestaurantResponse
	if err := json.NewDecoder(resp.Body).Decode(&restaurant); err != nil {
		return errors.New("failed to decode restaurant data")
	}

	menuItemIDMap := make(map[uint]bool)
	for _, menuItem := range restaurant.MenuItems {
		menuItemIDMap[menuItem.MenuItemID] = true
	}

	for _, id := range menuItemIDs {
		if !menuItemIDMap[uint(id)] {
			return errors.New("invalid menu item ID")
		}
	}

	return nil
}
