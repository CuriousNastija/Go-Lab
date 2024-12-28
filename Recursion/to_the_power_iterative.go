package sprint

func ToThePowerIterative(n int, power int) int {

	if power < 0 {
		return 0
	}

	if n == 0 {
		if power == 0 {
			return 1
		}
		return 0
	}

	if power == 0 {
		return 1
	}

	result := 1
	for i := 0; i < power; i++ {
		result *= n
	}

	return result

}