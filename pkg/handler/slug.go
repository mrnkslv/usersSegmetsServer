package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrnkslv/user-segmentation-service/models"
)

func (h *Handler) createSlug(c *gin.Context) {
	var input models.Slug
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Slugger.CreateSlug(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) deleteSlug(c *gin.Context) {

}
