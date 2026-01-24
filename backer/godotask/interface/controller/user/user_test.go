package user

import (
	"os"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/godotask/infrastructure/db/model" 
)

func setupTestDB() {
	// テスト用のデータベースURLを設定
	os.Setenv("DATABASE_URL", "postgres://dbgodotask:dbgodotask@localhost:5432/test_db?sslmode=disable")
	model.InitDB()
	model.DB.AutoMigrate(&model.User{})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/api/register", Register) // Register関数を直接使用
	return r
}

func TestRegisterUser(t *testing.T) {
	// Ginのテストモードを設定
	gin.SetMode(gin.TestMode)
	setupTestDB()

	// テスト用のルーターをセットアップ
	r := setupRouter()

	// テスト用のリクエストボディを作成
	body := `{
		"username": "testuser",
		"email": "testuser@example.com",
		"password": "password123"
	}`

	req, err := http.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer([]byte(body)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// レスポンスを記録するためのRecorderを作成
	w := httptest.NewRecorder()

	// リクエストをルーターに送信
	r.ServeHTTP(w, req)

	// ステータスコードを確認
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスボディを確認
	assert.Contains(t, w.Body.String(), "User registered successfully")
}
