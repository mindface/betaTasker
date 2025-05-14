package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/godotask/server"
)

func IndexDisplayAction(c *gin.Context) {
	c.String(http.StatusOK, "Expected content")
}

func TestGetRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := server.GetRouter()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "<h1>記録追加</h1>")
}