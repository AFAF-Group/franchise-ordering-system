package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	authHTTPDelivery "afaf-group.com/pkg/auth/delivery/http"
	authRepository "afaf-group.com/pkg/auth/repository/mysql"
	authUseCase "afaf-group.com/pkg/auth/usecase"
)

func InitAuthRoutes(r *echo.Echo, db *gorm.DB) {
	authRepo := authRepository.NewAuthMySQLRepository(db)
	authUCase := authUseCase.NewAuthUseCase(authRepo)
	authController := authHTTPDelivery.NewController(authUCase)

	router := r.Group("/auth")
	router.POST("/login", authController.Login)
}
