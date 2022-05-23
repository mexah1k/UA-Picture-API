package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"ukraine-picture/src/common"
	"ukraine-picture/src/ports"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StoriesController struct {
	app ports.StoriesPort
}

func NewStoriesController(app ports.StoriesPort) *StoriesController {
	return &StoriesController{app: app}
}

func (api StoriesController) Create(c *gin.Context) {
	var request UpsertStoryRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	upsertStory := MapUpsertStoryRequest(&request)
	id, err := api.app.Create(&upsertStory)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (api StoriesController) Update(c *gin.Context) {
	inputId := c.Param("id")

	parsedId, err := strconv.Atoi(inputId)
	if err != nil || parsedId < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	var request UpsertStoryRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	upsertStory := MapUpsertStoryRequest(&request)
	id, err := api.app.Update(uint(parsedId), &upsertStory)

	if err != nil {
		if errors.Is(err, common.ItemNotFoundError) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (api StoriesController) Delete(c *gin.Context) {
	inputId := c.Param("id")

	parsedId, err := strconv.Atoi(inputId)
	if err != nil && parsedId > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	err = api.app.Delete(uint(parsedId))

	if err != nil {
		if errors.Is(err, common.ItemNotFoundError) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (api StoriesController) Get(c *gin.Context) {
	inputId := c.Param("id")

	parsedId, err := strconv.Atoi(inputId)
	if err != nil && parsedId > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	story, err := api.app.Find(uint(parsedId))

	if err != nil {
		if errors.Is(err, common.ItemNotFoundError) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, MapStoryResponse(story))
}

func (api StoriesController) Query(c *gin.Context) {
	inputId := c.DefaultQuery("id", "0")
	parsedId, err := strconv.Atoi(inputId)
	if err != nil && parsedId > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	stories, err := api.app.Query(uint(parsedId))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error: " + err.Error()})
		return
	}

	response := make([]StoryResponse, len(*stories))
	for id, story := range *stories {
		response[id] = MapStoryResponse(&story)
	}

	c.JSON(http.StatusOK, response)
}
