package sprint

func GCD(a, b int) int {

	if b == 0 {
		return a
	}

	tmp := a
	a = b
	b = tmp % a
	return GCD(a, b)

}