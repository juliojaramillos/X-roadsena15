package routing

import (
	"github.com/gin-gonic/gin"

	"unal.edu.co/rest-example/api/controller"
)

func UserRouting(r *gin.RouterGroup) {
	r.GET("/citizen/runt-registry", controller.GETUser)
	r.POST("/citizen/transit-registry/create", controller.POSTUser)
}
