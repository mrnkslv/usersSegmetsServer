package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrnkslv/user-segmentation-service/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")

	segments := api.Group("/segments")
	{
		segments.POST("/", h.createSegment)
		segments.DELETE("/", h.deleteSegment)
	}
	users := api.Group("/users")
	{
		users.POST("/", h.addUserToSegment)
		users.GET("/", h.getActiveSegmentsByID)
	}
	return router
}
