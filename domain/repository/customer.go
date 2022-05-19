package repository

import (
	"afaf-group.com/domain/models"
	"afaf-group.com/domain/request"
	"afaf-group.com/pkg/common"
	"github.com/labstack/echo/v4"
)

type CustomerMySqlRepository interface {
	Create(ctx echo.Context, createCustomerRequest *request.CreateCustomerRequest) (*models.Customer, error)
	GetByEmail(email string) *models.Customer
	GetAll(ctx echo.Context, customerRequest *request.GetAllCustomerRequest) (*common.Pagination, error)
}
