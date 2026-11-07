package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func request(t *testing.T, method, target string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(method, target, nil)
	rec := httptest.NewRecorder()
	New().Handler().ServeHTTP(rec, req)
	return rec
}

func TestHealth(t *testing.T) {
	rec := request(t, http.MethodGet, "/healthz")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Body.String(); got != "ok\n" {
		t.Fatalf("body = %q, want %q", got, "ok\n")
	}
}

func TestReady(t *testing.T) {
	rec := request(t, http.MethodGet, "/readyz")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestHelloDefault(t *testing.T) {
	rec := request(t, http.MethodGet, "/api/v1/hello")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	var body helloResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if body.Message != "hello, world" {
		t.Fatalf("message = %q, want %q", body.Message, "hello, world")
	}
}

func TestHelloWithName(t *testing.T) {
	rec := request(t, http.MethodGet, "/api/v1/hello?name=opencode")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	var body helloResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if body.Message != "hello, opencode" {
		t.Fatalf("message = %q, want %q", body.Message, "hello, opencode")
	}
}

func TestMethodNotAllowed(t *testing.T) {
	rec := request(t, http.MethodPost, "/healthz")
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}
}
