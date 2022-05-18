package mysql

import (
	"afaf-group.com/domain/models"
	"afaf-group.com/domain/repository"
	"afaf-group.com/domain/request"
	"afaf-group.com/pkg/common"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type foodMySQLRepository struct {
	db *gorm.DB
}

func NewFoodMySQLRepository(db *gorm.DB) repository.FoodMySQLRepository {
	return &foodMySQLRepository{db: db}
}

func (r foodMySQLRepository) GetAll(ctx echo.Context, foodRequest *request.GetAllFoodRequest) (*common.Pagination, error) {
	var food []models.Food
	query := r.db.Model(&models.Food{}).Find(&food)
	if foodRequest.Search != "" {
		query = query.Where("name ILIKE ?", foodRequest.Search+"%")
	}

	paging, err := common.NewPagination().Paging(query, foodRequest.Page, foodRequest.Limit, []string{"id"}, &food)

	return paging, err
}
