package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zakiyalmaya/assetfindr-assignment/application"
	"github.com/zakiyalmaya/assetfindr-assignment/infrastructure/repository"
	"github.com/zakiyalmaya/assetfindr-assignment/transport"
)

func main() {
	// init database
	db := repository.Database()

	repository := repository.NewRepository(db)

	// init service
	service := application.NewService(repository)

	// init router
	r := gin.Default()

	transport.Handler(service, r)

	r.Run(":8000")
}
