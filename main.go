package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jackc/pgx/v4/pgxpool"
)

type City struct {
	TitleRU  *string `json:"title_ru"`
	AreaRU   *string `json:"area_ru"`
	RegionRU *string `json:"region_ru"`
}

type Country struct {
	Id      *int64  `json:"country_id"`
	TitleRU *string `json:"title_ru"`
}

type headers struct {
	ContentType string `header:"Content-Type"`
}

var err error
var conn *pgxpool.Pool

func getCities(c *gin.Context) {
	str := c.Query("str")
	lim := c.Query("lim")
	country_id := c.Query("country_id")

	h := headers{}

	var cities []City

	if str == "" {
		c.JSON(200, cities)
		return
	}

	if lim == "" {
		lim = "10"
	} else {
		if l, err := strconv.Atoi(lim); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else if l > 100 {
			lim = "100"
		}
	}

	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sql_str := fmt.Sprintf(`SELECT title_ru, area_ru, region_ru FROM _cities 
			WHERE LOWER(title_ru) LIKE LOWER('%s%%') `, str)

	if country_id != "" {
		if _, err := strconv.Atoi(country_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		sql_str += fmt.Sprintf(` AND country_id=%s`, country_id)
	}
	sql_str += fmt.Sprintf(`LIMIT %s`, lim)
	fmt.Println(sql_str)

	rows, err := conn.Query(context.Background(), sql_str)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var city City
		err = rows.Scan(&city.TitleRU, &city.AreaRU, &city.RegionRU)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		cities = append(cities, city)
	}

	c.JSON(200, cities)
}

func getCountries(c *gin.Context) {
	str := c.Query("str")
	lim := c.Query("lim")

	h := headers{}

	var countries []Country

	if str == "" {
		c.JSON(200, countries)
		return
	}

	if lim == "" {
		lim = "10"
	} else {
		if l, err := strconv.Atoi(lim); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else if l > 100 {
			lim = "100"
		}
	}

	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sql_str := fmt.Sprintf(`SELECT country_id, title_ru FROM _countries 
			WHERE LOWER(title_ru) LIKE LOWER('%s%%') LIMIT %s`, str, lim)

	fmt.Println(sql_str)

	rows, err := conn.Query(context.Background(), sql_str)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var country Country
		err = rows.Scan(&country.Id, &country.TitleRU)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		countries = append(countries, country)
	}

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
	conn, err = pgxpool.Connect(context.Background(), "postgres://postgres:Wc9T6pwQ7@localhost:5432/geodata?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	r := gin.Default()
	r.GET("/city", getCities)
	r.GET("/country", getCountries)
	r.GET("/", getDocs)
	r.Run(":80")
}
