package teaching_free_control

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

// AddTeachingFreeControl: POST /api/teaching_free_control
func (ctl *TeachingFreeControlController) AddTeachingFreeControl(c *gin.Context) {
	var teachingFreeControl model.TeachingFreeControl
	if err := c.ShouldBindJSON(&teachingFreeControl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.CreateTeachingFreeControl(&teachingFreeControl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add teaching free control"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teaching free control added", "teaching_free_control": teachingFreeControl})
}
