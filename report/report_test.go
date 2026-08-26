package report

import (
	"libraryassistant/model"
	"testing"
)

func TestSummary(t *testing.T) {
	s := Build([]model.Record{{Status: "returned", FineCents: 2}})
	if s.FineCents != 2 {
		t.Fail()
	}
}
