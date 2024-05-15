package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/zakiyalmaya/assetfindr-assignment/application"
	"github.com/zakiyalmaya/assetfindr-assignment/transport/controller"
)

func Handler(service application.Service, r *gin.Engine) {
	c := controller.NewController(service)

	r.POST("/api/posts", c.Create)
	r.GET("/api/posts", c.GetAll)
	r.GET("/api/posts/:id", c.GetByID)
	r.PUT("/api/posts/:id", c.Update)
	r.DELETE("/api/posts/:id", c.Delete)
}