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

func (r authMySQLRepository) FindOneByEmail(ctx echo.Context, loginRequest *request.AuthRequest) (*models.User, error) {
	var user models.User
	err := r.db.Model(&models.User{}).First(&user, "email=?", loginRequest.Email).Error

	return &user, err
}

func (r authMySQLRepository) Register(ctx echo.Context, user *models.User) error {
	return r.db.Model(&models.User{}).Create(&user).Error
}
