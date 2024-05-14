package post

import (
	"net/http"

	"github.com/zakiyalmaya/assetfindr-assignment/application/post"
	"github.com/zakiyalmaya/assetfindr-assignment/model"
	"github.com/zakiyalmaya/assetfindr-assignment/utils"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postSvc post.PostService
}

func NewPostController(postSvc post.PostService) *PostController {
	return &PostController{
		postSvc: postSvc,
	}
}

func (p *PostController) Create(c *gin.Context) {
	defer c.Request.Body.Close()

	postReq := &model.CreatePostRequest{}
	if err := c.ShouldBindJSON(postReq); err != nil {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse(err.Error()))
		return
	}

	if err := utils.Validator(postReq); err != nil {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse(err.Error()))
		return
	}

	if err := p.postSvc.Create(postReq); err != nil {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, model.HTTPSuccessResponse(nil))
}

