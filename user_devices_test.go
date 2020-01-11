package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeviceAttach(t *testing.T) {
	attach := AttachDevice{
		UserID:     "5e0e2480031e0b729112ff2d",
		DeviceID:   "test_id",
		DeviceName: "Test Device Name",
	}
	body, err := json.Marshal(attach)
	if err != nil {
		t.Fatal(err)
	}
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/attach/device", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Status should be 201, got %d", w.Code)
	}
}
