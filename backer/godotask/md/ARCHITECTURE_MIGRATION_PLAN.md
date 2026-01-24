# アーキテクチャ移行計画書

## 現状分析

### 現在のアーキテクチャ状況

#### ✅ Auth機能（クリーンアーキテクチャ実装済み）

```
domain/
├── entity/
│   └── user.go                    # Userエンティティ
├── repository/
│   └── user_repository.go          # UserRepositoryインターフェース
└── service/
    ├── password_service.go         # PasswordServiceインターフェース
    └── token_service.go            # TokenServiceインターフェース

infrastructure/
├── repository/
│   └── user_repository_gorm.go    # UserRepository実装（domain依存）
└── security/
    ├── password_service.go         # PasswordService実装
    └── jwt_service.go              # TokenService実装

usecase/
├── auth_usecase.go                 # ビジネスロジック
└── register_usecase.go

interface/http/
├── controller/
│   └── auth_controller.go          # HTTPハンドラー
└── middleware/
    └── auth_middleware.go          # 認証ミドルウェア
```

#### ❌ その他の機能（旧構造）

```
controller/                         # 直接HTTPハンドラー
├── task/
├── memory/
├── assessment/
└── ...

service/                            # ビジネスロジック（model直接使用）
├── task_service.go
├── memory_service.go
└── ...

infrastructure/repository/         # 実装（model直接使用、domain層なし）
├── task_repository.go
├── memory_repository.go
└── interface.go                    # リポジトリインターフェース（model依存）

model/                              # GORMモデル（データベース層）
├── Task.go
├── Memory.go
└── ...
```

### 問題点

1. **アーキテクチャの不統一**
   - Authのみクリーンアーキテクチャ
   - 他の機能は旧構造（Controller → Service → Repository → Model）

2. **依存関係の逆転**
   - 旧構造: `infrastructure/repository/interface.go`が`model`に依存
   - 正しい構造: `domain/repository`が`domain/entity`に依存し、`infrastructure`が`domain`に依存

3. **責務の混在**
   - Controllerがビジネスロジックを含む場合がある
   - ServiceがHTTP層の詳細を知っている

4. **テストの困難さ**
   - インターフェースが`infrastructure`層にあるため、モック作成が困難
   - Domain層の独立性がない

## 移行戦略

### 段階的移行アプローチ

既存のAuthアーキテクチャをテンプレートとして、各機能を段階的に移行します。

### 移行順序（優先度順）

1. **Task** - 最も使用頻度が高く、他の機能の基盤
2. **Memory** - Taskと密接に関連
3. **Assessment** - Task評価機能
4. **Heuristics** - ML機能群
5. **その他の機能** - Book, KnowledgePattern, LanguageOptimization等

## 移行手順（Task機能を例に）

### Phase 1: Domain層の作成

#### 1.1 Entityの作成

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
```

#### 1.2 Repositoryインターフェースの作成

```go
// domain/repository/task_repository.go
package repository

import "github.com/godotask/domain/entity"

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

### Phase 2: Infrastructure層の移行

#### 2.1 Repository実装の移行

```go
// infrastructure/repository/task_repository_gorm.go
package repository

import (
	"github.com/godotask/domain/entity"
	domainRepo "github.com/godotask/domain/repository"
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
	// model.Taskからentity.Taskへの変換が必要
	modelTask := convertEntityToModel(task)
	return r.db.Create(modelTask).Error
}

func (r *TaskRepositoryGorm) FindByID(id int) (*entity.Task, error) {
	var modelTask model.Task
	if err := r.db.First(&modelTask, id).Error; err != nil {
		return nil, err
	}
	return convertModelToEntity(&modelTask), nil
}

// ... 他のメソッドも同様に実装

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

### Phase 3: Usecase層の作成

#### 3.1 Usecaseの実装

```go
// usecase/task_usecase.go
package usecase

import (
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
	// ビジネスロジック
	if task.Title == "" {
		return errors.New("title is required")
	}
	task.UserID = userID
	return u.taskRepo.Create(task)
}

func (u *TaskUsecase) GetTask(userID, taskID int) (*entity.Task, error) {
	task, err := u.taskRepo.FindByID(taskID)
	if err != nil {
		return nil, err
	}
	// 権限チェック
	if task.UserID != userID {
		return nil, errors.New("unauthorized")
	}
	return task, nil
}

func (u *TaskUsecase) ListTasks(userID int) ([]*entity.Task, error) {
	return u.taskRepo.FindAll(userID)
}

// ... 他のメソッド
```

### Phase 4: Interface層の移行

#### 4.1 Controllerの移行

```go
// interface/http/controller/task_controller.go
package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/godotask/usecase"
	"github.com/godotask/domain/entity"
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
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userID, _ := ctx.Get("user_id")
	task := &entity.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
	}

	if err := c.usecase.CreateTask(userID.(int), task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// ... 他のメソッド
```

### Phase 5: DI設定の更新

#### 5.1 server/init.goの更新

```go
// server/init.go
func Init() {
	// ... 既存のauth設定 ...

	// Task機能のDI設定
	taskRepo := repository.NewGormTaskRepository(model.DB)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskController = controller.NewTaskController(taskUsecase)

	// ... 他の機能も同様に ...
	
	router = setupRouter()
}
```

#### 5.2 router.goの更新

```go
// server/router.go
func setupRouter() *gin.Engine {
	r := gin.Default()

	// ... 既存の設定 ...

	protected := r.Group("/api")
	protected.Use(authMiddleware)
	{
		// Task API
		protected.POST("/task", taskController.AddTask)
		protected.GET("/task", taskController.ListTasks)
		protected.GET("/task/:id", taskController.GetTask)
		protected.PUT("/task/:id", taskController.EditTask)
		protected.DELETE("/task/:id", taskController.DeleteTask)
		
		// ... 他のエンドポイント ...
	}

	return r
}
```

## 移行チェックリスト

### Task機能移行

- [ ] Phase 1: Domain層の作成
  - [ ] `domain/entity/task.go` 作成
  - [ ] `domain/repository/task_repository.go` 作成
- [ ] Phase 2: Infrastructure層の移行
  - [ ] `infrastructure/repository/task_repository_gorm.go` 作成
  - [ ] Model ↔ Entity変換関数の実装
- [ ] Phase 3: Usecase層の作成
  - [ ] `usecase/task_usecase.go` 作成
  - [ ] ビジネスロジックの移行
- [ ] Phase 4: Interface層の移行
  - [ ] `interface/http/controller/task_controller.go` 作成
  - [ ] 旧`controller/task/`の削除
- [ ] Phase 5: DI設定の更新
  - [ ] `server/init.go`の更新
  - [ ] `server/router.go`の更新
- [ ] Phase 6: テストと検証
  - [ ] 単体テストの作成
  - [ ] 統合テストの実行
  - [ ] 旧コードの削除

### その他の機能

- [ ] Memory機能の移行
- [ ] Assessment機能の移行
- [ ] Heuristics機能の移行
- [ ] Book機能の移行
- [ ] その他の機能の移行

## 注意事項

### 1. 後方互換性の維持

- 移行中は旧エンドポイントと新エンドポイントを並行運用
- 段階的に旧エンドポイントを非推奨化

### 2. データ変換の考慮

- `model.Task`と`entity.Task`の変換が必要
- リレーション（Preload）の扱いに注意
- JSONフィールドの変換処理

### 3. エラーハンドリング

- Domain層でのエラー定義
- Usecase層でのビジネスロジックエラー
- Controller層でのHTTPエラー変換

### 4. テスト戦略

- Domain層: 純粋なGoテスト（モック不要）
- Infrastructure層: 統合テスト（テストDB使用）
- Usecase層: モックリポジトリを使用したテスト
- Controller層: HTTPテスト

## 推奨される移行タイムライン

1. **Week 1-2**: Task機能の移行（テンプレート作成）
2. **Week 3-4**: Memory機能の移行
3. **Week 5-6**: Assessment機能の移行
4. **Week 7-8**: Heuristics機能の移行
5. **Week 9-10**: その他の機能の移行
6. **Week 11-12**: テストとリファクタリング

## 参考: Auth機能の実装パターン

Auth機能が既にクリーンアーキテクチャで実装されているため、これをテンプレートとして使用できます：

- `domain/entity/user.go` → `domain/entity/task.go`
- `domain/repository/user_repository.go` → `domain/repository/task_repository.go`
- `infrastructure/repository/user_repository_gorm.go` → `infrastructure/repository/task_repository_gorm.go`
- `usecase/auth_usecase.go` → `usecase/task_usecase.go`
- `interface/http/controller/auth_controller.go` → `interface/http/controller/task_controller.go`

## 移行後のアーキテクチャ

```
domain/
├── entity/              # ビジネスエンティティ
├── repository/          # リポジトリインターフェース
└── service/            # ドメインサービスインターフェース

infrastructure/
├── repository/          # リポジトリ実装（domain依存）
└── security/           # セキュリティ実装

usecase/                 # ビジネスロジック（domain依存）

interface/http/
├── controller/          # HTTPハンドラー（usecase依存）
└── middleware/          # HTTPミドルウェア

model/                   # GORMモデル（infrastructure層でのみ使用）
```

## 次のステップ

1. Task機能の移行を開始
2. 移行パターンを確立
3. 他の機能に適用
4. 旧コードの削除
5. ドキュメントの更新