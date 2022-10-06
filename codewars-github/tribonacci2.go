package kata

func Tribonacci(signature [3]float64, n int) []float64 {

	res := make([]float64, n)

	for i := 0; i < 3; i++ {
    if i < n {      
		  res[i] = signature[i] // copy signature in resultant slice;
    }
	}

	for i := 3; i < n; i++ {
		res[i] = res[i-1] + res[i-2] + res[i-3] // tribonacci algorithm;
	}

	return res
}