package handlers

import (
	"cmd/catapi/internal/models"
	"cmd/catapi/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateMission creates a new mission.
// @Summary Create a new mission
// @Description Create a new mission in the system
// @Tags Missions
// @Accept json
// @Produce json
// @Param mission body models.CreateMission true "Mission object that needs to be added"
// @Success 200 {object} models.Mission
// @Router /api/Missions/CreateMission [post]
func CreateMission(c *gin.Context) {
	var mission models.CreateMission
	if err := c.ShouldBindJSON(&mission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateMission(&mission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mission)
}

// RemoveMission removes a mission by ID.
// @Summary Remove a mission
// @Description Remove a mission from the system by ID
// @Tags Missions
// @Accept json
// @Produce json
// @Param id query int true "Mission ID to be deleted"
// @Success 200 {string} string "Mission deleted"
// @Router /api/Missions/RemoveMission [delete]
func RemoveMission(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := services.RemoveMission(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mission deleted"})
}

// UpdateMission updates an existing mission.
// @Summary Update a mission
// @Description Update an existing mission in the system
// @Tags Missions
// @Accept json
// @Produce json
// @Param mission body models.UpdateMission true "Updated mission object"
// @Success 200 {object} models.Mission
// @Router /api/Missions/UpdateMission [put]
func UpdateMission(c *gin.Context) {
	var mission models.UpdateMission
	if err := c.ShouldBindJSON(&mission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateMission(&mission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mission)
}

// UpdateMissionTargets updates targets for a mission.
// @Summary Update mission targets
// @Description Update targets for an existing mission in the system
// @Tags Missions
// @Accept json
// @Produce json
// @Param missionID query int true "Mission ID"
// @Param targets body []models.UpdateTarget true "Updated targets"
// @Success 200 {string} string "Mission targets updated"
// @Router /api/Missions/UpdateMissionTargets [put]
func UpdateMissionTargets(c *gin.Context) {
	missionID, err := strconv.Atoi(c.Query("missionID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mission ID"})
		return
	}

	var missionTargets []models.UpdateTarget
	if err := c.ShouldBindJSON(&missionTargets); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateMissionTargets(missionID, missionTargets); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mission targets updated"})
}

// AssignCat assigns a cat to a mission.
// @Summary Assign a cat to a mission
// @Description Assign a cat to an existing mission in the system
// @Tags Missions
// @Accept json
// @Produce json
// @Param missionID query int true "Mission ID"
// @Param catID query int true "Cat ID"
// @Success 200 {string} string "Cat assigned to mission"
// @Router /api/Missions/AssignCat [put]
func AssignCat(c *gin.Context) {
	missionID, err := strconv.Atoi(c.Query("missionID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mission ID"})
		return
	}

	catID, err := strconv.Atoi(c.Query("catID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cat ID"})
		return
	}

	if err := services.AssignCat(missionID, catID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cat assigned to mission"})
}

// GetMissions retrieves a list of all missions.
// @Summary Get missions
// @Description Retrieve a list of all missions in the system
// @Tags Missions
// @Accept json
// @Produce json
// @Success 200 {array} models.Mission
// @Router /api/Missions/GetMissions [get]
func GetMissions(c *gin.Context) {
	missions, err := services.GetMissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, missions)
}

// GetMission retrieves a mission by ID.
// @Summary Get mission
// @Description Retrieve a mission by ID
// @Tags Missions
// @Accept json
// @Produce json
// @Param id query int true "Mission ID"
// @Success 200 {object} models.Mission
// @Router /api/Missions/GetMission [get]
func GetMission(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	mission, err := services.GetMission(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mission)
}
