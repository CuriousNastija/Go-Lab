package sprint

func FactorialIterative(n int) int {

	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	var factorial int = 1
	for i := 1; i <= n; i++ {
		if factorial > (1<<63-1)/i {
			return 0
		}
		factorial *= i
	}

	return factorial

}