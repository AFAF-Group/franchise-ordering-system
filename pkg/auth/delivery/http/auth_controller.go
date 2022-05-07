package http

import (
	"net/http"

	"afaf-group.com/domain/request"
	"afaf-group.com/domain/response"
	"afaf-group.com/domain/usecase"
	"afaf-group.com/pkg/common"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	common.Controller
	authUseCase usecase.AuthUseCase
}

func NewController(authUseCase usecase.AuthUseCase) *Controller {
	return &Controller{authUseCase: authUseCase}
}

func (c Controller) Login(ctx echo.Context) error {
	var loginRequest request.LoginRequest
	if err := c.BindAndValidate(ctx, &loginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	user, err := c.authUseCase.Login(ctx, &loginRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.APIResponse{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.APIResponse{
		Message: "Anda berhasil Login",
		Data:    user,
	})
}
