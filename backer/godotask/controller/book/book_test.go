package book

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/api/book", AddBookAction)
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
    assert.Contains(t, w.Body.String(), "Book added successfully")
}

func TestUpdateBook(t *testing.T) {
	// Ginのテストモードを設定
	gin.SetMode(gin.TestMode)

	// テスト用のルーターをセットアップ
	r := gin.Default()
	r.PUT("/api/updatebook/:id", UpdateBookAction) // UpdateBookActionを直接使用

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
	assert.Contains(t, w.Body.String(), "Book updated successfully")
}

func TestDeleteBook(t *testing.T) {
	// Ginのテストモードを設定
	gin.SetMode(gin.TestMode)

	// テスト用のルーターをセットアップ
	r := gin.Default()
	r.DELETE("/api/deletebook/:id", DeleteBookAction) // DeleteBookActionを直接使用

	req, err := http.NewRequest(http.MethodDelete, "/api/deletebook/1", nil)
	if err != nil {
			t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Book deleted successfully")
}
