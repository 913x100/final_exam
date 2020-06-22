package customer

import (
	"final_exam/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Customer) UpdateCustomer(c *gin.Context) {
	id := c.Param("id")

	// row, err := database.GetCustomersByID(id)
	row, err := r.DB.GetCustomer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	cus := models.Customer{}
	err = row.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = r.DB.UpdateCustomer(cus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, cus)
}
