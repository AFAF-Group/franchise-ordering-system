package repository

import (
	"afaf-group.com/domain/models"
	"afaf-group.com/domain/request"
	"github.com/labstack/echo/v4"
)

type AuthMySQLRepository interface {
	FindOneByEmail(ctx echo.Context, loginRequest *request.AuthRequest) (*models.User, error)
	Register(ctx echo.Context, user *models.User) error
}
