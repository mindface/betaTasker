package teaching_free_control

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListTeachingFreeControls: GET /api/teaching_free_control
func (ctl *TeachingFreeControlController) ListTeachingFreeControls(c *gin.Context) {
	teachingFreeControls, err := ctl.Service.ListTeachingFreeControls()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list TeachingFreeControls"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"TeachingFreeControls": teachingFreeControls})
}