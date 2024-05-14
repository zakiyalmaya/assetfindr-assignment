package controller

import (
	"github.com/zakiyalmaya/assetfindr-assignment/application/post"
	postCtrl "github.com/zakiyalmaya/assetfindr-assignment/transport/controller/post"
)

type Controller struct {
	PostController *postCtrl.PostController
}

func NewController(postSvc post.PostService) *Controller {
	return &Controller{
		PostController: postCtrl.NewPostController(postSvc),
	}
}