package mysql

import (
	"afaf-group.com/domain/models"
	"afaf-group.com/domain/repository"
	"afaf-group.com/domain/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type authMySQLRepository struct {
	db *gorm.DB
}

func NewAuthMySQLRepository(db *gorm.DB) repository.AuthMySQLRepository {
	return &authMySQLRepository{db: db}
}

func (r authMySQLRepository) Login(ctx echo.Context, loginRequest *request.LoginRequest) (*models.User, error) {
	var user models.User
	err := r.db.Model(&models.User{}).First(&user, "email=? AND password=?", loginRequest.Email, loginRequest.Password).Error

	return &user, err
}
