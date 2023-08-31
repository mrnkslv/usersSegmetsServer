package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrnkslv/user-segmentation-service/models"
)

func (h *Handler) addUserToSlug(c *gin.Context) {
	var input models.AddSlugstoUser
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	exists, err := h.services.Users.Exists(input.UserId)
	if exists == false {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newSlugs, err := h.services.Users.AddUserToSlug(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"new_slugs": newSlugs})
}

func (h *Handler) getActiveSlugsByID(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	userIdString := queryParams.Get("id")
	userId, err := strconv.Atoi(userIdString)
	userId64 := int64(userId)
	if err != nil {
		c.String(http.StatusBadRequest, "user id: %d is not of type", userId64)
	}

	if userId64 <= 0 {
		c.String(http.StatusBadRequest, "user id: %d is not valid", userId64)
	}
	activeSlugs, err := h.services.Users.GetActiveSlugsByID(userId64)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, activeSlugs)

}
