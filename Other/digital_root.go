package sprint

func DigitalRoot(n int) int {

	for n > 9 {
		n = sumOfDigits(n)
	}
	return n
}

func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}