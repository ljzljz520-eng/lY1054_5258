package model

import (
	"errors"
	"strings"
	"time"
)

func ValidateRecord(r Record) error {
	if strings.TrimSpace(r.ID) == "" {
		return errors.New("id required")
	}
	if r.BookID == "" {
		return errors.New("book required")
	}
	if r.MemberID == "" {
		return errors.New("member required")
	}
	if r.DueAt.IsZero() {
		return errors.New("due required")
	}
	return nil
}
func ValidateProfile(p Profile) error {
	if p.ID == "" || p.Name == "" {
		return errors.New("identity required")
	}
	if !ValidEmail(p.Email) {
		return errors.New("email invalid")
	}
	return nil
}
func ClampDue(t time.Time) time.Time {
	if t.IsZero() {
		return time.Now().UTC().AddDate(0, 0, 14)
	}
	return t.UTC()
}
