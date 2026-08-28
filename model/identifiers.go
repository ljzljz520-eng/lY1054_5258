package model

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func NormalizeID(v string) string { return strings.ToLower(strings.TrimSpace(v)) }
func StableID(parts ...string) string {
	h := sha256.New()
	for _, p := range parts {
		h.Write([]byte(p))
		h.Write([]byte{0})
	}
	return hex.EncodeToString(h.Sum(nil))[:16]
}
func ValidEmail(v string) bool { return strings.Contains(v, "@") && len(v) > 3 }
