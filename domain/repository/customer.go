package repository

import (
	"afaf-group.com/domain/models"
	"afaf-group.com/domain/request"
	"github.com/labstack/echo/v4"
)

type CustomerMySqlRepository interface {
	Create(ctx echo.Context, createCustomerRequest *request.CreateCustomerRequest) (*models.Customer, error)
	GetByEmail(email string) *models.Customer
}
