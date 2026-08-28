package circulation

import (
	"fmt"
	"libraryassistant/model"
)

type Notice struct{ Recipient, Subject, Body string }

func OverdueNotice(r model.Record, fine int64) Notice {
	return Notice{Recipient: r.MemberID, Subject: "Overdue item", Body: fmt.Sprintf("record %s has fine %d cents", r.ID, fine)}
}
func ReturnNotice(r model.Record) Notice {
	return Notice{Recipient: r.MemberID, Subject: "Return confirmed", Body: "book " + r.BookID + " returned"}
}
func RenderNotice(n Notice) string { return n.Subject + "\n" + n.Body }
