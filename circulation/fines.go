package circulation

import (
	"errors"
	"libraryassistant/model"
	"time"
)

func CalculateFine(r model.Record, now time.Time, daily int64) (int64, error) {
	if daily < 0 {
		return 0, errors.New("invalid rate")
	}
	if !r.IsOverdue(now) {
		return 0, nil
	}
	days := int(now.Sub(r.DueAt).Hours() / 24)
	if days < 1 {
		days = 1
	}
	return int64(days) * daily, nil
}
func ValidateLoan(book, member string, due time.Time) error {
	if book == "" || member == "" {
		return errors.New("book and member required")
	}
	if due.Before(time.Now().UTC()) {
		return errors.New("due date in past")
	}
	return nil
}
func CanReturn(r model.Record) error {
	if r.Status == "returned" {
		return errors.New("already returned")
	}
	if r.Status != "borrowed" {
		return errors.New("invalid status")
	}
	return nil
}
