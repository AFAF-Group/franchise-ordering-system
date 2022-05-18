package response

type APIResponse struct {
	Code     int         `json:"code,omitempty"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
	PageInfo interface{} `json:"page_info,omitempty"`
	Errors   interface{} `json:"errors,omitempty"`
} // @name  APIResponse
