package catalog

import "testing"

func TestCatalogSearch(t *testing.T) {
	c := New()
	c.Add(Book{ID: "1", Title: "Go", Author: "A", ISBN: "1234567890"})
	if len(c.Search("go")) != 1 {
		t.Fail()
	}
}
