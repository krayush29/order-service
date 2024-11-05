package request

import (
	_ "github.com/lib/pq"
)

type OrderRequest struct {
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	RestaurantID uint    `json:"restaurant_id"`
	MenuItemIDs  []int64 `json:"menu_item_ids"`
}
