package usecase

import (
	"afaf-group.com/domain/repository"
	"afaf-group.com/domain/request"
	"afaf-group.com/domain/usecase"
	"afaf-group.com/pkg/common"
	"github.com/labstack/echo/v4"
)

type foodUseCase struct {
	foodMySQLRepository repository.FoodMySQLRepository
}

func NewFoodUseCase(foodMySQLRepository repository.FoodMySQLRepository) usecase.FoodUseCase {
	return &foodUseCase{
		foodMySQLRepository: foodMySQLRepository,
	}
}

func (u foodUseCase) GetAll(ctx echo.Context, foodRequest *request.GetAllFoodRequest) (*common.Pagination, error) {
	paging, err := u.foodMySQLRepository.GetAll(ctx, foodRequest)
	if err != nil {
		return nil, err
	}
	return paging, nil
}
