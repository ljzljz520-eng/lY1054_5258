package model

import "time"

type Record struct {
	ID, BookID, MemberID          string
	BorrowedAt, DueAt, ReturnedAt time.Time
	FineCents                     int64
	Status                        string
}
type Profile struct {
	ID, Name, Email string
	Active          bool
	JoinedAt        time.Time
}
type Event struct {
	ID, RecordID, Kind, Detail string
	At                         time.Time
}
type Audit struct {
	ID, Actor, Action, Target string
	At                        time.Time
	Metadata                  map[string]string
}

func NewRecord(id, book, member string, due time.Time) Record {
	return Record{ID: id, BookID: book, MemberID: member, BorrowedAt: time.Now().UTC(), DueAt: due, Status: "borrowed"}
}
func (r Record) IsOverdue(now time.Time) bool { return r.Status != "returned" && now.After(r.DueAt) }
func (r *Record) Return(now time.Time, fine int64) {
	r.ReturnedAt = now
	r.FineCents = fine
	r.Status = "returned"
}
func NewProfile(id, name, email string) Profile {
	return Profile{ID: id, Name: name, Email: email, Active: true, JoinedAt: time.Now().UTC()}
}
func (p *Profile) Deactivate() { p.Active = false }
func NewEvent(id, record, kind, detail string) Event {
	return Event{ID: id, RecordID: record, Kind: kind, Detail: detail, At: time.Now().UTC()}
}
func NewAudit(id, actor, action, target string) Audit {
	return Audit{ID: id, Actor: actor, Action: action, Target: target, At: time.Now().UTC(), Metadata: map[string]string{}}
}
