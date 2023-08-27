package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")

	slugs := api.Group("/slugs")
	{
		slugs.POST("/", h.createSlug)
		slugs.DELETE("/:id", h.deleteSlug)
	}
	users := api.Group("/users")
	{
		users.POST("/:id", h.addUserToSlug)        //slugid
		users.DELETE("/:id", h.deleteUserFromSlug) //slugid
		users.GET("/:id", h.getActiveSlugsByID)    //userid
	}
	return router
}
