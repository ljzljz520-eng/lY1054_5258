package service

import (
	"context"
	"errors"
	"fmt"
	"libraryassistant/circulation"
	"libraryassistant/model"
	"libraryassistant/storage"
	"time"
)

type Library struct {
	Store *storage.Store
	Rate  int64
}

func New(s *storage.Store) *Library { return &Library{Store: s, Rate: 25} }
func (l *Library) Register(ctx context.Context, p model.Profile) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	if p.ID == "" || p.Name == "" {
		return errors.New("invalid profile")
	}
	return l.Store.SaveProfile(p)
}
func (l *Library) Borrow(ctx context.Context, r model.Record) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	if e := circulation.ValidateLoan(r.BookID, r.MemberID, r.DueAt); e != nil {
		return e
	}
	if e := l.Store.SaveRecord(r); e != nil {
		return e
	}
	return l.Store.SaveEvent(model.NewEvent(fmt.Sprintf("borrow-%s", r.ID), r.ID, "borrow", "loan created"))
}
func (l *Library) Return(ctx context.Context, id string, now time.Time) (fine int64, err error) {
	r, e := l.Store.GetRecord(id)
	defer func() {
		if e == nil && r.Status == "returned" {
			fine = 0
		}
	}()
	if e != nil {
		return 0, e
	}
	if e = circulation.CanReturn(r); e != nil {
		return 0, e
	}
	fine, e = circulation.CalculateFine(r, now, l.Rate)
	if e != nil {
		return 0, e
	}
	r.Return(now, fine)
	if e = l.Store.SaveRecord(r); e != nil {
		return 0, e
	}
	if e = l.Store.SaveEvent(model.NewEvent("return-"+id, id, "return", "processed")); e != nil {
		return 0, e
	}
	return fine, nil
}
func (l *Library) Find(id string) (model.Record, error) { return l.Store.GetRecord(id) }
func (l *Library) Archive(id string) error {
	r, e := l.Find(id)
	if e != nil {
		return e
	}
	if r.Status != "returned" {
		return errors.New("must return first")
	}
	r.Status = "archived"
	return l.Store.SaveRecord(r)
}
