package http

import (
	"net/http"

	"afaf-group.com/domain/models"
	"afaf-group.com/domain/request"
	"afaf-group.com/domain/response"
	"afaf-group.com/domain/usecase"
	"afaf-group.com/pkg/common"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	common.Controller
	foodUseCase usecase.FoodUseCase
}

func NewController(foodUseCase usecase.FoodUseCase) *Controller {
	return &Controller{foodUseCase: foodUseCase}
}

// GetAll godoc
// @Tags         Food
// @Summary      Get All Food With Pagination
// @Description  Get All Food With Pagination
// @Accept       json
// @Param        foodRequest  body  request.GetAllFoodRequest  true  "Page: Page Number; Limit: Request Limit; Search: Reqeust for search food"
// @Produce      json
// @Success      200  {object}  response.APIResponse{[]models.Food}
// @Failure      400  {object}  response.SwaggerHTTPErrorBadRequestValidation
// @Failure      401  {object}  response.SwaggerHTTPErrorUnauthorized
// @Failure      404  {object}  response.SwaggerHTTPErrorNotFound
// @Failure      500  {object}  response.SwaggerHTTPErrorInternalServerError
// @Security     ApiKeyAuth
// @Router       /foods [get]
func (c Controller) GetAll(ctx echo.Context) error {
	var foodRequest request.GetAllFoodRequest
	if err := c.BindAndValidate(ctx, &foodRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	if foodRequest.Limit == 0 {
		foodRequest.Limit = -1
	}

	paging, err := c.foodUseCase.GetAll(ctx, &foodRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.APIResponse{
			Message: err.Error(),
		})
	}
	foods := paging.Records.(*[]models.Food)
	return ctx.JSON(http.StatusOK, &response.APIResponse{
		Code:     http.StatusOK,
		Message:  http.StatusText(http.StatusOK),
		Data:     foods,
		PageInfo: response.NewPageInfo().ToPageInfo(paging),
	})
}
