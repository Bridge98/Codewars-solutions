package kata

import "strings"

func DuplicateEncode(word string) string {

	var res string
	word = strings.ToLower(word)

	for _, char := range word {
		if strings.Count(word, string(char)) == 1 {
			res += "("
		} else {
			res += ")"
		}
	}

	return res
}