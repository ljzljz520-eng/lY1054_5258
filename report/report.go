package report

import (
	"fmt"
	"libraryassistant/model"
	"sort"
)

type Summary struct {
	Total, Returned, Overdue int
	FineCents                int64
}

func Build(records []model.Record) Summary {
	var s Summary
	for _, r := range records {
		s.Total++
		if r.Status == "returned" || r.Status == "archived" {
			s.Returned++
			s.FineCents += r.FineCents
		}
		if r.IsOverdue(r.DueAt.AddDate(0, 0, 1)) {
			s.Overdue++
		}
	}
	return s
}
func Format(s Summary) string {
	return fmt.Sprintf("total=%d returned=%d overdue=%d fine=%d", s.Total, s.Returned, s.Overdue, s.FineCents)
}
func SortByDue(records []model.Record) {
	sort.Slice(records, func(i, j int) bool { return records[i].DueAt.Before(records[j].DueAt) })
}
