package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"order-service/dto/request"
	"order-service/dto/response"
	"testing"
)

type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) CreateOrder(orderRequest request.OrderRequest) (response.OrderResponse, error) {
	args := m.Called(orderRequest)
	return args.Get(0).(response.OrderResponse), args.Error(1)
}

func TestCreateOrder_BadRequest(t *testing.T) {
	req, err := http.NewRequest("POST", "/orders", bytes.NewBuffer([]byte("invalid json")))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateOrder)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetOrderInvalidOrderID(t *testing.T) {
	req, err := http.NewRequest("GET", "/orders/invalid", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOrder)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid order ID")
}

func TestUpdateOrderSuccess(t *testing.T) {
	orderUpdateRequest := request.OrderUpdateRequest{Status: "DONE"}
	body, _ := json.Marshal(orderUpdateRequest)
	req, err := http.NewRequest("PUT", "/orders/1", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	mockService := new(MockOrderService)
	mockService.On("UpdateOrder", mock.Anything).Return(response.OrderResponse{}, nil)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UpdateOrder(w, r)
	})
	handler.ServeHTTP(rr, req)
}

func TestUpdateOrderInvalidRequestBody(t *testing.T) {
	req, err := http.NewRequest("PUT", "/orders/1", bytes.NewBuffer([]byte("invalid json")))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateOrder)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid request body")
}

func TestUpdateOrderInvalidStatus(t *testing.T) {
	orderUpdateRequest := request.OrderUpdateRequest{Status: "INVALID"}
	body, _ := json.Marshal(orderUpdateRequest)
	req, err := http.NewRequest("PUT", "/orders/1", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateOrder)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid status")
}

func TestUpdateOrderInvalidOrderID(t *testing.T) {
	orderUpdateRequest := request.OrderUpdateRequest{Status: "DONE"}
	body, _ := json.Marshal(orderUpdateRequest)
	req, err := http.NewRequest("PUT", "/orders/invalid", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateOrder)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid order ID")
}
