package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	authHTTPDelivery "afaf-group.com/pkg/auth/delivery/http"
	authRepository "afaf-group.com/pkg/auth/repository/mysql"
	authUseCase "afaf-group.com/pkg/auth/usecase"

	customerHTTPDelivery "afaf-group.com/pkg/customer/delivery/http"
	customerRepository "afaf-group.com/pkg/customer/repository/mysql"
	customerUseCase "afaf-group.com/pkg/customer/usecase"
)

func InitAuthRoutes(r *echo.Echo, db *gorm.DB) {
	authRepo := authRepository.NewAuthMySQLRepository(db)
	authUCase := authUseCase.NewAuthUseCase(authRepo)
	authController := authHTTPDelivery.NewController(authUCase)

	customerRepo := customerRepository.NewCustomerMySQLRepository(db)
	customerUCase := customerUseCase.NewCustomerUseCase(customerRepo)
	customerController := customerHTTPDelivery.NewController(customerUCase)

	router := r.Group("/auth")
	router.POST("/login", authController.Login)

	routerCustomer := r.Group("/customers")
	routerCustomer.GET("", customerController.CreateCustomer)
	routerCustomer.POST("", customerController.CreateCustomer)
}
