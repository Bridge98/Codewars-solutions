package kata

func Multiple3And5(number int) int {

	sum := 0

	// find all the multiples of a number;
	for i := 1; i < number; i++ {

		if i%3 == 0 || i%5 == 0 {
			if i < number {
				sum += i
			}
		}
	}

	return sum
}