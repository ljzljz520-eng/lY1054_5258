package report

import (
	"libraryassistant/model"
	"time"
)

func Aging(records []model.Record, now time.Time) map[string]int {
	m := map[string]int{"0-7": 0, "8-30": 0, "31+": 0}
	for _, r := range records {
		d := int(now.Sub(r.DueAt).Hours() / 24)
		if d <= 7 {
			m["0-7"]++
		} else if d <= 30 {
			m["8-30"]++
		} else {
			m["31+"]++
		}
	}
	return m
}
func Active(records []model.Record) int {
	n := 0
	for _, r := range records {
		if r.Status == "borrowed" {
			n++
		}
	}
	return n
}
