package mysql

import (
	"afaf-group.com/domain/models"
	"afaf-group.com/domain/repository"
	"afaf-group.com/domain/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type customerMySQLRepository struct {
	db *gorm.DB
}

func NewCustomerMySQLRepository(db *gorm.DB) repository.CustomerMySqlRepository {
	return &customerMySQLRepository{db: db}
}

func (r customerMySQLRepository) GetByEmail(email string) *models.Customer {
	var customer = models.Customer{}
	r.db.Where(&models.Customer{Email: email}).First(&customer)

	return &customer
}

func (r customerMySQLRepository) Create(ctx echo.Context, createRequest *request.CreateCustomerRequest) (*models.Customer, error) {
	var customer = models.Customer{
		Name:  createRequest.Name,
		Email: createRequest.Email,
		Phone: createRequest.Phone,
	}
	err := r.db.Create(&customer).Error

	return &customer, err
}
