package api

import (
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	s := &Server{}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/health", nil)
	s.Health(w, r)
	if w.Code != 200 {
		t.Fail()
	}
}
