package http

import (
	"net/http"

	"afaf-group.com/domain/request"
	"afaf-group.com/domain/response"
	"afaf-group.com/domain/usecase"
	"afaf-group.com/pkg/common"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	common.Controller
	customerUseCase usecase.CustomerUseCase
}

func NewController(customerUseCase usecase.CustomerUseCase) *Controller {
	return &Controller{customerUseCase: customerUseCase}
}

func (c *Controller) CreateCustomer(ctx echo.Context) error {
	var customerRequest request.CreateCustomerRequest
	if err := c.BindAndValidate(ctx, &customerRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	customer, err := c.customerUseCase.Store(ctx, &customerRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.APIResponse{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.APIResponse{
		Message: "Customer successfully created!",
		Data:    customer,
	})
}
