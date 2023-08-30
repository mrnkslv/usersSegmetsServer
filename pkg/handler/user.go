package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrnkslv/user-segmentation-service/models"
)

func (h *Handler) addUserToSlug(c *gin.Context) {
	var input models.AddSlugstoUser
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newSlugs, err := h.services.Users.AddUserToSlug(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, newSlugs)
}

func (h *Handler) deleteUserFromSlug(c *gin.Context) {

}

func (h *Handler) getActiveSlugsByID(c *gin.Context) {

}
