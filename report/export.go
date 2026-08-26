package report

import (
	"encoding/json"
	"libraryassistant/model"
)

func JSON(records []model.Record) ([]byte, error) { return json.MarshalIndent(records, "", "  ") }
func GroupByStatus(records []model.Record) map[string]int {
	m := map[string]int{}
	for _, r := range records {
		m[r.Status]++
	}
	return m
}
func TotalFine(records []model.Record) int64 {
	var n int64
	for _, r := range records {
		n += r.FineCents
	}
	return n
}
