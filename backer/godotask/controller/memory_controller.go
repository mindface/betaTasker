package memory

import (
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
	"github.com/godotask/model"
	"github.com/godotask/service"
)

// MemoryController ...
type MemoryController struct {
	Service *service.MemoryService
}

func (ctl *MemoryController) Create(c *gin.Context) {
	var memory model.Memory
	if err := c.ShouldBindJSON(&memory); err != nil {
		appErr := errors.NewAppError(
			errors.VAL_INVALID_INPUT,
			errors.GetErrorMessage(errors.VAL_INVALID_INPUT),
			err.Error(),
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	
	// 必須フィールドのバリデーション
	if memory.Title == "" {
		appErr := errors.NewAppError(
			errors.VAL_MISSING_FIELD,
			errors.GetErrorMessage(errors.VAL_MISSING_FIELD),
			"タイトルは必須項目です",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	
	if err := ctl.Service.CreateMemory(&memory); err != nil {
		var appErr *errors.AppError
		errMsg := err.Error()
		
		if strings.Contains(errMsg, "duplicate") || strings.Contains(errMsg, "UNIQUE constraint") {
			appErr = errors.NewAppError(
				errors.VAL_DUPLICATE_ENTRY,
				errors.GetErrorMessage(errors.VAL_DUPLICATE_ENTRY),
				"同じタイトルのメモリーが既に存在します",
			)
		} else if strings.Contains(errMsg, "connection refused") {
			appErr = errors.NewAppError(
				errors.DB_CONNECTION_FAILED,
				errors.GetErrorMessage(errors.DB_CONNECTION_FAILED),
				"",
			)
		} else {
			appErr = errors.NewAppError(
				errors.SYS_INTERNAL_ERROR,
				errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
				"",
			)
		}
		
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "メモリーが正常に作成されました",
		"data": gin.H{
			"memory": memory,
		},
	})
}

func (ctl *MemoryController) Get(c *gin.Context) {
	id := c.Param("id")
	memory, err := ctl.Service.GetMemoryByID(id)
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			"指定されたメモリーが見つかりません",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"memory": memory,
		},
	})
}

func (ctl *MemoryController) List(c *gin.Context) {
	memories, err := ctl.Service.ListMemories()
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			"メモリー一覧の取得に失敗しました",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"memories": memories,
		},
	})
}

func (ctl *MemoryController) Update(c *gin.Context) {
	id := c.Param("id")
	var memory model.Memory
	if err := c.ShouldBindJSON(&memory); err != nil {
		appErr := errors.NewAppError(
			errors.VAL_INVALID_INPUT,
			errors.GetErrorMessage(errors.VAL_INVALID_INPUT),
			err.Error(),
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	if err := ctl.Service.UpdateMemory(id, &memory); err != nil {
		var appErr *errors.AppError
		errMsg := err.Error()
		
		if strings.Contains(errMsg, "not found") {
			appErr = errors.NewAppError(
				errors.RES_NOT_FOUND,
				errors.GetErrorMessage(errors.RES_NOT_FOUND),
				"更新対象のメモリーが見つかりません",
			)
		} else {
			appErr = errors.NewAppError(
				errors.SYS_INTERNAL_ERROR,
				errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
				"",
			)
		}
		
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "メモリーが正常に更新されました",
		"data": gin.H{
			"memory": memory,
		},
	})
}

func (ctl *MemoryController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteMemory(id); err != nil {
		var appErr *errors.AppError
		errMsg := err.Error()
		
		if strings.Contains(errMsg, "not found") {
			appErr = errors.NewAppError(
				errors.RES_NOT_FOUND,
				errors.GetErrorMessage(errors.RES_NOT_FOUND),
				"削除対象のメモリーが見つかりません",
			)
		} else {
			appErr = errors.NewAppError(
				errors.SYS_INTERNAL_ERROR,
				errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
				"",
			)
		}
		
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "メモリーが正常に削除されました",
	})
}
