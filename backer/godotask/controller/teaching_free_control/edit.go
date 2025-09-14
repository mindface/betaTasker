package teaching_free_control

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

// EditTeachingFreeControl: PUT /api/teaching_free_control/:id
func (ctl *TeachingFreeControlController) EditTeachingFreeControl(c *gin.Context) {
	id := c.Param("id")
	var teachingFreeControl model.TeachingFreeControl
	if err := c.ShouldBindJSON(&teachingFreeControl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.UpdateTeachingFreeControl(id, &teachingFreeControl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit teaching free control"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teaching free control edited", "teaching_free_control": teachingFreeControl})
}
