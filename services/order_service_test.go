package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"order-service/dto/response"
	"order-service/models"
	"strconv"
	"testing"
)

// Mocking the repositories
type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) CreateOrder(order models.Order) (models.Order, error) {
	args := m.Called(order)
	return args.Get(0).(models.Order), args.Error(1)
}

func TestValidateUser(t *testing.T) {
	userResponse := response.UserResponse{
		ID:       1,
		Username: "username1",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userResponse)
	}))
	defer server.Close()

	var validateUser = func(username, password string) (response.UserResponse, error) {
		userCredentials := map[string]string{
			"username": username,
			"password": password,
		}
		jsonData, err := json.Marshal(userCredentials)
		if err != nil {
			return response.UserResponse{}, err
		}

		req, err := http.NewRequest("GET", server.URL+"/users", bytes.NewBuffer(jsonData))
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

	user, err := validateUser("username1", "password1")
	assert.NoError(t, err)
	assert.Equal(t, userResponse, user)
}

func TestValidateRestaurant(t *testing.T) {
	restaurantResponse := response.RestaurantResponse{
		RestaurantID: 1,
		Name:         "Test Restaurant",
		Address:      "123 Test St",
		MenuItems: []response.MenuItemResponse{
			{MenuItemID: 1, Name: "Item 1", Price: 10.0},
			{MenuItemID: 2, Name: "Item 2", Price: 20.0},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(restaurantResponse)
	}))
	defer server.Close()

	var validateRestaurant = func(restaurantID uint, menuItemIDs []int64) error {
		url := server.URL + "/restaurants/" + strconv.FormatUint(uint64(restaurantID), 10)
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

	err := validateRestaurant(1, []int64{1, 2})
	assert.NoError(t, err)
}
