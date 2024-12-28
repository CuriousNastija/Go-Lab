package sprint

func IsPrime(n int) bool {

	if n <= 1 {
		return false
	}

	if n == 1 {
		return true
	}

	if n % 2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n % i == 0 {
			return false
		}
	}

	return true

}