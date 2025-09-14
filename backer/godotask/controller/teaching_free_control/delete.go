package teaching_free_control

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteTeachingFreeControl: DELETE /api/teaching_free_control/:id
func (ctl *TeachingFreeControlController) DeleteTeachingFreeControl(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteTeachingFreeControl(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete teaching free control"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teaching free control deleted"})
}
