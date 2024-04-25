package routers

import (
	"github.com/gin-gonic/gin"
	"myGPT/controller"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	{
		v1.POST("/request", controller.RequestHandler)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
