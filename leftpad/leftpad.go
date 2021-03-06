package leftpad

import (
	"strings"
	"unicode/utf8"
)

var defaultChar = ' '

// Format takes in a string and an int and returns the string
// left-padded with spaces to the length of the int.
func Format(s string, size int) string {
	return FormatRune(s, size, defaultChar)
}

// FormatRune takes in a string, an int and a rune and returns
// the string left-padded with spaces to the lenght of the int
func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}
	out := strings.Repeat(string(r), size-l) + s
	return out
}
