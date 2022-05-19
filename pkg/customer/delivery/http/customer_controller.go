package http

import (
	"net/http"

	"afaf-group.com/domain/models"
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

// GetMap godoc
// @Summary      Get Map Example
// @Description  get map
// @ID           get-map
// @Accept       json
// @Produce      json
// @Router       /customers [get]
func (c *Controller) GetCustomerList(ctx echo.Context) error {
	var customerRequest request.GetAllCustomerRequest
	if err := c.BindAndValidate(ctx, &customerRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	if customerRequest.Limit == 0 {
		customerRequest.Limit = -1
	}

	paging, err := c.customerUseCase.GetAll(ctx, &customerRequest)

	if err != nil {
		return err
	}

	customers := paging.Records.(*[]models.Customer)
	return ctx.JSON(http.StatusOK, &response.APIResponse{
		Code:     http.StatusOK,
		Message:  http.StatusText(http.StatusOK),
		Data:     customers,
		PageInfo: response.NewPageInfo().ToPageInfo(paging),
	})
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
		return err
	}

	return ctx.JSON(http.StatusOK, response.APIResponse{
		Message: "Customer successfully created!",
		Data:    customer,
	})
}
