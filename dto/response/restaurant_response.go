package response

type RestaurantResponse struct {
	RestaurantID uint               `json:"restaurantId"`
	Name         string             `json:"name"`
	Address      string             `json:"address"`
	MenuItems    []MenuItemResponse `json:"menuItems"`
}
