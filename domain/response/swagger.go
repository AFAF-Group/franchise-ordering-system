package response

type SwaggerHTTPStatusOK struct {
	Message string      `json:"message"  example:"OK"`
	Data    interface{} `json:"data"`
} // @name SwaggerHttpStatusOK

type SwaggerHTTPErrorBadRequest struct {
	Message string `json:"message"  example:"Bad Request"`
} // @name HttpErrorBadRequest

type SwaggerHTTPErrorBadRequestValidation struct {
	Message string `json:"message"  example:"Error Validation"`
	Errors  []struct {
		Field   string `json:"field"    example:"email"`
		Message string `json:"message"  example:"The email must be a valid email address."`
	} `json:"errors"`
} // @name HttpErrorBadRequestValidation

type SwaggerHTTPErrorNotFound struct {
	Message string `json:"message"  example:"record not found"`
} // @name HttpErrorNotFound

type SwaggerHTTPErrorUnauthorized struct {
	Message string `json:"message"  example:"missing key in request header"`
} // @name HttpErrorUnauthorized

type SwaggerHTTPErrorInternalServerError struct {
	Message string `json:"message"  example:"Internal Server Error"`
} // @name HttpErrorInternalServerError
