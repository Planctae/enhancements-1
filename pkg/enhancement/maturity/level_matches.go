package maturity

import (
	"strings"
)

func LevelMatches(given, matches string) bool {
	a := strings.ToLower(given)
	b := strings.ToLower(matches)

	return strings.EqualFold(a, b)
}
