package teaching_free_control

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetTeachingFreeControl: GET /api/teaching_free_control/:id
func (ctl *TeachingFreeControlController) GetTeachingFreeControl(c *gin.Context) {
	id := c.Param("id")
	teachingFreeControl, err := ctl.Service.GetTeachingFreeControlByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teaching free control not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"teaching_free_control": teachingFreeControl})
}
