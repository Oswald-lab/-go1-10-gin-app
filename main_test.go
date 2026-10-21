package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func do(method, path, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	newRouter().ServeHTTP(rec, req)
	return rec
}

func TestHello(t *testing.T) {
	rec := do(http.MethodGet, "/hello", "")
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "Hello from Gin") {
		t.Fatalf("unexpected response: %d %s", rec.Code, rec.Body.String())
	}
}

func TestPathParam(t *testing.T) {
	rec := do(http.MethodGet, "/items/42", "")
	if !strings.Contains(rec.Body.String(), `"id":"42"`) {
		t.Fatalf("unexpected body: %s", rec.Body.String())
	}
}

func TestEchoRejectsMissingMessage(t *testing.T) {
	rec := do(http.MethodPost, "/echo", `{}`)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}

func TestEchoAcceptsMessage(t *testing.T) {
	rec := do(http.MethodPost, "/echo", `{"message":"hi"}`)
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), `"echo":"hi"`) {
		t.Fatalf("unexpected response: %d %s", rec.Code, rec.Body.String())
	}
}
