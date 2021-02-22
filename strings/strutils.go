package strings

import "strings"

// IsEmpty empty string
func IsEmpty(s string) bool {
	if s == "" {
		return true
	}

	s = strings.Replace(s, " ", "", -1)
	if s == "" {
		return true
	}

	return false
}
