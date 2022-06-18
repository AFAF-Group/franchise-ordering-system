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

// GetCustomerList godoc
// @Tags         Customer
// @Summary      Get Customer List With Pagination
// @Description  Get Customer List With Pagination
// @Accept       json
// @Param        customerRequest  body  request.GetAllCustomerRequest  true  "Page: Page Number; Limit: Request Limit; Search: Reqeust for search customer"
// @Produce      json
// @Success      200  {object}  response.APIResponse{data=[]models.Customer}
// @Failure      400  {object}  response.SwaggerHTTPErrorBadRequestValidation
// @Failure      401  {object}  response.SwaggerHTTPErrorUnauthorized
// @Failure      404  {object}  response.SwaggerHTTPErrorNotFound
// @Failure      500  {object}  response.SwaggerHTTPErrorInternalServerError
// @Security     ApiKeyAuth
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

// CreateCustomer godoc
// @Tags         Customer
// @Summary      Create New Customer
// @Description  Create New Customer
// @Accept       json
// @Param        customerRequest  body  request.CreateCustomerRequest  true  "Name: Customer name; Email: Customer email; Phone: Customer phone number"
// @Produce      json
// @Success      200  {object}  response.APIResponse{data=models.Customer}
// @Failure      400  {object}  response.SwaggerHTTPErrorBadRequestValidation
// @Failure      401  {object}  response.SwaggerHTTPErrorUnauthorized
// @Failure      404  {object}  response.SwaggerHTTPErrorNotFound
// @Failure      500  {object}  response.SwaggerHTTPErrorInternalServerError
// @Security     ApiKeyAuth
// @Router       /customers [get]
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
