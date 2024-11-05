package response

type OrderResponse struct {
	OrderID      uint    `json:"order_id"`
	Username     string  `json:"username"`
	RestaurantID uint    `json:"restaurant_id"`
	MenuItemIDs  []int64 `json:"menu_item_ids"`
}
