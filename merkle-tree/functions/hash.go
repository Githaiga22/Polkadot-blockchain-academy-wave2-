package functions

import (
	"crypto/sha256"
	"fmt"
)

// Hash data using SHA256
func hash(data string) string {
	h := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", h)
}
