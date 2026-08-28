package catalog

import (
	"fmt"
	"strings"
)

func Label(b Book) string { return fmt.Sprintf("%s (%d) - %s", b.Title, b.Year, b.Author) }
func Slug(b Book) string {
	return strings.ToLower(strings.ReplaceAll(strings.TrimSpace(b.Title), " ", "-"))
}
func Categories(books []Book) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, b := range books {
		if !seen[b.Category] {
			seen[b.Category] = true
			out = append(out, b.Category)
		}
	}
	return out
}
