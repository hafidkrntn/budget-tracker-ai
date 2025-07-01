package utilities

import "strings"

func IsNilOrEmpty(s *string) bool {
	return s == nil || strings.TrimSpace(*s) == ""
}
