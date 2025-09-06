package process_monitoring

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/godotask/service"
)

// GetMonitorStatus: GET /api/process-monitoring/:id/status
func (ctrl *ProcessMonitoringController) GetMonitorStatus(c *gin.Context) {
	monitorID := c.Param("id")
	if monitorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Monitor ID is required"})
		return
	}

	activeMonitors := ctrl.Service.GetActiveMonitors()
	monitor, exists := activeMonitors[monitorID]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitor not found or not active"})
		return
	}

	status := gin.H{
		"id":               monitor.ID,
		"status":           monitor.Status,
		"process_type":     monitor.ProcessType,
		"start_time":       monitor.StartTime,
		"last_update":      monitor.LastUpdate,
		"runtime_seconds":  monitor.LastUpdate.Sub(monitor.StartTime).Seconds(),
		"current_metrics":  monitor.Metrics,
		"anomaly_count":    len(monitor.Anomalies),
		"recent_anomalies": ctrl.getRecentAnomalies(monitor.Anomalies, 3),
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   status,
	})
}

// GetMonitoringSummary: GET /api/process-monitoring/:id/summary
func (ctrl *ProcessMonitoringController) GetMonitoringSummary(c *gin.Context) {
	monitorID := c.Param("id")
	if monitorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Monitor ID is required"})
		return
	}

	activeMonitors := ctrl.Service.GetActiveMonitors()
	monitor, exists := activeMonitors[monitorID]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitor not found or not active"})
		return
	}

	// Calculate summary statistics
	runtime := monitor.LastUpdate.Sub(monitor.StartTime)
	
	// Count anomalies by severity
	anomalyCounts := make(map[string]int)
	recentAnomalies := 0
	cutoff := monitor.LastUpdate.Add(-5 * time.Minute) // 5 minutes ago

	for _, anomaly := range monitor.Anomalies {
		anomalyCounts[anomaly.Severity]++
		if anomaly.Timestamp.After(cutoff) {
			recentAnomalies++
		}
	}

	summary := gin.H{
		"monitor_id":         monitor.ID,
		"process_type":       monitor.ProcessType,
		"status":            monitor.Status,
		"runtime": gin.H{
			"seconds": runtime.Seconds(),
			"minutes": runtime.Minutes(),
			"hours":   runtime.Hours(),
		},
		"performance_overview": gin.H{
			"current_metrics":     monitor.Metrics,
			"target_performance":  monitor.TargetPerformance,
		},
		"anomaly_summary": gin.H{
			"total_anomalies":    len(monitor.Anomalies),
			"recent_anomalies":   recentAnomalies,
			"by_severity":        anomalyCounts,
		},
		"health_indicators": gin.H{
			"overall_health": ctrl.calculateOverallHealth(monitor),
			"stability":      ctrl.calculateStability(monitor),
			"efficiency":     ctrl.calculateEfficiency(monitor),
		},
		"last_updated": monitor.LastUpdate,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   summary,
	})
}

// Helper functions for health calculations
func (ctrl *ProcessMonitoringController) calculateOverallHealth(monitor *service.ProcessMonitor) string {
	recentAnomalies := 0
	cutoff := monitor.LastUpdate.Add(-5 * time.Minute) // 5 minutes

	for _, anomaly := range monitor.Anomalies {
		if anomaly.Timestamp.After(cutoff) && anomaly.Severity == "high" {
			recentAnomalies++
		}
	}

	if recentAnomalies == 0 {
		return "Healthy"
	} else if recentAnomalies < 3 {
		return "Warning"
	} else {
		return "Critical"
	}
}

func (ctrl *ProcessMonitoringController) calculateStability(monitor *service.ProcessMonitor) float64 {
	if len(monitor.Anomalies) == 0 {
		return 1.0
	}

	recentAnomalies := 0
	cutoff := monitor.LastUpdate.Add(-10 * time.Minute) // 10 minutes

	for _, anomaly := range monitor.Anomalies {
		if anomaly.Timestamp.After(cutoff) {
			recentAnomalies++
		}
	}

	stability := 1.0 - (float64(recentAnomalies) * 0.1)
	if stability < 0 {
		stability = 0
	}

	return stability
}

func (ctrl *ProcessMonitoringController) calculateEfficiency(monitor *service.ProcessMonitor) float64 {
	// Simple efficiency calculation based on current metrics
	// In a real implementation, this would be more sophisticated
	
	if monitor.Metrics == nil {
		return 0.8 // Default efficiency
	}

	// Look for efficiency-related metrics
	if efficiency, exists := monitor.Metrics["efficiency"]; exists {
		if effFloat, ok := efficiency.(float64); ok {
			return effFloat
		}
	}

	// Calculate based on other metrics if efficiency not directly available
	efficiency := 0.8 // Base efficiency

	// Adjust based on anomalies
	recentAnomalies := 0
	cutoff := monitor.LastUpdate.Add(-5 * time.Minute) // 5 minutes

	for _, anomaly := range monitor.Anomalies {
		if anomaly.Timestamp.After(cutoff) {
			recentAnomalies++
		}
	}

	efficiency -= float64(recentAnomalies) * 0.05

	if efficiency < 0 {
		efficiency = 0
	}
	if efficiency > 1 {
		efficiency = 1
	}

	return efficiency
}