package main

import (
	"net/http"
	"testing"
)

func TestValidDeviceGet(t *testing.T) {
	//Object ID - 5e0e2480031e0b729112ff2d is always available as a test document
	testRequestStatusCode("GET", "/users/5e0e2480031e0b729112ff2d", nil, http.StatusOK, t)
}

func TestInvalidDeviceGet(t *testing.T) {
	testRequestStatusCode("GET", "/users/000000000000000000000000", nil, http.StatusNotFound, t)
}
