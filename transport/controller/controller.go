package controller

import (
	"net/http"
	"strconv"

	"github.com/zakiyalmaya/assetfindr-assignment/application"
	"github.com/zakiyalmaya/assetfindr-assignment/model"
	"github.com/zakiyalmaya/assetfindr-assignment/utils"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service application.Service
}

func NewController(svc application.Service) *Controller {
	return &Controller{
		service: svc,
	}
}

func (p *Controller) Create(c *gin.Context) {
	defer c.Request.Body.Close()

	postReq := &model.PostRequest{}
	if err := c.ShouldBindJSON(postReq); err != nil {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse(err.Error()))
		return
	}

	if err := utils.Validator(postReq); err != nil {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse(err.Error()))
		return
	}

	if err := p.service.Create(postReq); err != nil {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, model.HTTPSuccessResponse(nil))
}

func (p *Controller) GetAll(c *gin.Context) {
	posts, err := p.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.HTTPSuccessResponse(posts))
}

func (p *Controller) GetByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse("invalid id"))
		return
	}

	post, err := p.service.GetByID(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.HTTPSuccessResponse(post))
}

func (p *Controller) Update(c *gin.Context) {
	defer c.Request.Body.Close()

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse("invalid id"))
		return
	}

	postReq := &model.PostRequest{}
	if err := c.ShouldBindJSON(postReq); err != nil {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse(err.Error()))
		return
	}

	if err := utils.Validator(postReq); err != nil {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse(err.Error()))
		return
	}

	if err := p.service.Update(idInt, postReq); err != nil {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.HTTPSuccessResponse(nil))
}

func (p *Controller) Delete(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse("invalid id"))
		return
	}

	if err := p.service.Delete(idInt); err != nil {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.HTTPSuccessResponse(nil))
}