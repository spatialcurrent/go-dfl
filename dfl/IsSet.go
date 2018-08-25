package dfl

import (
	"strings"
)

func IsSet(s string) bool {
	return len(s) >= 2 && strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}")
}
