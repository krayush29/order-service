package response

type OrderResponse struct {
	OrderID      uint    `json:"order_id"`
	UserId       uint    `json:"user_id"`
	RestaurantID uint    `json:"restaurant_id"`
	MenuItemIDs  []int64 `json:"menu_item_ids"`
}
