package handlers

import (
	"cmd/catapi/internal/models"
	"cmd/catapi/internal/services"
	"cmd/catapi/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new SpyCat
// @Description Create a new SpyCat in the system
// @Tags SpyCats
// @Accept json
// @Produce json
// @Param spyCat body models.CreateSpyCat true "SpyCat object that needs to be added"
// @Success 200 {object} models.SpyCat
// @Router /api/SpyCats/CreateSpyCat [post]
func CreateSpyCat(c *gin.Context) {
	var spyCat models.CreateSpyCat
	if err := c.ShouldBindJSON(&spyCat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateBreed(spyCat.Breed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateSpyCat(&spyCat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, spyCat)
}

// @Summary Remove a SpyCat
// @Description Remove a SpyCat from the system by ID
// @Tags SpyCats
// @Accept json
// @Produce json
// @Param id query integer true "SpyCat ID to be deleted"
// @Success 200 {string} string "SpyCat deleted"
// @Router /api/SpyCats/RemoveSpyCat [delete]
func RemoveSpyCat(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := services.RemoveSpyCat(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SpyCat deleted"})
}

// UpdateSpyCat updates the salary of a SpyCat by ID.
// @Summary Update SpyCat Salary
// @Description Update the salary of a SpyCat by ID
// @Tags SpyCats
// @Accept json
// @Produce json
// @Param req body models.UpdateSalaryRequest true "Request body containing ID and new Salary"
// @Success 200 {string} string "SpyCat salary updated"
// @Router /api/SpyCats/UpdateSpyCat [put]
func UpdateSpyCat(c *gin.Context) {
	var req models.UpdateSalaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateSpyCat(req.ID, req.Salary); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "SpyCat salary updated"})
}

// GetSpyCats retrieves a list of all SpyCats.
// @Summary Get SpyCats
// @Description Retrieve a list of all SpyCats in the system
// @Tags SpyCats
// @Accept json
// @Produce json
// @Success 200 {array} models.SpyCat
// @Router /api/SpyCats/GetSpyCats [get]
func GetSpyCats(c *gin.Context) {
	spyCats, err := services.GetSpyCats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, spyCats)
}

// GetSpyCat retrieves a single SpyCat by ID.
// @Summary Get SpyCat
// @Description Retrieve a single SpyCat by ID
// @Tags SpyCats
// @Accept json
// @Produce json
// @Param id query int true "SpyCat ID"
// @Success 200 {object} models.SpyCat
// @Router /api/SpyCats/GetSpyCat [get]
func GetSpyCat(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	spyCat, err := services.GetSpyCat(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, spyCat)
}
