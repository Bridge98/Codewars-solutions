package kata

import (
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	sol := regexp.MustCompile("-|_").Split(s, -1)

	for i, word := range sol[1:] {
		sol[i+1] = strings.Title(word)
	}

	return strings.Join(sol, "")
}
