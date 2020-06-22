package routers

import (
	"final_exam/customer"
	"final_exam/database"
	"final_exam/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	db := database.New()
	customerHandler := customer.NewCustomerHandler(db)
	r := gin.Default()

	api_v1 := r.Group("/v1")

	api_v1.Use(middleware.BasicAuth)

	customer := api_v1.Group("customer")
	customer.POST("/", customerHandler.CreateCustomer)
	customer.GET("/", customerHandler.GetCustomers)
	customer.GET("/:id", customerHandler.GetCustomer)
	customer.PUT("/:id", customerHandler.UpdateCustomer)
	customer.DELETE("/:id", customerHandler.DeleteCustomer)

	return r
}
