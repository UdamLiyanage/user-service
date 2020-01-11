package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
	DB.Collection = connect()
	// Run the other tests
	os.Exit(m.Run())
}

func newRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users/:id", readUser)
	r.POST("/users", createUser)
	r.POST("/users/attach/device", attachDevice)
	r.DELETE("/users/:id", deleteUser)
	r.POST("/users/remove/device", removeDevice)
	r.PUT("/users/:id", updateUser)
	return r
}

func testRequestStatusCode(method string, url string, body []byte, code int, t *testing.T) *httptest.ResponseRecorder {
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != code {
		t.Errorf("Status should be 201, got %d", w.Code)
	}
	return w
}

func testRequestBody(w *httptest.ResponseRecorder, field string, condition int, t *testing.T) {
	var response map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	value, exists := response[field]
	if !exists {
		t.Errorf("Wrong Response Format")
	}
	count, _ := strconv.Atoi(value)
	if count != condition {
		t.Errorf("Operation Not Working Properly!")
	}
}
