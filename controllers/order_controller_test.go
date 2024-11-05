package controllers

import (
	"bytes"
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
