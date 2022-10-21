package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	router.GET("/api/v1/labels", controllers.getLabels)
	router.GET("/api/v1/albums", controllers.getAlbums)
	router.Run("0.0.0.0:8080")
}
