package kata

import (
	"sort"
	"strings"
)

func RemoveDuplictaesStrings(sl_str []string) []string {
	
	keys := make(map[string]bool)
	list := []string{} // = var list []string

	//if the key is not equal to the already present value than we append it
	for _, entry := range sl_str {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
} 

func InArray(a1 []string, a2 []string) []string {

	// sort the two slice of strings
	sort.Strings(a1)
	sort.Strings(a2)
	var r []string

	for _, word := range a1 {
		for _, w := range a2 {
			if strings.Contains(w, word) { // check if a2[w] contains word
				r = append(r, word)
			}
		}
	}

	return (RemoveDuplictaesStrings(r)) // return the slice whitout repetition
}