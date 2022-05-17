package usecase

import (
	"afaf-group.com/domain/models"
	"afaf-group.com/domain/request"
	"github.com/labstack/echo/v4"
)

type CustomerUseCase interface {
	Store(ctx echo.Context, a *request.CreateCustomerRequest) (*models.Customer, error)
}
