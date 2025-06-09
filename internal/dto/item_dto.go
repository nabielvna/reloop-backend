package dto

type CreateItemRequest struct {
	CategoryID  uint    `json:"categoryId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type UpdateItemRequest struct {
	CategoryID  *uint    `json:"categoryId,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Quantity    *int     `json:"quantity,omitempty"`
}

type UpdateItemStatusRequest struct {
	Status *string `json:"status"`
}

type ItemResponse struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       float64          `json:"price"`
	Quantity    int              `json:"quantity"`
	Status      string           `json:"status"`
	Category    CategoryResponse `json:"category"`
	Seller      SellerResponse   `json:"seller"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type SellerResponse struct {
	ID           uint         `json:"id"`
	BusinessName string       `json:"businessName"`
	User         UserResponse `json:"user"`
}

type BrowseItemsRequest struct {
	CategoryID *uint    `json:"categoryId,omitempty"`
	MinPrice   *float64 `json:"minPrice,omitempty"`
	MaxPrice   *float64 `json:"maxPrice,omitempty"`
	Search     *string  `json:"search,omitempty"`
	Status     *string  `json:"status,omitempty"`
}
