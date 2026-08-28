package api

import (
	"encoding/json"
	"libraryassistant/model"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
func DecodeRecord(r *http.Request) (model.Record, error) {
	var v model.Record
	e := json.NewDecoder(r.Body).Decode(&v)
	return v, e
}
func (s *Server) MethodGuard(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method"})
		return false
	}
	return true
}
