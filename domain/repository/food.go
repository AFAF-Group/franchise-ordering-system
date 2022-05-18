package repository

import (
	"afaf-group.com/domain/request"
	"afaf-group.com/pkg/common"
	"github.com/labstack/echo/v4"
)

type FoodMySQLRepository interface {
	GetAll(ctx echo.Context, foodRequest *request.GetAllFoodRequest) (*common.Pagination, error)
}
