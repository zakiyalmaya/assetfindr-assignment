package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/zakiyalmaya/assetfindr-assignment/application"
	"github.com/zakiyalmaya/assetfindr-assignment/transport/controller"
)

func Handler(service application.Service, r *gin.Engine) {
	c := controller.NewController(service)

	r.POST("/posts", c.Create)
	r.GET("/posts", c.GetAll)
	r.GET("/posts/:id", c.GetByID)
}