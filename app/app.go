package app

import (
	"gogin/controllers"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	router.GET("/api/v1/labels", controllers.GetLabels)
	router.GET("/api/v1/albums", controllers.GetAlbums)
	router.Run("0.0.0.0:8080")
}
