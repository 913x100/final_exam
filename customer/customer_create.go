package customer

import (
	"final_exam/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Customer) CreateCustomer(c *gin.Context) {
	cus := models.Customer{}

	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// id, err := database.CreateCustomer(cus.Name, cus.Email, cus.Status)
	id, err := r.DB.CreateCustomer(cus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	cus.ID = id
	c.JSON(http.StatusCreated, cus)
}
