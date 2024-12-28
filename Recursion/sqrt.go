package sprint

func Sqrt(n int) int {

	if n < 0 {
		return 0
	}
	
	if n == 0 || n == 1 {
		return n
	}

	for i := 2; i*i <= n; i++ {
		if i*i == n {
			return i
		}
	}

	return 0

}