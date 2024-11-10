package repositories

import (
	"order-service/models"
	"order-service/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Order{})
	utils.DB = db
	return db
}

func TestGetOrder(t *testing.T) {
	db := setupTestDB()

	order := models.Order{
		RestaurantID: 1,
		UserID:       1,
		MenuItemIDs:  []int64{1, 2, 3},
		Status:       "PENDING",
	}
	db.Create(&order)

	retrievedOrder, err := GetOrder(order.ID)
	assert.NoError(t, err)
	assert.Equal(t, order.ID, retrievedOrder.ID)
	assert.Equal(t, order.RestaurantID, retrievedOrder.RestaurantID)
	assert.Equal(t, order.UserID, retrievedOrder.UserID)
	assert.Equal(t, order.MenuItemIDs, retrievedOrder.MenuItemIDs)
	assert.Equal(t, order.Status, retrievedOrder.Status)
}

func TestUpdateOrder(t *testing.T) {
	db := setupTestDB()

	order := models.Order{
		RestaurantID: 1,
		UserID:       1,
		MenuItemIDs:  []int64{1, 2, 3},
		Status:       "PENDING",
	}
	db.Create(&order)

	updatedOrder, err := UpdateOrder(order.ID, "DONE")
	assert.NoError(t, err)
	assert.Equal(t, "DONE", updatedOrder.Status)
}

func TestGetOrders(t *testing.T) {
	db := setupTestDB()

	order1 := models.Order{
		RestaurantID: 1,
		UserID:       1,
		MenuItemIDs:  []int64{1, 2, 3},
		Status:       "PENDING",
	}
	order2 := models.Order{
		RestaurantID: 2,
		UserID:       2,
		MenuItemIDs:  []int64{4, 5, 6},
		Status:       "DONE",
	}
	db.Create(&order1)
	db.Create(&order2)

	userIds := []uint{1, 2}
	statuses := []string{"PENDING", "DONE"}

	orders, err := GetOrders(userIds, statuses)
	assert.NoError(t, err)
	assert.Len(t, orders, 2)
}
