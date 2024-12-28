package sprint

func LCM(a, b int) int {

	var result int

	result = (a * b)/GCD(a, b)

	return result

}

func GCD(a, b int) int {

	if b == 0 {
		return a
	}

	tmp := a
	a = b
	b = tmp % a
	return GCD(a, b)

}