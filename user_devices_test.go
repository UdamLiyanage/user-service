package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAttachDevice(t *testing.T) {
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

func TestRemoveAttachedDevice(t *testing.T) {
	remove := make(map[string]string)
	remove["user_id"] = "5e0e2480031e0b729112ff2d"
	remove["device_id"] = "test_id"
	body, err := json.Marshal(&remove)
	if err != nil {
		t.Fatal(err)
	}
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/remove/device", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("Status should be 404, got %d", w.Code)
	}
}
