package service

import (
	"libraryassistant/model"
	"libraryassistant/storage"
	"strings"
)

type Searcher struct{ Store *storage.Store }

func (q Searcher) ByMember(member string) ([]model.Record, error) {
	all, e := q.Store.ListRecords()
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	for _, r := range all {
		if strings.EqualFold(r.MemberID, member) {
			out = append(out, r)
		}
	}
	return out, nil
}
func (q Searcher) ByStatus(status string) ([]model.Record, error) {
	all, e := q.Store.ListRecords()
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	for _, r := range all {
		if r.Status == status {
			out = append(out, r)
		}
	}
	return out, nil
}
