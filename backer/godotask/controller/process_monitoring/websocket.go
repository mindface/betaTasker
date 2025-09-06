package process_monitoring

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleWebSocket: GET /api/process-monitoring/:id/ws
func (ctrl *ProcessMonitoringController) HandleWebSocket(c *gin.Context) {
	ctrl.Service.HandleWebSocket(c)
}

// UpdateAlertThresholds: POST /api/process-monitoring/:id/alert-thresholds
func (ctrl *ProcessMonitoringController) UpdateAlertThresholds(c *gin.Context) {
	monitorID := c.Param("id")
	if monitorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Monitor ID is required"})
		return
	}

	var req struct {
		Thresholds map[string]float64 `json:"thresholds" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update thresholds in active monitor
	activeMonitors := ctrl.Service.GetActiveMonitors()
	monitor, exists := activeMonitors[monitorID]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitor not found or not active"})
		return
	}

	// Update thresholds
	for key, value := range req.Thresholds {
		monitor.Thresholds[key] = value
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"message":    "Alert thresholds updated successfully",
		"thresholds": monitor.Thresholds,
	})
}