package usecase

import (
	"afaf-group.com/domain/models"
	"afaf-group.com/domain/repository"
	"afaf-group.com/domain/request"
	"afaf-group.com/domain/usecase"
	"github.com/labstack/echo/v4"
)

type authUseCase struct {
	authMySQLRepository repository.AuthMySQLRepository
}

func NewAuthUseCase(authMySQLRepository repository.AuthMySQLRepository) usecase.AuthUseCase {
	return &authUseCase{
		authMySQLRepository: authMySQLRepository,
	}
}

func (u authUseCase) Login(ctx echo.Context, loginRequest *request.LoginRequest) (*models.User, error) {
	user, err := u.authMySQLRepository.Login(ctx, loginRequest)
	if err != nil {
		return nil, err
	}

	return user, err
}
