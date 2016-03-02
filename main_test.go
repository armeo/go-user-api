package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handleIndex(res, req)

	if res.Code != 200 {
		t.Errorf("Expected: 200, Actual: %d", res.Code)
	}

	if res.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Expected: application/json, Actual: %s", res.Header().Get("Content-Type"))
	}

	expected := `{"message":"Hello"}`
	if res.Body.String() != expected {
		t.Errorf(`Expected: %s, Actual: %s`, expected, res.Body)
	}
}
