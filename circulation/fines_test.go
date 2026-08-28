package circulation

import (
	"libraryassistant/model"
	"testing"
	"time"
)

func TestCalculateFine(t *testing.T) {
	r := model.NewRecord("1", "b", "m", time.Now().Add(-49*time.Hour))
	n, e := CalculateFine(r, time.Now(), 25)
	if e != nil || n < 50 {
		t.Fail()
	}
}
