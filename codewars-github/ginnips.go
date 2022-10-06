package kata

import (
	_ "fmt"
	"regexp"
	"strings"
)

// func to reverse a string in golang
func ReverseWord(s string) string {
    
    rns := []rune(s) // convert to rune
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
  
        // swap the letters of the string,
        // like first with last and so on.
        rns[i], rns[j] = rns[j], rns[i]
    }
  
    // return the reversed string.
    return string(rns)
}

// Spin the words!
func SpinWords(str string) string {

	res := regexp.MustCompile(" ").Split(str, -1) // split the words without space and verify that there are no errors, else panic
	var sol []string

	for i, word := range res {
		if len(word) >= 5 {
			res[i] = ReverseWord(word)
		}

		sol = append(sol, string(res[i]))
	}

	return strings.Join(sol, " ")
}