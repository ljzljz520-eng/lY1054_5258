package catalog

import (
	"encoding/csv"
	"io"
	"strconv"
)

func ImportCSV(r io.Reader) ([]Book, error) {
	rows, e := csv.NewReader(r).ReadAll()
	if e != nil {
		return nil, e
	}
	out := []Book{}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) < 5 {
			continue
		}
		year, _ := strconv.Atoi(row[4])
		out = append(out, Book{ID: row[0], Title: NormalizeTitle(row[1]), Author: row[2], ISBN: row[3], Year: year})
	}
	return out, nil
}
func Merge(c *Catalog, books []Book) int {
	n := 0
	for _, b := range books {
		if c.Add(b) == nil {
			n++
		}
	}
	return n
}
func ValidateBook(b Book) bool {
	return b.ID != "" && b.Title != "" && b.Author != "" && IsValidISBN(b.ISBN)
}
