package utils

import "strings"

// Multiline concatenates multiple strings into a series of newlines
// This function exists because it is not possible to escape ` in raw strings
// So there is no way we can write something like
// `json:"foo"` inside an assert
func Multiline(parts ...string) string {
	return strings.Join(parts, "\n")
}
