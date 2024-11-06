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

func GetOrders(userIds []uint, statuses []string) ([]models.Order, error) {
	var orders []models.Order
	query := utils.DB

	if len(userIds) > 0 {
		query = query.Where("user_id IN (?)", userIds)
	}

	if len(statuses) > 0 {
		query = query.Where("status IN (?)", statuses)
	}

	if err := query.Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}
