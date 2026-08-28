package storage

import (
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"libraryassistant/model"
	"path/filepath"
	"sync"
)

var buckets = []string{"records", "profiles", "events", "audits"}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(filepath.Clean(path), 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, n := range buckets {
			if _, x := tx.CreateBucketIfNotExists([]byte(n)); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return errors.New("closed")
	}
	e := s.db.Close()
	s.db = nil
	return e
}
func put[T any](s *Store, bucket, key string, v T) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), data) })
}
func get[T any](s *Store, bucket, key string, out *T) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("closed")
	}
	return s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if v == nil {
			return errors.New("not found")
		}
		return json.Unmarshal(v, out)
	})
}
func (s *Store) SaveRecord(v model.Record) error { return put(s, "records", v.ID, v) }
func (s *Store) GetRecord(id string) (model.Record, error) {
	var v model.Record
	e := get(s, "records", id, &v)
	return v, e
}
func (s *Store) SaveProfile(v model.Profile) error { return put(s, "profiles", v.ID, v) }
func (s *Store) GetProfile(id string) (model.Profile, error) {
	var v model.Profile
	e := get(s, "profiles", id, &v)
	return v, e
}
func (s *Store) SaveEvent(v model.Event) error { return put(s, "events", v.ID, v) }
func (s *Store) SaveAudit(v model.Audit) error { return put(s, "audits", v.ID, v) }
func (s *Store) Count(bucket string) (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return 0, errors.New("closed")
	}
	n := 0
	e := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return errors.New("bucket")
		}
		return b.ForEach(func(_, v []byte) error {
			if v != nil {
				n++
			}
			return nil
		})
	})
	return n, e
}
