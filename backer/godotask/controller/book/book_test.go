package book

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/godotask/model"
    "github.com/godotask/service"
)

// モックリポジトリを作成
type MockBookRepository struct{}

func (m *MockBookRepository) Create(book *model.Book) error {
    return nil
}

func (m *MockBookRepository) FindByID(id string) (*model.Book, error) {
    return &model.Book{
        ID:      1,
        Title:   "Test Book",
        Name:    "Test Name",
        Text:    "Test Text",
        Disc:    "Test Description",
        ImgPath: "Test Image Path",
    }, nil
}

func (m *MockBookRepository) FindAll() ([]model.Book, error) {
    return []model.Book{
        {
            ID:      1,
            Title:   "Test Book",
            Name:    "Test Name",
            Text:    "Test Text",
            Disc:    "Test Description",
            ImgPath: "Test Image Path",
        },
    }, nil
}

func (m *MockBookRepository) Update(id string, book *model.Book) error {
    return nil // テスト用なので常に成功
}

func (m *MockBookRepository) Delete(id string) error {
    return nil // テスト用なので常に成功
}

func setupRouter() *gin.Engine {
    r := gin.Default()
    
    // モックサービスとリポジトリを作成
    mockRepo := &MockBookRepository{}
    mockService := &service.BookService{Repo: mockRepo}
    ctl := &BookController{Service: mockService}
    
    r.POST("/api/book", ctl.AddBook)
    return r
}

func TestAddBook(t *testing.T) {
    // Ginのテストモードを設定
    gin.SetMode(gin.TestMode)

    // テスト用のルーターをセットアップ
    r := setupRouter()

    body := `{
        "title": "Test Title",
        "name": "Test Name",
        "text": "Test Text",
        "disc": "Test Description",
        "imgPath": "Test Image Path"
    }`

    req, err := http.NewRequest(http.MethodPost, "/api/book", bytes.NewBuffer([]byte(body)))
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Book added")
}

func TestUpdateBook(t *testing.T) {
    // Ginのテストモードを設定
    gin.SetMode(gin.TestMode)

    // テスト用のルーターをセットアップ
    r := gin.Default()
    
    // モックサービスとリポジトリを作成
    mockRepo := &MockBookRepository{}
    mockService := &service.BookService{Repo: mockRepo}
    ctl := &BookController{Service: mockService}
    
    r.PUT("/api/updatebook/:id", ctl.EditBook)

    body := `{
        "title": "Updated Title",
        "name": "Updated Name",
        "text": "Updated Text",
        "disc": "Updated Description",
        "imgPath": "Updated Image Path"
    }`

    req, err := http.NewRequest(http.MethodPut, "/api/updatebook/1", bytes.NewBuffer([]byte(body)))
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Book edited")
}

func TestDeleteBook(t *testing.T) {
    // Ginのテストモードを設定
    gin.SetMode(gin.TestMode)

    // テスト用のルーターをセットアップ
    r := gin.Default()
    
    // モックサービスとリポジトリを作成
    mockRepo := &MockBookRepository{}
    mockService := &service.BookService{Repo: mockRepo}
    ctl := &BookController{Service: mockService}
    
    r.DELETE("/api/deletebook/:id", ctl.DeleteBook)

    req, err := http.NewRequest(http.MethodDelete, "/api/deletebook/1", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Book deleted")
}