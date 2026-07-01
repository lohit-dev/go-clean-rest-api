package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
)

func TestHealthRoute(t *testing.T) {
	srv := New("4545", zap.NewNop(), nil)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	srv.Router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	expected := "{\"status\":\"Ok\",\"result\":{\"status\":\"online\"}}\n"
	if rec.Body.String() != expected {
		t.Fatalf("expected body %q, got %q", expected, rec.Body.String())
	}
}
