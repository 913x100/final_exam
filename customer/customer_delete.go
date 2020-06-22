package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Customer) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	// err := database.DeleteCustomersById(id)
	err := r.DB.DeleteCustomer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, "customer deleted.")
}
