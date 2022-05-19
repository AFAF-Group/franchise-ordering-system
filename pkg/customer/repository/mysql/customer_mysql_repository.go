package mysql

import (
	"fmt"

	"afaf-group.com/domain/models"
	"afaf-group.com/domain/repository"
	"afaf-group.com/domain/request"
	"afaf-group.com/pkg/common"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type customerMySQLRepository struct {
	db *gorm.DB
}

func NewCustomerMySQLRepository(db *gorm.DB) repository.CustomerMySqlRepository {
	return &customerMySQLRepository{db: db}
}

func (r customerMySQLRepository) GetAll(ctx echo.Context, customerRequest *request.GetAllCustomerRequest) (*common.Pagination, error) {
	var customer []models.Customer
	query := r.db.Model(&models.Customer{}).Find(&customer)
	if customerRequest.Search != "" {
		query = query.Where("email ILike ?", customerRequest.Search+"%")
	}

	paging, err := common.NewPagination().Paging(query, customerRequest.Page, customerRequest.Limit, []string{"id"}, &customer)

	return paging, err
}

func (r customerMySQLRepository) GetByEmail(email string) *models.Customer {
	var customer = models.Customer{}
	if result := r.db.Where(&models.Customer{Email: email}).First(&customer); result.Error != nil {
		// error handling...
	}
	fmt.Println(customer.ID, "fadfsdf")

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
