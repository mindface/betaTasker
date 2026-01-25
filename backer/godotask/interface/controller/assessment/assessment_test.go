package assessment

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
type MockAssessmentRepository struct{}

func (m *MockAssessmentRepository) Create(assessment *model.Assessment) error {
    return nil
}

func (m *MockAssessmentRepository) FindByID(id string) (*model.Assessment, error) {
    return &model.Assessment{
        ID:                  1,
        TaskID:              1,
        UserID:              1,
        EffectivenessScore:  85,
        EffortScore:         70,
        ImpactScore:         90,
        QualitativeFeedback: "Test feedback",
        CreatedAt:           time.Now(),
        UpdatedAt:           time.Now(),
    }, nil
}

func (m *MockAssessmentRepository) FindAll() ([]model.Assessment, error) {
    return []model.Assessment{
        {
            ID:                  1,
            TaskID:              1,
            UserID:              1,
            EffectivenessScore:  85,
            EffortScore:         70,
            ImpactScore:         90,
            QualitativeFeedback: "Test feedback",
            CreatedAt:           time.Now(),
            UpdatedAt:           time.Now(),
        },
    }, nil
}

func (m *MockAssessmentRepository) Update(id string, assessment *model.Assessment) error {
    return nil // テスト用なので常に成功
}

func (m *MockAssessmentRepository) Delete(id string) error {
    return nil // テスト用なので常に成功
}

func setupRouter() *gin.Engine {
    r := gin.Default()
    
    // モックサービスとリポジトリを作成
    mockRepo := &MockAssessmentRepository{}
    mockService := &service.AssessmentService{Repo: mockRepo}
    ctl := &AssessmentController{Service: mockService}
    
    r.POST("/api/assessment", ctl.AddAssessment)
    r.GET("/api/assessment", ctl.ListAssessments)
    r.GET("/api/assessment/:id", ctl.GetAssessment)
    r.PUT("/api/assessment/:id", ctl.EditAssessment)
    r.DELETE("/api/assessment/:id", ctl.DeleteAssessment)
    return r
}

func TestAddAssessment(t *testing.T) {
    // Ginのテストモードを設定
    gin.SetMode(gin.TestMode)

    // テスト用のルーターをセットアップ
    r := setupRouter()

    body := `{
        "task_id": 1,
        "user_id": 1,
        "effectiveness_score": 85,
        "effort_score": 70,
        "impact_score": 90,
        "qualitative_feedback": "Test feedback"
    }`
    req, err := http.NewRequest(http.MethodPost, "/api/assessment", bytes.NewBuffer([]byte(body)))
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Assessment added")
}

func TestUpdateAssessment(t *testing.T) {
    // Ginのテストモードを設定
    gin.SetMode(gin.TestMode)

    // テスト用のルーターをセットアップ
    r := setupRouter()

    body := `{
        "task_id": 1,
        "user_id": 1,
        "effectiveness_score": 95,
        "effort_score": 80,
        "impact_score": 85,
        "qualitative_feedback": "Updated feedback"
    }`

    req, err := http.NewRequest(http.MethodPut, "/api/assessment/1", bytes.NewBuffer([]byte(body)))
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Assessment edited")
}

func TestDeleteAssessment(t *testing.T) {
    // Ginのテストモードを設定
    gin.SetMode(gin.TestMode)

    // テスト用のルーターをセットアップ
    r := setupRouter()

    req, err := http.NewRequest(http.MethodDelete, "/api/assessment/1", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Assessment deleted")
}

func TestListAssessments(t *testing.T) {
    // Ginのテストモードを設定
    gin.SetMode(gin.TestMode)

    // テスト用のルーターをセットアップ
    r := setupRouter()

    req, err := http.NewRequest(http.MethodGet, "/api/assessment", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Test feedback")
}

func TestGetAssessment(t *testing.T) {
    // Ginのテストモードを設定
    gin.SetMode(gin.TestMode)

    // テスト用のルーターをセットアップ
    r := setupRouter()

    req, err := http.NewRequest(http.MethodGet, "/api/assessment/1", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Test feedback")
}

func setupRouterWithMockService(repo repository.AssessmentRepository) *gin.Engine {
    gin.SetMode(gin.TestMode)
    r := gin.Default()

    service := service.NewAssessmentService(repo)
    controller := NewAssessmentController(service)

    r.GET("/api/assessment/:id", controller.GetAssessment)
    return r
}
