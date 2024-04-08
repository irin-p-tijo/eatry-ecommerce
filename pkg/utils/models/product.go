package models

type AddProduct struct {
	ID         int     `json:"id"`
	CategoryID int     `json:"category_id"`
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	Stock      int     `json:"stock"`
	Price      float64 `json:"price"`
}
type ProductResponse struct {
	ID         int     `json:"id" `
	CategoryID int     `json:"category_id"`
	Name       string  `json:"name" `
	Quantity   int     `json:"quantity"`
	Stock      int     `json:"stock"`
	Price      float64 `json:"price"`
}
type SearchItems struct {
	Name string `json:"name" binding:"required"`
}
type ProductBrief struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	Stock         int     `json:"stock"`
	ProductStatus string  `json:"product_status"`
}
type ProductDetails struct {
	Name       string  `json:"name"`
	FinalPrice float64 `json:"final_price"`
	Price      float64 `json:"price" `
	Total      float64 `json:"total_price"`
	Quantity   int     `json:"quantity"`
}
