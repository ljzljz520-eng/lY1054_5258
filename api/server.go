package api

import (
	"encoding/json"
	"libraryassistant/service"
	"net/http"
	"time"
)

type Server struct{ Library *service.Library }

func New(l *service.Library) *Server { return &Server{Library: l} }
func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func (s *Server) Return(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fine, e := s.Library.Return(r.Context(), id, time.Now().UTC())
	resp := map[string]any{"success": e == nil, "fine_cents": fine}
	if e != nil {
		resp["error"] = e.Error()
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(resp)
}
func (s *Server) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/health", s.Health)
	m.HandleFunc("/return", s.Return)
	return m
}
