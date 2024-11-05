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
