package service

import (
	"context"
	"libraryassistant/model"
	"libraryassistant/storage"
	"testing"
	"time"
)

func TestBusinessChain23(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x.db")
	defer s.Close()
	l := New(s)
	base := time.Now().UTC()
	r := model.NewRecord("r1", "b1", "m1", base.Add(time.Hour))
	if e := l.Borrow(context.Background(), r); e != nil {
		t.Fatal(e)
	}
	fine, e := l.Return(context.Background(), "r1", base.Add(49*time.Hour))
	if e != nil || fine == 0 {
		t.Fatalf("fine=%d err=%v", fine, e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x.db")
	defer s.Close()
	l := New(s)
	if e := l.Receive(context.Background(), "r2", "b2", "m2", time.Now().Add(time.Hour)); e != nil {
		t.Fatal(e)
	}
	if e := l.Confirm(context.Background(), "r2"); e != nil {
		t.Fatal(e)
	}
	if e := l.Archive("r2"); e == nil {
		t.Fail()
	}
}
func TestWorkflowThree(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x.db")
	defer s.Close()
	l := New(s)
	if e := l.AuditAction(context.Background(), "staff", "submit", "r3"); e != nil {
		t.Fatal(e)
	}
	if e := l.RecordEvent("r3", "review", "ok"); e != nil {
		t.Fatal(e)
	}
	if _, e := l.Track("r3"); e == nil {
		t.Fail()
	}
}
