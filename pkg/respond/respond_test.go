package respond

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOkWritesJSONResponse(t *testing.T) {
	rec := httptest.NewRecorder()

	Ok(rec, http.StatusCreated, map[string]string{"id": "123"})

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rec.Code)
	}

	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("expected application/json content type, got %q", got)
	}

	expected := "{\"status\":\"Ok\",\"result\":{\"id\":\"123\"}}\n"
	if rec.Body.String() != expected {
		t.Fatalf("expected body %q, got %q", expected, rec.Body.String())
	}
}

func TestErrorMessageWritesJSONResponse(t *testing.T) {
	rec := httptest.NewRecorder()

	ErrorMessage(rec, http.StatusBadRequest, "invalid input")

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}

	expected := "{\"status\":\"Error\",\"error\":\"invalid input\"}\n"
	if rec.Body.String() != expected {
		t.Fatalf("expected body %q, got %q", expected, rec.Body.String())
	}
}
