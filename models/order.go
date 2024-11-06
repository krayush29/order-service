package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	RestaurantID uint          `json:"restaurant_id"`
	UserID       uint          `json:"user_id"`
	MenuItemIDs  pq.Int64Array `gorm:"type:integer[]" json:"menu_item_ids"`
	Status       string        `json:"status" gorm:"default:PENDING"`
}
