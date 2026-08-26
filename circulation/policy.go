package circulation

import (
	"libraryassistant/model"
	"time"
)

type Policy struct {
	GraceDays int
	DailyFine int64
	MaxLoans  int
}

func DefaultPolicy() Policy                           { return Policy{GraceDays: 2, DailyFine: 25, MaxLoans: 5} }
func DueWithGrace(r model.Record, p Policy) time.Time { return r.DueAt.AddDate(0, 0, p.GraceDays) }
func Eligible(active bool, current int, p Policy) bool {
	if !active {
		return false
	}
	if current < 0 {
		return false
	}
	return current < p.MaxLoans
}
