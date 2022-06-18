package http

import (
	"errors"
	"net/http"

	"afaf-group.com/domain"
	_ "afaf-group.com/domain/models"
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

// Login godoc
// @Tags         Authorization
// @Summary      Login
// @Description  Login with Email and Password
// @Accept       json
// @Param        loginRequest  body  request.AuthRequest  true  "Email: User Email; Password: User Password"
// @Produce      json
// @Success      200  {object}  response.APIResponse{data=models.Auth}
// @Failure      400  {object}  response.SwaggerHTTPErrorBadRequestValidation
// @Failure      401  {object}  response.SwaggerHTTPErrorUnauthorized
// @Failure      404  {object}  response.SwaggerHTTPErrorNotFound
// @Failure      500  {object}  response.SwaggerHTTPErrorInternalServerError
// @Security     ApiKeyAuth
// @Router       /auth/login [post]
func (c Controller) Login(ctx echo.Context) error {
	var loginRequest request.AuthRequest
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

// Register godoc
// @Tags         Authorization
// @Summary      Register
// @Description  Register with Email and Password
// @Accept       json
// @Param        registerReqest  body  request.AuthRequest  true  "Email: User Email; Password: User Password"
// @Produce      json
// @Success      200  {object}  response.APIResponse{}
// @Failure      400  {object}  response.SwaggerHTTPErrorBadRequestValidation
// @Failure      401  {object}  response.SwaggerHTTPErrorUnauthorized
// @Failure      404  {object}  response.SwaggerHTTPErrorNotFound
// @Failure      500  {object}  response.SwaggerHTTPErrorInternalServerError
// @Security     ApiKeyAuth
// @Router       /auth/register [post]
func (c Controller) Register(ctx echo.Context) error {
	var registerRequest request.AuthRequest
	if err := c.BindAndValidate(ctx, &registerRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	errRegister := c.authUseCase.Register(ctx, &registerRequest)
	if errRegister != nil {
		if errors.Is(errRegister, domain.ErrEmailAlreadyExists) {
			return ctx.JSON(http.StatusBadRequest, response.APIResponse{
				Message: errRegister.Error(),
			})
		}
		return ctx.JSON(http.StatusInternalServerError, errRegister)
	}
	return ctx.JSON(http.StatusOK, "succesfully register new user")
}
