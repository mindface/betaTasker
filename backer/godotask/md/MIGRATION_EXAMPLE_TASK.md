# Task機能の移行実装例

このドキュメントは、Task機能をクリーンアーキテクチャに移行する際の具体的な実装例を示します。

## 移行前後の比較

### 移行前（旧構造）

```
controller/task/add.go
  ↓
service/task_service.go
  ↓
infrastructure/repository/task_repository.go (model.Task直接使用)
  ↓
model/Task.go (GORM)
```

### 移行後（クリーンアーキテクチャ）

```
interface/http/controller/task_controller.go
  ↓
usecase/task_usecase.go
  ↓
domain/repository/task_repository.go (interface)
  ↓
infrastructure/repository/task_repository_gorm.go (entity.Task使用)
  ↓
model/Task.go (変換層として使用)
```

## 実装例

### 1. Domain Entity

```go
// domain/entity/task.go
package entity

import "time"

type Task struct {
	ID          int
	UserID      int
	MemoryID    *int
	Title       string
	Description string
	Date        *time.Time
	Status      string
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// バリデーションメソッド
func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	if t.UserID == 0 {
		return errors.New("user_id is required")
	}
	return nil
}
```

### 2. Domain Repository Interface

```go
// domain/repository/task_repository.go
package repository

import (
	"github.com/godotask/domain/entity"
)

type TaskRepository interface {
	Create(task *entity.Task) error
	FindByID(id int) (*entity.Task, error)
	FindAll(userID int) ([]*entity.Task, error)
	ListTasksByUser(userID int) ([]*entity.Task, error)
	ListTasksByUserPager(userID int, offset, limit int) ([]*entity.Task, int64, error)
	Update(id int, task *entity.Task) error
	Delete(id int) error
}
```

### 3. Infrastructure Repository Implementation

```go
// infrastructure/repository/task_repository_gorm.go
package repository

import (
	"github.com/godotask/domain/entity"
	domainRepo "github.com/godotask/domain/repository"
	"github.com/godotask/infrastructure/db/model"
	"gorm.io/gorm"
)

type TaskRepositoryGorm struct {
	db *gorm.DB
}

// compile-time check
var _ domainRepo.TaskRepository = (*TaskRepositoryGorm)(nil)

func NewGormTaskRepository(db *gorm.DB) domainRepo.TaskRepository {
	return &TaskRepositoryGorm{db: db}
}

func (r *TaskRepositoryGorm) Create(task *entity.Task) error {
	modelTask := convertEntityToModel(task)
	if err := r.db.Create(modelTask).Error; err != nil {
		return err
	}
	// IDをエンティティに反映
	task.ID = modelTask.ID
	task.CreatedAt = modelTask.CreatedAt
	task.UpdatedAt = modelTask.UpdatedAt
	return nil
}

func (r *TaskRepositoryGorm) FindByID(id int) (*entity.Task, error) {
	var modelTask model.Task
	if err := r.db.First(&modelTask, id).Error; err != nil {
		return nil, err
	}
	return convertModelToEntity(&modelTask), nil
}

func (r *TaskRepositoryGorm) FindAll(userID int) ([]*entity.Task, error) {
	var modelTasks []model.Task
	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at DESC, id DESC").
		Find(&modelTasks).Error
	
	if err != nil {
		return nil, err
	}
	
	tasks := make([]*entity.Task, len(modelTasks))
	for i, mt := range modelTasks {
		tasks[i] = convertModelToEntity(&mt)
	}
	return tasks, nil
}

func (r *TaskRepositoryGorm) ListTasksByUser(userID int) ([]*entity.Task, error) {
	return r.FindAll(userID)
}

func (r *TaskRepositoryGorm) ListTasksByUserPager(userID int, offset, limit int) ([]*entity.Task, int64, error) {
	var modelTasks []model.Task
	var total int64
	
	q := r.db.Model(&model.Task{}).Where("user_id = ?", userID)
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	if err := q.Order("created_at DESC, id DESC").
		Limit(limit).
		Offset(offset).
		Find(&modelTasks).Error; err != nil {
		return nil, 0, err
	}
	
	tasks := make([]*entity.Task, len(modelTasks))
	for i, mt := range modelTasks {
		tasks[i] = convertModelToEntity(&mt)
	}
	return tasks, total, nil
}

func (r *TaskRepositoryGorm) Update(id int, task *entity.Task) error {
	modelTask := convertEntityToModel(task)
	return r.db.Model(&model.Task{}).
		Where("id = ?", id).
		Updates(modelTask).Error
}

func (r *TaskRepositoryGorm) Delete(id int) error {
	return r.db.Delete(&model.Task{}, id).Error
}

// 変換ヘルパー関数
func convertModelToEntity(m *model.Task) *entity.Task {
	return &entity.Task{
		ID:          m.ID,
		UserID:      m.UserID,
		MemoryID:    m.MemoryID,
		Title:       m.Title,
		Description: m.Description,
		Date:        m.Date,
		Status:      m.Status,
		Priority:    m.Priority,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func convertEntityToModel(e *entity.Task) *model.Task {
	return &model.Task{
		ID:          e.ID,
		UserID:      e.UserID,
		MemoryID:    e.MemoryID,
		Title:       e.Title,
		Description: e.Description,
		Date:        e.Date,
		Status:      e.Status,
		Priority:    e.Priority,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
```

### 4. Usecase

```go
// usecase/task_usecase.go
package usecase

import (
	"errors"
	"github.com/godotask/domain/entity"
	domainRepo "github.com/godotask/domain/repository"
)

type TaskUsecase struct {
	taskRepo domainRepo.TaskRepository
}

func NewTaskUsecase(taskRepo domainRepo.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		taskRepo: taskRepo,
	}
}

func (u *TaskUsecase) CreateTask(userID int, task *entity.Task) error {
	// バリデーション
	if err := task.Validate(); err != nil {
		return err
	}
	
	// ユーザーIDの設定
	task.UserID = userID
	
	// デフォルト値の設定
	if task.Status == "" {
		task.Status = "pending"
	}
	if task.Priority == 0 {
		task.Priority = 1
	}
	
	return u.taskRepo.Create(task)
}

func (u *TaskUsecase) GetTask(userID, taskID int) (*entity.Task, error) {
	task, err := u.taskRepo.FindByID(taskID)
	if err != nil {
		return nil, err
	}
	
	// 権限チェック
	if task.UserID != userID {
		return nil, errors.New("unauthorized: task does not belong to user")
	}
	
	return task, nil
}

func (u *TaskUsecase) ListTasks(userID int) ([]*entity.Task, error) {
	return u.taskRepo.FindAll(userID)
}

func (u *TaskUsecase) ListTasksPager(userID int, page, perPage int) ([]*entity.Task, int64, error) {
	offset := (page - 1) * perPage
	return u.taskRepo.ListTasksByUserPager(userID, offset, perPage)
}

func (u *TaskUsecase) UpdateTask(userID, taskID int, task *entity.Task) error {
	// 既存タスクの取得と権限チェック
	existing, err := u.GetTask(userID, taskID)
	if err != nil {
		return err
	}
	
	// 更新可能なフィールドのみ更新
	existing.Title = task.Title
	existing.Description = task.Description
	existing.Status = task.Status
	existing.Priority = task.Priority
	if task.Date != nil {
		existing.Date = task.Date
	}
	if task.MemoryID != nil {
		existing.MemoryID = task.MemoryID
	}
	
	return u.taskRepo.Update(taskID, existing)
}

func (u *TaskUsecase) DeleteTask(userID, taskID int) error {
	// 権限チェック
	_, err := u.GetTask(userID, taskID)
	if err != nil {
		return err
	}
	
	return u.taskRepo.Delete(taskID)
}
```

### 5. Controller

```go
// interface/http/controller/task_controller.go
package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/godotask/usecase"
	"github.com/godotask/domain/entity"
	"github.com/godotask/errors"
)

type TaskController struct {
	usecase *usecase.TaskUsecase
}

func NewTaskController(u *usecase.TaskUsecase) *TaskController {
	return &TaskController{usecase: u}
}

func (c *TaskController) AddTask(ctx *gin.Context) {
	var req struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		Status      string `json:"status"`
		Priority    int    `json:"priority"`
		MemoryID    *int   `json:"memory_id"`
		Date        string `json:"date"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		appErr := errors.NewAppError(
			errors.VAL_INVALID_INPUT,
			errors.GetErrorMessage(errors.VAL_INVALID_INPUT),
			err.Error(),
		)
		ctx.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	userID, _ := ctx.Get("user_id")
	
	task := &entity.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		MemoryID:    req.MemoryID,
	}

	if err := c.usecase.CreateTask(userID.(int), task); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error(),
		)
		ctx.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Task added",
		"task":    task,
	})
}

func (c *TaskController) ListTasks(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	
	tasks, err := c.usecase.ListTasks(userID.(int))
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error(),
		)
		ctx.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (c *TaskController) GetTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appErr := errors.NewAppError(
			errors.VAL_INVALID_INPUT,
			errors.GetErrorMessage(errors.VAL_INVALID_INPUT),
			"invalid task id",
		)
		ctx.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	userID, _ := ctx.Get("user_id")
	
	task, err := c.usecase.GetTask(userID.(int), id)
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error(),
		)
		ctx.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task": task})
}

func (c *TaskController) EditTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appErr := errors.NewAppError(
			errors.VAL_INVALID_INPUT,
			errors.GetErrorMessage(errors.VAL_INVALID_INPUT),
			"invalid task id",
		)
		ctx.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		Priority    int    `json:"priority"`
		MemoryID    *int   `json:"memory_id"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		appErr := errors.NewAppError(
			errors.VAL_INVALID_INPUT,
			errors.GetErrorMessage(errors.VAL_INVALID_INPUT),
			err.Error(),
		)
		ctx.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	userID, _ := ctx.Get("user_id")
	
	task := &entity.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		MemoryID:    req.MemoryID,
	}

	if err := c.usecase.UpdateTask(userID.(int), id, task); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error(),
		)
		ctx.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Task updated",
	})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appErr := errors.NewAppError(
			errors.VAL_INVALID_INPUT,
			errors.GetErrorMessage(errors.VAL_INVALID_INPUT),
			"invalid task id",
		)
		ctx.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	userID, _ := ctx.Get("user_id")
	
	if err := c.usecase.DeleteTask(userID.(int), id); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error(),
		)
		ctx.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Task deleted",
	})
}
```

### 6. DI設定の更新

```go
// server/init.go
package server

import (
	"time"
	"github.com/godotask/interface/http/controller"
	"github.com/godotask/usecase"
	"github.com/godotask/infrastructure/db/repository"
	"github.com/godotask/infrastructure/security"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/interface/http/middleware"
)

var (
	authController  *controller.AuthController
	taskController  *controller.TaskController
	authMiddleware  gin.HandlerFunc
	router          *gin.Engine
)

func Init() {
	// Auth機能のDI設定
	userRepo := repository.NewGormUserRepository(model.DB)
	passwordSvc := security.NewBcryptPasswordService()
	tokenSvc := security.NewJWTService(
		[]byte("secret"),
		24*time.Hour,
	)

	authUsecase := usecase.NewAuthUsecase(
		userRepo,
		passwordSvc,
		tokenSvc,
	)
	authController = controller.NewAuthController(authUsecase)
	authMiddleware = middleware.AuthMiddleware(tokenSvc)

	// Task機能のDI設定
	taskRepo := repository.NewGormTaskRepository(model.DB)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskController = controller.NewTaskController(taskUsecase)

	router = setupRouter()
}
```

```go
// server/router.go
func setupRouter() *gin.Engine {
	r := gin.Default()

	// ... 既存の設定 ...

	protected := r.Group("/api")
	protected.Use(authMiddleware)
	{
		// Auth endpoints
		public.POST("/login", authController.Login)
		public.POST("/register", authController.Register)
		public.POST("/logout", authController.Logout)

		// Task API
		protected.POST("/task", taskController.AddTask)
		protected.GET("/task", taskController.ListTasks)
		protected.GET("/task/:id", taskController.GetTask)
		protected.PUT("/task/:id", taskController.EditTask)
		protected.DELETE("/task/:id", taskController.DeleteTask)
	}

	return r
}
```

## 移行手順

1. **新規ファイルの作成**
   - `domain/entity/task.go`
   - `domain/repository/task_repository.go`
   - `infrastructure/repository/task_repository_gorm.go`
   - `usecase/task_usecase.go`
   - `interface/http/controller/task_controller.go`

2. **既存コードの確認**
   - 旧`controller/task/`の機能を確認
   - 旧`service/task_service.go`のビジネスロジックを確認
   - 旧`infrastructure/repository/task_repository.go`の実装を確認

3. **段階的な移行**
   - まず新構造で動作確認
   - 旧構造と並行運用
   - テスト完了後に旧構造を削除

4. **テスト**
   - 単体テスト（Usecase層）
   - 統合テスト（Repository層）
   - E2Eテスト（Controller層）

## 注意点

1. **リレーションの扱い**
   - Preloadが必要な場合は、Repository層で実装
   - Entityにリレーション情報を含める場合は、別途Entity定義が必要

2. **エラーハンドリング**
   - Domain層で定義したエラーをUsecase層で処理
   - Controller層でHTTPステータスコードに変換

3. **バリデーション**
   - Entity層で基本的なバリデーション
   - Usecase層でビジネスルールのバリデーション
   - Controller層でHTTPリクエストのバリデーション