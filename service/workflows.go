package service

import (
	"context"
	"fmt"
	"libraryassistant/model"
	"time"
)

func (l *Library) Receive(ctx context.Context, id, book, member string, due time.Time) error {
	return l.Borrow(ctx, model.NewRecord(id, book, member, due))
}
func (l *Library) ProcessReturn(ctx context.Context, id string, now time.Time) (int64, error) {
	return l.Return(ctx, id, now)
}
func (l *Library) Confirm(ctx context.Context, id string) error {
	r, e := l.Find(id)
	if e != nil {
		return e
	}
	return l.Store.SaveEvent(model.NewEvent("confirm-"+id, id, "confirm", fmt.Sprintf("status=%s", r.Status)))
}
func (l *Library) Track(id string) (string, error) {
	r, e := l.Find(id)
	if e != nil {
		return "", e
	}
	return r.Status, nil
}
