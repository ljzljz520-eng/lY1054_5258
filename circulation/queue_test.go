package circulation

import (
	"libraryassistant/model"
	"testing"
)

func TestQueue(t *testing.T) {
	q := &Queue{}
	q.Enqueue(model.Record{ID: "x"})
	if q.Len() != 1 {
		t.Fail()
	}
	if _, ok := q.Dequeue(); !ok {
		t.Fail()
	}
}
