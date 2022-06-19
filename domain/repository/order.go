package repository

import (
	"afaf-group.com/domain/models"
	"afaf-group.com/domain/request"
	"afaf-group.com/pkg/common"
	"github.com/labstack/echo/v4"
)

type OrderMySQLRepository interface {
	FindAllWithPagination(ctx echo.Context, orderRequest *request.GetAllOrderRequest) (*common.Pagination, error)
	Create(ctx echo.Context, order *models.Order) error
	Update(ctx echo.Context, order *models.Order) error
}
