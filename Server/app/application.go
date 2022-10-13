package app

import (
	"github.com/gin-gonic/gin"

	_ "shirt/Server/routes"
)

var (
	router = gin.Default()
)

func StartApplication() {
	routes.mapUrls()
	router.Run(":8081")
}
