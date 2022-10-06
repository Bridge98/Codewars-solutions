package kata

import "strings"

func duplicate_count(s1 string) int {

	res := strings.ToLower(s1)
	count := 0
	occurences_string := make(map[rune]int) // map to count occurences

	// finding occurences in a string:
	for _, char := range res {
		occurences_string[char]++
	}

	// finding occurences that will repeat in the
	for _, v := range occurences_string {
		if v > 1 {
			count++
		}
	}

	return count
}