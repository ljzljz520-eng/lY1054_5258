package storage

import (
	"libraryassistant/model"
	"testing"
	"time"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/db"
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	s.SaveRecord(model.NewRecord("x", "b", "m", time.Now().Add(time.Hour)))
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.GetRecord("x"); e != nil {
		t.Fatal(e)
	}
}
