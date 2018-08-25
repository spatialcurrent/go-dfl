package dfl

import (
	"strings"
	"unicode"
)

func IsAttribute(s string) bool {
	return strings.HasPrefix(strings.TrimLeftFunc(s, unicode.IsSpace), "@")
}
