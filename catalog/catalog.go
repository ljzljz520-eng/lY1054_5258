package catalog

import (
	"errors"
	"sort"
	"strings"
	"sync"
	"time"
)

type Book struct {
	ID, Title, Author, ISBN, Category string
	Year                              int
	Available                         bool
	AddedAt                           time.Time
}
type Catalog struct {
	mu    sync.RWMutex
	books map[string]Book
}

func New() *Catalog { return &Catalog{books: map[string]Book{}} }
func (c *Catalog) Add(b Book) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if b.ID == "" || b.Title == "" {
		return errors.New("book identity required")
	}
	if _, ok := c.books[b.ID]; ok {
		return errors.New("duplicate book")
	}
	b.AddedAt = time.Now().UTC()
	b.Available = true
	c.books[b.ID] = b
	return nil
}
func (c *Catalog) Remove(id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.books[id]; !ok {
		return errors.New("not found")
	}
	delete(c.books, id)
	return nil
}
func (c *Catalog) Get(id string) (Book, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	b, ok := c.books[id]
	if !ok {
		return Book{}, errors.New("not found")
	}
	return b, nil
}
func (c *Catalog) SetAvailability(id string, v bool) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	b, ok := c.books[id]
	if !ok {
		return errors.New("not found")
	}
	b.Available = v
	c.books[id] = b
	return nil
}
func (c *Catalog) Search(q string) []Book {
	c.mu.RLock()
	defer c.mu.RUnlock()
	q = strings.ToLower(strings.TrimSpace(q))
	out := []Book{}
	for _, b := range c.books {
		if q == "" || strings.Contains(strings.ToLower(b.Title), q) || strings.Contains(strings.ToLower(b.Author), q) || strings.Contains(strings.ToLower(b.Category), q) {
			out = append(out, b)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Title < out[j].Title })
	return out
}
func (c *Catalog) Available() []Book {
	all := c.Search("")
	out := []Book{}
	for _, b := range all {
		if b.Available {
			out = append(out, b)
		}
	}
	return out
}
func (c *Catalog) Count() int     { c.mu.RLock(); defer c.mu.RUnlock(); return len(c.books) }
func (c *Catalog) Export() []Book { return c.Search("") }
func IsValidISBN(v string) bool {
	v = strings.ReplaceAll(v, "-", "")
	return len(v) == 10 || len(v) == 13
}
func NormalizeTitle(v string) string { return strings.Join(strings.Fields(strings.TrimSpace(v)), " ") }
