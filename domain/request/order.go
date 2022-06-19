package request

type GetAllOrderRequest struct {
	PaginationRequest
	Search string `json:"search"`
}

type OrderRequest struct {
	CustomerID uint   `json:"customer_id"`
	Status     string `json:"status"`
	TotalPrice int    `json:"total_price"`
}
