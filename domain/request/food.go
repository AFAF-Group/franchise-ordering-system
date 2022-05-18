package request

type GetAllFoodRequest struct {
	PaginationRequest
	Search string `json:"search"`
}
