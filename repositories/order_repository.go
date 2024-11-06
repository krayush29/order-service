package repositories

import (
	"order-service/models"
	"order-service/utils"
)

func CreateOrder(order models.Order) (models.Order, error) {
	if err := utils.DB.Create(&order).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func GetOrder(orderId uint) (models.Order, error) {
	var order models.Order
	if err := utils.DB.First(&order, orderId).Error; err != nil {
		return order, err
	}
	return order, nil
}

func UpdateOrder(orderId uint, status string) (models.Order, error) {
	var order models.Order
	if err := utils.DB.First(&order, orderId).Error; err != nil {
		return order, err
	}

	order.Status = status
	if err := utils.DB.Save(&order).Error; err != nil {
		return order, err
	}

	return order, nil
}
