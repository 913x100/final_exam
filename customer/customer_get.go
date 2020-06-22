package customer

import (
	"final_exam/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Customer) GetCustomers(c *gin.Context) {
	rows, err := r.DB.GetCustomers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	custs := []models.Customer{}
	for rows.Next() {
		cus := models.Customer{}

		err := rows.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		custs = append(custs, cus)
	}

	c.JSON(http.StatusOK, custs)
}

func (r *Customer) GetCustomer(c *gin.Context) {
	id := c.Param("id")

	row, err := r.DB.GetCustomer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	cus := &models.Customer{}
	err = row.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, cus)
}
