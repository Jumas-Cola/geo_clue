package main

import (
	"geo_clue/controllers"
	"geo_clue/models"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "geo_clue/docs"
)

// @title Swagger Geo API
// @version 1.0
// @description This is a api server for geo clue.
// @schemes http
// @host 188.225.72.165
// @BasePath /
func main() {

	models.ConnectDB()

	r := gin.Default()

	r.GET("/city", controllers.GetCities)
	r.GET("/country", controllers.GetCountries)

	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":80")

}
