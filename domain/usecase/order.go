package usecase

import (
	"afaf-group.com/domain/request"
	"afaf-group.com/pkg/common"
	"github.com/labstack/echo/v4"
)

type OrderUseCase interface {
	FindAllWithPagination(ctx echo.Context, orderRequest *request.GetAllOrderRequest) (*common.Pagination, error)
	Create(ctx echo.Context, orderRequest *request.OrderRequest) error
	Update(ctx echo.Context, orderRequest *request.OrderRequest) error
}
