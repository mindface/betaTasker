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


// 修正: userID パラメータを受け取り、userID=0 で全データ、それ以外でフィルタ
func (m *MockBookRepository) FindAll(userID uint) ([]model.Book, error) {
    if userID == 0 {
        // 全データを返す
        return []model.Book{
            {
                ID:      1,
                Title:   "Test Book 1",
                Name:    "Test Name 1",
                Text:    "Test Text 1",
                Disc:    "Test Description 1",
                ImgPath: "Test Image Path 1",
            },
            {
                ID:      2,
                Title:   "Test Book 2",
                Name:    "Test Name 2",
                Text:    "Test Text 2",
                Disc:    "Test Description 2",
                ImgPath: "Test Image Path 2",
            },
        }, nil
    }
    // userID でフィルタ
    if userID == 1 {
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
    return []model.Book{}, nil
}

// 追加: ページネーション対応
func (m *MockBookRepository) FindAllPager(userID uint, offset, limit int) ([]model.Book, int64, error) {
    books, _ := m.FindAll(userID)
    total := int64(len(books))
    
    // オフセット・リミットを適用
    if offset >= len(books) {
        return []model.Book{}, total, nil
    }
    end := offset + limit
    if end > len(books) {
        end = len(books)
    }
    return books[offset:end], total, nil
}

func (m *MockBookRepository) Update(id string, book *model.Book) error {
    return nil // テスト用なので常に成功
}

func (m *MockBookRepository) Delete(id string) error {
    return nil // テスト用なので常に成功
}


// AuthMiddleware のモック版
func authMiddlewareMock(userID uint) gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Set("user_id", userID)
    c.Next()
  }
}

func setupRouter() *gin.Engine {
  r := gin.Default()
  
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

func TestListBooks(t *testing.T) {
    gin.SetMode(gin.TestMode)

    r := gin.Default()
    
    mockRepo := &MockBookRepository{}
    mockService := &service.BookService{Repo: mockRepo}
    ctl := &BookController{Service: mockService}
    
    // AuthMiddleware を適用（userID=1）
    r.GET("/api/book", authMiddlewareMock(1), ctl.ListBooks)

    req, err := http.NewRequest(http.MethodGet, "/api/book", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Test Book")
    assert.Contains(t, w.Body.String(), "success")
}