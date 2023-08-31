package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrnkslv/user-segmentation-service/models"
)

func (h *Handler) addUserToSegment(c *gin.Context) {
	var input models.AddSegmentstoUser
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	exists, err := h.services.Users.Exists(input.UserId)
	if !exists {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newSegments, deletedSegments, err := h.services.Users.AddUserToSegments(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"new_segments": newSegments, "deleted_segments": deletedSegments})
}

func (h *Handler) getActiveSegmentsByID(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	userIdString := queryParams.Get("id")
	userId, err := strconv.Atoi(userIdString)
	userId64 := int64(userId)
	if err != nil {
		c.String(http.StatusBadRequest, "user id: %d is not of type", userId64)
	}
	exists, err := h.services.Users.Exists(userId64)
	if !exists {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	activeSegments, err := h.services.Users.GetActiveSegmentsByID(userId64)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, activeSegments)

}
