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
