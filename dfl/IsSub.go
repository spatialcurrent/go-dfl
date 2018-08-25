package dfl

import (
	"strings"
)

func IsSub(s string) bool {
	return len(s) >= 2 && strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")")
}
