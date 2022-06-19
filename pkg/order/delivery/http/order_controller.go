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
	orderUseCase usecase.OrderUseCase
}

func NewController(orderUseCase usecase.OrderUseCase) *Controller {
	return &Controller{orderUseCase: orderUseCase}
}

// GetAllOrder godoc
// @Tags         Orders
// @Summary      Get All Order
// @Description  Show List All Order
// @Accept       json
// @Param        orderRequest  body  request.GetAllOrderRequest  false  "Page: page number; Limit: limit number; Search: search order"
// @Produce      json
// @Success      200  {object}  response.APIResponse{data=[]models.Order}
// @Failure      400  {object}  response.SwaggerHTTPErrorBadRequestValidation
// @Failure      401  {object}  response.SwaggerHTTPErrorUnauthorized
// @Failure      404  {object}  response.SwaggerHTTPErrorNotFound
// @Failure      500  {object}  response.SwaggerHTTPErrorInternalServerError
// @Security     ApiKeyAuth
// @Router       /orders [get]
func (c Controller) GetAll(ctx echo.Context) error {
	var orderRequest request.GetAllOrderRequest
	if err := c.BindAndValidate(ctx, &orderRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	if orderRequest.Limit == 0 {
		orderRequest.Limit = -1
	}
	paging, err := c.orderUseCase.FindAllWithPagination(ctx, &orderRequest)
	if err != nil {
		return err
	}

	orders := paging.Records.(*[]models.Order)
	return ctx.JSON(http.StatusOK, &response.APIResponse{
		Code:     http.StatusOK,
		Message:  http.StatusText(http.StatusOK),
		Data:     orders,
		PageInfo: response.NewPageInfo().ToPageInfo(paging),
	})
}

// CreateOrder godoc
// @Tags         Orders
// @Summary      Create Order
// @Description  Create New Order
// @Accept       json
// @Param        orderRequest  body  request.OrderRequest  true  "CustomerID: ID from seleted customer; Status: status order; TotalPrice: total price order"
// @Produce      json
// @Success      200  {object}  response.APIResponse{}
// @Failure      400  {object}  response.SwaggerHTTPErrorBadRequestValidation
// @Failure      401  {object}  response.SwaggerHTTPErrorUnauthorized
// @Failure      404  {object}  response.SwaggerHTTPErrorNotFound
// @Failure      500  {object}  response.SwaggerHTTPErrorInternalServerError
// @Security     ApiKeyAuth
// @Router       /orders/create [post]
func (c Controller) Create(ctx echo.Context) error {
	var orderRequest request.OrderRequest
	if err := c.BindAndValidate(ctx, &orderRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	err := c.orderUseCase.Create(ctx, &orderRequest)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.APIResponse{
		Message: "Order successfully created!",
	})
}

// UpdateOrder godoc
// @Tags         Orders
// @Summary      Update Order
// @Description  Update Existing Order
// @Accept       json
// @Param        orderRequest  body  request.OrderRequest  false  "CustomerID: ID from seleted customer; Status: status order; TotalPrice: total price order"
// @Produce      json
// @Success      200  {object}  response.APIResponse{}
// @Failure      400  {object}  response.SwaggerHTTPErrorBadRequestValidation
// @Failure      401  {object}  response.SwaggerHTTPErrorUnauthorized
// @Failure      404  {object}  response.SwaggerHTTPErrorNotFound
// @Failure      500  {object}  response.SwaggerHTTPErrorInternalServerError
// @Security     ApiKeyAuth
// @Router       /orders/update [put]
func (c Controller) Update(ctx echo.Context) error {
	var orderRequest request.OrderRequest
	if err := c.BindAndValidate(ctx, &orderRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	err := c.orderUseCase.Update(ctx, &orderRequest)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.APIResponse{
		Message: "Order successfully updated!",
	})
}
