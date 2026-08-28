package storage

import (
	"errors"
	"go.etcd.io/bbolt"
	"time"
)

func (s *Store) Backup(path string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("closed")
	}
	return s.db.View(func(tx *bbolt.Tx) error { return tx.CopyFile(path, 0600) })
}
func (s *Store) Touch() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("audits"))
		return b.Put([]byte("last-touch"), []byte(time.Now().UTC().Format(time.RFC3339Nano)))
	})
}
