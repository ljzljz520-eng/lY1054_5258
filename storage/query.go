package storage

import (
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"libraryassistant/model"
)

func (s *Store) ListRecords() ([]model.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return nil, errors.New("closed")
	}
	out := []model.Record{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			if v == nil {
				return nil
			}
			var r model.Record
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
