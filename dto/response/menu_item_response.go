package response

type MenuItemResponse struct {
	MenuItemID uint    `json:"menuItemId"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
}
