package lstrings

import (
	"strings"
	"unicode"
)

func SplitByUpper(s string) []string {
	var words []string
	l := 0
	for s := s; s != ""; s = s[l:] {
		l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if l <= 0 {
			l = len(s)
		}
		words = append(words, s[:l])
	}

	return words
}
