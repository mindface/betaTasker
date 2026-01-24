package memory

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/godotask/infrastructure/db/model"
    "github.com/godotask/usecase/service"
)

// モックリポジトリを作成
type MockMemoryRepository struct{}

func (m *MockMemoryRepository) Create(memory *model.Memory) error {
    return nil
}

func (m *MockMemoryRepository) FindByID(id string) (*model.Memory, error) {
    now := time.Now()
    return &model.Memory{
        ID:                1,
        UserID:            1,
        SourceType:        "book",
        Title:             "Test Title",
        Author:            "Test Author",
        Notes:             "Some notes",
        Tags:              "tag1,tag2",
        ReadStatus:        "read",
        ReadDate:          &now,
        Factor:            "Insightful",
        Process:           "Summarized",
        EvaluationAxis:    "Relevance",
        InformationAmount: "High",
        CreatedAt:         now,
        UpdatedAt:         now,
    }, nil
}

func (m *MockMemoryRepository) FindAll() ([]model.Memory, error) {
    memory, _ := m.FindByID("1")
    return []model.Memory{*memory}, nil
}

func (m *MockMemoryRepository) Update(id string, memory *model.Memory) error {
    return nil
}

func (m *MockMemoryRepository) Delete(id string) error {
    return nil
}

func setupRouter() *gin.Engine {
    r := gin.Default()
    mockRepo := &MockMemoryRepository{}
    mockService := &service.MemoryService{Repo: mockRepo}
    ctl := &MemoryController{Service: mockService}

    r.POST("/api/memory", ctl.AddMemory)
    r.GET("/api/memory", ctl.ListMemories)
    r.GET("/api/memory/:id", ctl.GetMemory)
    r.PUT("/api/memory/:id", ctl.EditMemory)
    r.DELETE("/api/memory/:id", ctl.DeleteMemory)
    return r
}

func TestAddMemory(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    body := `{
        "user_id": 1,
        "source_type": "book",
        "title": "Test Title",
        "author": "Test Author",
        "notes": "Some notes",
        "tags": "tag1,tag2",
        "read_status": "read",
        "factor": "Insightful",
        "process": "Summarized",
        "evaluation_axis": "Relevance",
        "information_amount": "High"
    }`

    req, err := http.NewRequest(http.MethodPost, "/api/memory", bytes.NewBuffer([]byte(body)))
    if err != nil {
        t.Fatalf("Couldn't create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Memory added")
}

func TestUpdateMemory(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    body := `{
        "user_id": 1,
        "source_type": "book",
        "title": "Updated Title",
        "author": "Updated Author",
        "notes": "Updated notes",
        "tags": "tag3,tag4",
        "read_status": "read",
        "factor": "Updated Insight",
        "process": "Analyzed",
        "evaluation_axis": "Importance",
        "information_amount": "Medium"
    }`

    req, err := http.NewRequest(http.MethodPut, "/api/memory/1", bytes.NewBuffer([]byte(body)))
    if err != nil {
        t.Fatalf("Couldn't create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Memory edited")
}

func TestDeleteMemory(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    req, err := http.NewRequest(http.MethodDelete, "/api/memory/1", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v", err)
    }

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Memory deleted")
}

func TestListMemory(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    req, err := http.NewRequest(http.MethodGet, "/api/memory", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v", err)
    }

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Test Title")
}

func TestGetMemory(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    req, err := http.NewRequest(http.MethodGet, "/api/memory/1", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v", err)
    }

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Test Title")
}
