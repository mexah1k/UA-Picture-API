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

type MediaController struct {
	app ports.MediaPort
}

func NewMediaController(app ports.MediaPort) *MediaController {
	return &MediaController{app: app}
}

func (api MediaController) Create(c *gin.Context) {
	var request UpsertMediaRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	upsertMedia := MapUpsertMediaRequest(&request)
	id, err := api.app.Create(&upsertMedia)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (api MediaController) Update(c *gin.Context) {
	inputId := c.Param("id")

	parsedId, err := strconv.Atoi(inputId)
	if err != nil || parsedId < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	var request UpsertMediaRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	upsertMedia := MapUpsertMediaRequest(&request)
	id, err := api.app.Update(uint(parsedId), &upsertMedia)

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

func (api MediaController) Delete(c *gin.Context) {
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

func (api MediaController) Get(c *gin.Context) {
	inputId := c.Param("id")

	parsedId, err := strconv.Atoi(inputId)
	if err != nil && parsedId > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	media, err := api.app.Find(uint(parsedId))

	if err != nil {
		if errors.Is(err, common.ItemNotFoundError) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, MapMediaResponse(media))
}

func (api MediaController) Query(c *gin.Context) {
	inputId := c.DefaultQuery("id", "0")
	parsedId, err := strconv.Atoi(inputId)
	if err != nil && parsedId > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	media, err := api.app.Query(uint(parsedId))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error: " + err.Error()})
		return
	}

	response := make([]MediaResponse, len(*media))
	for id, media := range *media {
		response[id] = MapMediaResponse(&media)
	}

	c.JSON(http.StatusOK, response)
}
