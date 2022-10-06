package kata

func DigitalRoot(n int) int {

	var sum int
	if n < 10 {
    return n
  } else {
  	for n != 0 {
		  sum += n%10
		  n /= 10
	  } 
  }

	return DigitalRoot(sum)
}