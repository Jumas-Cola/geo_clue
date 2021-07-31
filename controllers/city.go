package controllers

import (
	"fmt"
	"geo_clue/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCities godoc
// @Summary Show a city list
// @Produce  json
// @Param str query string true "search city by str"
// @Param country query int false "search city by country"
// @Param lim query string false "query limit"
// @Router /city [get]
func GetCities(c *gin.Context) {
	str := c.Query("str")
	lim := 10
	country_id := c.Query("country")

	h := headers{}

	var cities []models.City

	if str == "" {
		c.JSON(200, cities)
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

	if country_id != "" {
		models.DB.Limit(lim).Where("LOWER(title_ru) LIKE LOWER(?) AND country_id = ?",
			fmt.Sprintf("%s%%", str), country_id).Find(&cities)
	} else {
		models.DB.Limit(lim).Where("LOWER(title_ru) LIKE LOWER(?)",
			fmt.Sprintf("%s%%", str)).Find(&cities)
	}

	c.JSON(200, cities)
}
