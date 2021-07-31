package controllers

import (
	"fmt"
	"geo_clue/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCountries godoc
// @Summary Show a country list
// @Produce  json
// @Param str query string true "search country by str"
// @Param lim query string false "query limit"
// @Router /country [get]
func GetCountries(c *gin.Context) {
	str := c.Query("str")
	lim := 10

	h := headers{}

	var countries []models.Country

	if str == "" {
		c.JSON(200, countries)
		return
	}

	lim_str := c.Query("lim")

	if lim_str != "" {
		if lim, err = strconv.Atoi(lim_str); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else if lim > 100 {
			lim = 100
		}
	}

	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Limit(lim).Where("LOWER(title_ru) LIKE LOWER(?)", fmt.Sprintf("%s%%", str)).Find(&countries)

	c.JSON(200, countries)
}
