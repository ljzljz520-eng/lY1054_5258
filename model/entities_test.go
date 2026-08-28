package model

import (
	"testing"
	"time"
)

func TestRecordLifecycle(t *testing.T) {
	r := NewRecord("1", "b", "m", time.Now())
	if r.Status != "borrowed" {
		t.Fail()
	}
	r.Return(time.Now(), 4)
	if r.Status != "returned" {
		t.Fail()
	}
}
