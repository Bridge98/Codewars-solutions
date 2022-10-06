package kata

import (
  "fmt"
  "strings"
  "strconv"
)

func HighAndLow(in string) string {

	var high, low int

	for i, char := range strings.Fields(in) {

		num, _ := strconv.Atoi(string(char))
    
    // assign the first number of the string to high and low, to return the correct number; 
		if i == 0 {
			high, low = num, num
		}
  
    // plan to find an extreme value:
    // find the minimum value
		if num < low {
			low = num
		}
    
    // find the maximum value
		if num > high {
			high = num
		}
	}
  
  // easiest way to return the string "high, low"
	return fmt.Sprintf("%d %d",high,low) 
}