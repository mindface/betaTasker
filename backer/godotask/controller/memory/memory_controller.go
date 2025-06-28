package memory

// import (
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// 	"github.com/godotask/model"
// 	"github.com/godotask/service"
// )

// MemoryControllerの定義はこのファイルにのみ残し、他ファイル（add.go, edit.go, delete.go, get.go, list.go）からは削除してください。
// type MemoryController struct {
// 	Service *service.MemoryService
// }

// // AddMemory: POST /api/memory
// func (ctl *MemoryController) AddMemory(c *gin.Context) {
// 	var memory model.Memory
// 	if err := c.ShouldBindJSON(&memory); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if err := ctl.Service.CreateMemory(&memory); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add memory"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Memory added", "memory": memory})
// }

// // EditMemory: PUT /api/memory/:id
// func (ctl *MemoryController) EditMemory(c *gin.Context) {
// 	id := c.Param("id")
// 	var memory model.Memory
// 	if err := c.ShouldBindJSON(&memory); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if err := ctl.Service.UpdateMemory(id, &memory); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit memory"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Memory edited", "memory": memory})
// }

// func (ctl *MemoryController) Get(c *gin.Context) {
// 	id := c.Param("id")
// 	memory, err := ctl.Service.GetMemoryByID(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Memory not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"memory": memory})
// }

// func (ctl *MemoryController) List(c *gin.Context) {
// 	memories, err := ctl.Service.ListMemories()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list memories"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"memories": memories})
// }

// // DeleteMemory: DELETE /api/memory/:id
// func (ctl *MemoryController) DeleteMemory(c *gin.Context) {
// 	id := c.Param("id")
// 	if err := ctl.Service.DeleteMemory(id); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete memory"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Memory deleted"})
// }
