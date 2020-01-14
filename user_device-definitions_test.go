package main

import (
	"encoding/json"
	"testing"
)

func TestAttachDeviceDefinition(t *testing.T) {
	attach := AttachDeviceDefinition{
		UserID:       "5e0e2480031e0b729112ff2d",
		DefinitionID: "test_definition_id",
	}
	body, err := json.Marshal(attach)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("POST", "/users/attach/device-definition", body, 200, t)
	testRequestBody(w, "ModifiedCount", 0, t)
}

func TestRemoveDeviceDefinition(t *testing.T) {
	attach := AttachDeviceDefinition{
		UserID:       "5e0e2480031e0b729112ff2d",
		DefinitionID: "test_definition_id",
	}
	body, err := json.Marshal(attach)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("POST", "/users/remove/device-definition", body, 200, t)
	testRequestBody(w, "ModifiedCount", 0, t)
}
