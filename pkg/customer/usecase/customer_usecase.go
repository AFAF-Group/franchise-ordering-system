package usecase

import (
	"afaf-group.com/domain"
	"afaf-group.com/domain/models"
	"afaf-group.com/domain/repository"
	"afaf-group.com/domain/request"
	"afaf-group.com/domain/usecase"
	"github.com/labstack/echo/v4"
)

type customerUseCase struct {
	customerMySQLRepository repository.CustomerMySqlRepository
}

func NewCustomerUseCase(customerMySQLRepository repository.CustomerMySqlRepository) usecase.CustomerUseCase {
	return &customerUseCase{
		customerMySQLRepository: customerMySQLRepository,
	}
}

func (u customerUseCase) Store(ctx echo.Context, createCustomerRequest *request.CreateCustomerRequest) (*models.Customer, error) {
	// check user exists
	customerExist := u.customerMySQLRepository.GetByEmail(createCustomerRequest.Email)
	if customerExist.Email != "" {
		return nil, domain.ErrConflict
	}
	customer, err := u.customerMySQLRepository.Create(ctx, createCustomerRequest)
	if err != nil {
		return nil, err
	}

	return customer, err
}
