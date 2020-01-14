package main

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestAttachDevice(t *testing.T) {
	attach := AttachDevice{
		UserID:   "5e0e2480031e0b729112ff2d",
		DeviceID: "test_id",
	}
	body, err := json.Marshal(attach)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("POST", "/users/attach/device", body, http.StatusCreated, t)
	testRequestBody(w, "ModifiedCount", 0, t)
}

func TestRemoveAttachedDevice(t *testing.T) {
	remove := make(map[string]string)
	remove["user_id"] = "5e0e2480031e0b729112ff2d"
	remove["device_id"] = "test_id"
	body, err := json.Marshal(&remove)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("POST", "/users/remove/device", body, http.StatusNotFound, t)
	testRequestBody(w, "ModifiedCount", 0, t)
}
