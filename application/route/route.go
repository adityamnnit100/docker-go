package route

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.GET("/liveliness", GetHealth)
	r.GET("/visits", GetServerVisit)
	r.Run("0.0.0.0:8080")
}
