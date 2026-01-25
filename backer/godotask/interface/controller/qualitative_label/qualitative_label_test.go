package qualitative_label

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
type MockQualitativeLabelRepository struct{}

func (m *MockQualitativeLabelRepository) Create(ql *model.QualitativeLabel) error {
  return nil
}

func (m *MockQualitativeLabelRepository) FindByID(id string) (*model.QualitativeLabel, error) {
  return &model.QualitativeLabel{
		ID:        "1",
		TaskID:    100,
		UserID:    1,
		Content:   "Test label content",
		Category:  "CategoryA",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
  }, nil
}

func (m *MockQualitativeLabelRepository) FindAll() ([]model.QualitativeLabel, error) {
  return []model.QualitativeLabel{
		{
			ID:        "1",
			TaskID:    100,
			UserID:    1,
			Content:   "Test label content",
			Category:  "CategoryA",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
  }, nil
}

func (m *MockQualitativeLabelRepository) Update(id string, ql *model.QualitativeLabel) error {
    return nil
}

func (m *MockQualitativeLabelRepository) Delete(id string) error {
    return nil
}

// テスト用ルーター
func setupRouter() *gin.Engine {
    r := gin.Default()
    mockRepo := &MockQualitativeLabelRepository{}
    mockService := &service.QualitativeLabelService{Repo: mockRepo}
    ctl := &QualitativeLabelController{Service: mockService}

    r.POST("/api/qualitative_label", ctl.AddQualitativeLabel)
    r.PUT("/api/qualitative_label/:id", ctl.EditQualitativeLabel)
    r.DELETE("/api/qualitative_label/:id", ctl.DeleteQualitativeLabel)
    return r
}

func TestAddQualitativeLabel(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    body := `{
      "task_id": 100,
      "user_id": 1,
      "content": "Test label content",
      "category": "CategoryA"
    }`

    req, _ := http.NewRequest(http.MethodPost, "/api/qualitative_label", bytes.NewBuffer([]byte(body)))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Qualitative label added")
}

func TestUpdateQualitativeLabel(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    body := `{
      "content": "Updated label content",
      "category": "CategoryB"
    }`

    req, _ := http.NewRequest(http.MethodPut, "/api/qualitative_label/1", bytes.NewBuffer([]byte(body)))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Qualitative label edited")
}

func TestDeleteQualitativeLabel(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    req, _ := http.NewRequest(http.MethodDelete, "/api/qualitative_label/1", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Qualitative label deleted")
}
