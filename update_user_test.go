package main

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestUpdateDeviceValid(t *testing.T) {
	device := User{
		FirstName: "Updated Name",
	}
	body, err := json.Marshal(device)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("PUT", "/users/5e103f782aa554ae4e6abb8b", body, http.StatusOK, t)
	testRequestBody(w, "ModifiedCount", 0, t)
}

func TestUpdateDeviceInvalid(t *testing.T) {
	device := User{
		FirstName: "Updated Name",
	}
	body, err := json.Marshal(device)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("PUT", "/users/000000000000000000000000", body, http.StatusOK, t)
	testRequestBody(w, "MatchedCount", 0, t)
}
