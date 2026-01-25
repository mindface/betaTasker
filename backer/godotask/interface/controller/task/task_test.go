package task

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

// モックリポジトリ
type MockTaskRepository struct{}

func (m *MockTaskRepository) Create(task *model.Task) error {
	return nil
}

func (m *MockTaskRepository) FindByID(id string) (*model.Task, error) {
	now := time.Now()
	return &model.Task{
		ID:          1,
		UserID:      1,
		Title:       "Test Task Title",
		Description: "Test Description",
		Date:        &now,
		Status:      "todo",
		Priority:    3,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (m *MockTaskRepository) FindAll() ([]model.Task, error) {
	task, _ := m.FindByID("1")
	return []model.Task{*task}, nil
}

func (m *MockTaskRepository) Update(id string, task *model.Task) error {
	return nil
}

func (m *MockTaskRepository) Delete(id string) error {
	return nil
}

// Router + Controller setup
func setupRouter() *gin.Engine {
	r := gin.Default()
	mockRepo := &MockTaskRepository{}
	mockService := &service.TaskService{Repo: mockRepo}
	ctl := &TaskController{Service: mockService}

	r.POST("/api/task", ctl.AddTask)
	r.GET("/api/task", ctl.ListTasks)
	r.GET("/api/task/:id", ctl.GetTask)
	r.PUT("/api/task/:id", ctl.EditTask)
	r.DELETE("/api/task/:id", ctl.DeleteTask)
	return r
}

// ===== テスト =====

func TestAddTask(t *testing.T) {
	r := setupRouter()
	body := `{
		"user_id": 1,
		"title": "Test Task Title",
		"description": "Test Description",
		"status": "todo",
		"priority": 3
	}`

	req, _ := http.NewRequest(http.MethodPost, "/api/task", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Task added")
}

func TestUpdateTask(t *testing.T) {
	r := setupRouter()
	body := `{
		"user_id": 1,
		"title": "Updated Task Title",
		"description": "Updated Description",
		"status": "in_progress",
		"priority": 2
	}`

	req, _ := http.NewRequest(http.MethodPut, "/api/task/1", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Task edited")
}

func TestDeleteTask(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest(http.MethodDelete, "/api/task/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Task deleted")
}

func TestListTask(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest(http.MethodGet, "/api/task", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Task Title")
}

func TestGetTask(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest(http.MethodGet, "/api/task/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Task Title")
}
