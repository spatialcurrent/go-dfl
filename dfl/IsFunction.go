package dfl

import (
	"strings"
)

func IsFunction(s string) bool {
	return len(s) >= 3 && strings.Contains(s, "(") && strings.HasSuffix(s, ")")
}
