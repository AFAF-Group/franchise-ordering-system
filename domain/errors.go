package domain

import (
	"errors"
	"net/http"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
)

// map error to get proper status code
// any others error not mapped here will result 500 status code
func NewErrorStatusCodeMaps() map[error]int {
	var errorStatusCodeMaps = make(map[error]int)
	errorStatusCodeMaps[ErrNotFound] = http.StatusNotFound
	errorStatusCodeMaps[ErrConflict] = http.StatusConflict
	errorStatusCodeMaps[ErrBadParamInput] = http.StatusBadRequest
	errorStatusCodeMaps[ErrInternalServerError] = http.StatusInternalServerError
	return errorStatusCodeMaps
}
