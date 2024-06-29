package handlers

import (
	"cmd/catapi/internal/models"
	"cmd/catapi/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddTarget adds a new target to a mission.
// @Summary Add a new target
// @Description Add a new target to a mission in the system
// @Tags Targets
// @Accept json
// @Produce json
// @Param target body models.AddTarget true "Target object that needs to be added"
// @Success 200 {object} models.Target
// @Router /api/Targets/AddTarget [post]
func AddTarget(c *gin.Context) {
	var target models.AddTarget
	if err := c.ShouldBindJSON(&target); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.AddTarget(&target); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, target)
}

// RemoveTarget removes a target from a mission by ID.
// @Summary Remove a target
// @Description Remove a target from a mission in the system by ID
// @Tags Targets
// @Accept json
// @Produce json
// @Param id query int true "Target ID to be deleted"
// @Success 200 {string} string "Target deleted"
// @Router /api/Targets/RemoveTarget [delete]
func RemoveTarget(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := services.RemoveTarget(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Target deleted"})
}
