package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateDeviceAndDelete(t *testing.T) {
	user := User{
		FirstName:     "First Name",
		LastName:      "Last Name",
		Email:         "test@email.com",
		Password:      "testpassword",
		ContactNumber: "0000000000",
		Address:       nil,
		Devices:       nil,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	body, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	r := newRouter()
	createRecorder := httptest.NewRecorder()
	createRequest, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	r.ServeHTTP(createRecorder, createRequest)
	if createRecorder.Code != http.StatusCreated {
		t.Errorf("Status should be 201, got %d", createRecorder.Code)
	}

	_ = json.NewDecoder(createRecorder.Body).Decode(&user)
	deleteRequest, _ := http.NewRequest("DELETE", "/users/"+user.ID.Hex(), nil)
	deleteRecorder := httptest.NewRecorder()
	r.ServeHTTP(deleteRecorder, deleteRequest)
	if deleteRecorder.Code != http.StatusNotFound {
		t.Errorf("Status should be 404, got %d", deleteRecorder.Code)
	}
}
