package task

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/service"
)

// DeleteTask: DELETE /api/task/:id
func (ctl *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
