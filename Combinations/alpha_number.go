package sprint

/* Create a Go function that takes an integer 
as input and returns a string representing that 
integer. If the number is negative, preserve 
the minus sign. Replace the digits with lowercase 
letters, where 0 becomes a, 1 becomes b, and so on up 
to 9, which becomes j. */

func AlphaNumber(n int) string {

	var sequence string
	var negativeNumber bool

	// handle zero case directly
	if n == 0 {
		return "a"
	}

	// handle negative numbers with bool var
	negativeNumber = n < 0

	if negativeNumber {
		n = -n
	}

	/*
	for n > 0 {
		digit := n % 10
		sequence += string('a' + digit) + sequence
		n /= 10
	}
	*/

	for i := n; i > 0; i /= 10 {
		digit := i % 10
		sequence = string('a' + digit) + sequence
	}

	if negativeNumber {
		sequence = "-" + sequence
	}

	return sequence
}