package request

type CreateCustomerRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type GetAllCustomerRequest struct {
	PaginationRequest
	Search string `json:"search"`
}
