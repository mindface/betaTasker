package task

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/infrastructure/db/model"
)

// EditTask: PUT /api/task/:id
func (ctl *TaskController) EditTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.UpdateTask(id, &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task edited", "task": task})
}
