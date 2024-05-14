package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zakiyalmaya/assetfindr-assignment/application/post"
	"github.com/zakiyalmaya/assetfindr-assignment/infrastructure/repository"
	"github.com/zakiyalmaya/assetfindr-assignment/transport"
)

func main() {
	// init database
	db := repository.Database()

	repository := repository.NewRepository(db)

	// init service
	postService := post.NewPostService(repository)

	// init router
	r := gin.Default()

	transport.Handler(postService, r)

	r.Run(":8000")
}
