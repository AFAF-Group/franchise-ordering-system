package repository

import (
	"afaf-group.com/domain/models"
	"afaf-group.com/domain/request"
	"github.com/labstack/echo/v4"
)

type AuthMySQLRepository interface {
	Login(ctx echo.Context, loginRequest *request.LoginRequest) (*models.User, error)
}
