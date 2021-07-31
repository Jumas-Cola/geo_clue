package main

import (
	"fmt"
	"geo_clue/models"
	"net/http"
	"strconv"

	"github.com/jumas-cola/geo_clue/models"

	"github.com/gin-gonic/gin"
)

type headers struct {
	ContentType string `header:"Content-Type"`
}

var err error

// func getCities(c *gin.Context) {
// 	str := c.Query("str")
// 	lim := c.Query("lim")
// 	country_id := c.Query("country_id")

// 	h := headers{}

// 	var cities []City

// 	if str == "" {
// 		c.JSON(200, cities)
// 		return
// 	}

// 	if lim == "" {
// 		lim = "10"
// 	} else {
// 		if l, err := strconv.Atoi(lim); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		} else if l > 100 {
// 			lim = "100"
// 		}
// 	}

// 	if err := c.ShouldBindHeader(&h); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	sql_str := fmt.Sprintf(`SELECT title_ru, area_ru, region_ru FROM _cities
// 			WHERE LOWER(title_ru) LIKE LOWER('%s%%')`, str)

// 	if country_id != "" {
// 		if _, err := strconv.Atoi(country_id); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		sql_str += fmt.Sprintf(` AND country_id=%s`, country_id)
// 	}
// 	sql_str += fmt.Sprintf(` LIMIT %s`, lim)
// 	fmt.Println(sql_str)

// 	rows, err := conn.Query(context.Background(), sql_str)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var city City
// 		err = rows.Scan(&city.TitleRU, &city.AreaRU, &city.RegionRU)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		cities = append(cities, city)
// 	}

// 	c.JSON(200, cities)
// }

func getCountries(c *gin.Context) {
	str := c.Query("str")
	lim := 10

	h := headers{}

	var countries []Country

	if str == "" {
		c.JSON(200, countries)
		return
	}

	ls := c.Query("lim")

	if ls != "" {
		if lim, err = strconv.Atoi(ls); err != nil {
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

	// if _, err := conn.Prepare("country-query", `SELECT country_id, title_ru FROM _countries WHERE LOWER(title_ru) LIKE LOWER('$1::string%%') LIMIT $2::string`); err != nil {

	db.Limit(lim).Where("LOWER(title_ru) LIKE LOWER(?)", fmt.Sprintf("%s%%", str)).Find(&countries)

	c.JSON(200, countries)
}

func getDocs(c *gin.Context) {
	h := headers{}

	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"/city": "str:string - string for search; lim:int - limit values(max 100); country_id:int - country number in table _countries", "/country": "str:string - string for search; lim:int - limit values(max 100)"})
}

func main() {

	models.ConnectDB()

	r := gin.Default()
	// r.GET("/city", getCities)
	r.GET("/country", getCountries)
	r.GET("/", getDocs)
	r.Run(":8000")
}
