package usecase

import (
	"afaf-group.com/domain/models"
	"afaf-group.com/domain/repository"
	"afaf-group.com/domain/request"
	"afaf-group.com/domain/usecase"
	"afaf-group.com/pkg/common"
	"github.com/labstack/echo/v4"
)

type orderUseCase struct {
	orderMySQLRepository repository.OrderMySQLRepository
}

func NewOrderUseCase(orderMySQLRepository repository.OrderMySQLRepository) usecase.OrderUseCase {
	return &orderUseCase{
		orderMySQLRepository: orderMySQLRepository,
	}
}

func (u *orderUseCase) FindAllWithPagination(ctx echo.Context, orderRequest *request.GetAllOrderRequest) (*common.Pagination, error) {
	paging, err := u.orderMySQLRepository.FindAllWithPagination(ctx, orderRequest)

	if err != nil {
		return nil, err
	}

	return paging, nil
}

func (u *orderUseCase) Create(ctx echo.Context, orderRequest *request.OrderRequest) error {
	order := models.Order{
		Status:     orderRequest.Status,
		TotalPrice: orderRequest.TotalPrice,
		CustomerID: orderRequest.CustomerID,
	}
	return u.orderMySQLRepository.Create(ctx, &order)
}

func (u *orderUseCase) Update(ctx echo.Context, orderRequest *request.OrderRequest) error {
	order := models.Order{
		Status:     orderRequest.Status,
		TotalPrice: orderRequest.TotalPrice,
		CustomerID: orderRequest.CustomerID,
	}
	return u.orderMySQLRepository.Update(ctx, &order)
}
