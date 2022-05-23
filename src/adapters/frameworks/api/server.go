package api

import (
	"ukraine-picture/src/adapters/frameworks/api/controllers"
	"ukraine-picture/src/ports"

	"github.com/gin-gonic/gin"
)

type ApiAdapter struct {
	stories controllers.StoriesController
	media   controllers.MediaController
}

func NewAdapter(storiesController ports.StoriesPort, mediaController ports.MediaPort) *ApiAdapter {
	storiesApi := controllers.NewStoriesController(storiesController)
	mediaApi := controllers.NewMediaController(mediaController)

	return &ApiAdapter{stories: *storiesApi, media: *mediaApi}
}

func (api ApiAdapter) RegisterStoriesRouting(router *gin.RouterGroup) {
	router.POST("/", api.stories.Create)
	router.PUT("/:id", api.stories.Update)
	router.GET("/:id", api.stories.Get)
	router.POST("/query", api.stories.Query)
	router.DELETE("/:id", api.stories.Delete)
}

func (api ApiAdapter) RegisterMediaRouting(router *gin.RouterGroup) {
	router.POST("/", api.media.Create)
	router.PUT("/:id", api.media.Update)
	router.GET("/:id", api.media.Get)
	router.POST("/query", api.media.Query)
	router.DELETE("/:id", api.media.Delete)
}

func (api ApiAdapter) Run() {
	r := gin.Default()

	// routing
	v1 := r.Group("/api")

	// register entities
	api.RegisterStoriesRouting(v1.Group("/stories"))
	api.RegisterMediaRouting(v1.Group("/media"))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
