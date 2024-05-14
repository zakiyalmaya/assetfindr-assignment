package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/zakiyalmaya/assetfindr-assignment/application/post"
	"github.com/zakiyalmaya/assetfindr-assignment/transport/controller"
)

func Handler(postSvc post.PostService, r *gin.Engine) {
	c := controller.NewController(postSvc)

	r.POST("/post", c.PostController.Create)
}