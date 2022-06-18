package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	authHTTPDelivery "afaf-group.com/pkg/auth/delivery/http"
	authRepository "afaf-group.com/pkg/auth/repository/mysql"
	authUseCase "afaf-group.com/pkg/auth/usecase"
	"afaf-group.com/pkg/common/utils"

	customerHTTPDelivery "afaf-group.com/pkg/customer/delivery/http"
	customerRepository "afaf-group.com/pkg/customer/repository/mysql"
	customerUseCase "afaf-group.com/pkg/customer/usecase"

	foodHTTPDelivery "afaf-group.com/pkg/food/delivery/http"
	foodRepository "afaf-group.com/pkg/food/repository/mysql"
	foodUseCase "afaf-group.com/pkg/food/usecase"
)

func InitAuthRoutes(r *echo.Echo, db *gorm.DB) {
	authRepo := authRepository.NewAuthMySQLRepository(db)
	authUCase := authUseCase.NewAuthUseCase(authRepo)
	authController := authHTTPDelivery.NewController(authUCase)

	customerRepo := customerRepository.NewCustomerMySQLRepository(db)
	customerUCase := customerUseCase.NewCustomerUseCase(customerRepo)
	customerController := customerHTTPDelivery.NewController(customerUCase)

	routerAuth := r.Group("/auth")
	routerAuth.POST("/login", authController.Login)
	routerAuth.POST("/register", authController.Register)

	routerCustomer := r.Group("/customers")
	routerCustomer.GET("", customerController.GetCustomerList, utils.AuthMidlleware())
	routerCustomer.POST("", customerController.CreateCustomer, utils.AuthMidlleware())
}

func InitFoodRoutes(r *echo.Echo, db *gorm.DB) {
	foodRepo := foodRepository.NewFoodMySQLRepository(db)
	foodUCase := foodUseCase.NewFoodUseCase(foodRepo)
	foodController := foodHTTPDelivery.NewController(foodUCase)

	foodGroup := r.Group("/foods")
	foodGroup.GET("", foodController.GetAll, utils.AuthMidlleware())
}
