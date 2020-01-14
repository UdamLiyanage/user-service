package main

import (
	"encoding/json"
	"net/http"
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
	createRecorder := testRequestStatusCode("POST", "/users", body, http.StatusCreated, t)

	_ = json.NewDecoder(createRecorder.Body).Decode(&user)
	testRequestStatusCode("DELETE", "/users/"+user.ID, nil, http.StatusNotFound, t)
}
