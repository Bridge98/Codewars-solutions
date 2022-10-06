package kata

func Arithmetic(a,b int, operator string) int {

	// var to memorize the result of the operation
	result := 0

	// all the cases to handel arithmetic operations:
	if a > 0 && b > 0 {
		switch (operator) {
		case"add":
			result = a+b
		case"subtract":
			result = a-b
		case"multiply":
			result = a*b 
		case"divide":
			result = a/b
		}
	} else {
		return 0
	}

	return result
}