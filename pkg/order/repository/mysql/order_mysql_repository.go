package mysql

import (
	"afaf-group.com/domain/models"
	"afaf-group.com/domain/repository"
	"afaf-group.com/domain/request"
	"afaf-group.com/pkg/common"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type orderMySQLRepository struct {
	db *gorm.DB
}

func NewOrderMySQLRepository(db *gorm.DB) repository.OrderMySQLRepository {
	return &orderMySQLRepository{db: db}
}

func (r *orderMySQLRepository) FindAllWithPagination(ctx echo.Context, orderRequest *request.GetAllOrderRequest) (*common.Pagination, error) {
	var order []models.Order
	query := r.db.Model(&models.Order{}).Find(&order)
	if orderRequest.Search != "" {
		query = query.Where("customer_id=?", orderRequest.Search)
	}
	paging, err := common.NewPagination().Paging(query, orderRequest.Page, orderRequest.Limit, []string{"id"}, &order)
	return paging, err
}

func (r *orderMySQLRepository) Create(ctx echo.Context, order *models.Order) error {
	return r.db.Model(&models.Order{}).Create(&order).Error
}

func (r *orderMySQLRepository) Update(ctx echo.Context, order *models.Order) error {
	return r.db.Model(&models.Order{}).Updates(&order).Error
}
