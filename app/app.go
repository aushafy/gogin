package app

import (
	"gogin/controllers"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	router.GET("/endpoint", controllers.GetLabels)
	router.GET("/weather", controllers.GetWeather)
	router.Run("0.0.0.0:8080")
}
