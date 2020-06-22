package customer

import (
	"final_exam/database"

	"github.com/gin-gonic/gin"
)

type ICustomer interface {
	GetCustomers(c *gin.Context)
	GetCustomer(c *gin.Context)
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

type Customer struct {
	DB database.IDatabase
}

func NewCustomerHandler(db database.IDatabase) *Customer {
	return &Customer{
		DB: db,
	}
}
