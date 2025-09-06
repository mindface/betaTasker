package process_monitoring

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
)

// StartMonitoring: POST /api/process-monitoring/start
func (ctrl *ProcessMonitoringController) StartMonitoring(c *gin.Context) {
	var req model.ProcessMonitoringRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	monitoring, err := ctrl.Service.StartMonitoring(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   monitoring,
		"message": "Process monitoring started successfully",
	})
}

// StopMonitoring: POST /api/process-monitoring/:id/stop
func (ctrl *ProcessMonitoringController) StopMonitoring(c *gin.Context) {
	monitorID := c.Param("id")
	if monitorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Monitor ID is required"})
		return
	}

	err := ctrl.Service.StopMonitoring(monitorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Process monitoring stopped successfully",
	})
}