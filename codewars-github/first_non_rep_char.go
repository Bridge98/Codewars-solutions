package kata

import (
	_ "fmt"
	"strings"
	_ "unicode"
)

func FirstNonRepeating(str string) string {
  
  res := strings.ToLower(str)// add letter cases
  
	for _, char := range str {
		if strings.Count(res, strings.ToLower(string(char))) <= 1 {
			return string(char)
		}
	}

	return ""
}