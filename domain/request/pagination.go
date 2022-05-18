package request

type PaginationRequest struct {
	Page  int `json:"page"  form:"page"   query:"page"   example:"1"`  // This is page, if load all data set value to 0
	Limit int `json:"limit" form:"limit"  query:"limit"  example:"10"` // This is limit, if load all data set value to -1
}
