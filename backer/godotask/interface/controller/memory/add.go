package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/errors"
	"github.com/rs/zerolog/log"
)

// AddMemory: POST /api/memory
func (ctl *MemoryController) AddMemory(c *gin.Context) {
	var memory model.Memory
	if err := c.ShouldBindJSON(&memory); err != nil {
		appErr := errors.NewAppError(
			errors.VAL_INVALID_INPUT,
			errors.GetErrorMessage(errors.VAL_INVALID_INPUT),
			err.Error(),
		)
		log.Error().Msgf("AddMemory BindJSON error: %v", err)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	if err := ctl.Service.CreateMemory(&memory); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error(),
		)
		log.Error().Msgf("AddMemory CreateMemory error: %v", err)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	log.Info().Msgf("AddMemory success: %+v", memory)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Memory added",
		"memory": memory,
	})
}