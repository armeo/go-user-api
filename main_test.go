package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleIndexReturnWithStatusOK(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handleIndex(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected: %v, Actual: %v", "200", res.Code)
	}
}
