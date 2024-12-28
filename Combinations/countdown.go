package sprint

/* Create a function that takes a one-digit 
integer as input and returns a countdown string. 
The countdown should start from the given number, 
skip every second number, and end with 0!. 
For example, if the input is 7, the function 
should return "7, 5, 3, 1, 0!". */

func Countdown(n int) string {

	var sequence string
	var step int = 2

	if n == 0 {
		sequence = "0!"
	}

	for i := n; i >= 1; i -= step {
		if i < 10 {
			sequence += string('0' + i) // shift the unicode value of '0' (48) to the corresponding digit
		} else {
			sequence += string ('0' + (i / 10))
			sequence += string ('0' + (i % 10))
		}
			sequence += ", "
		if i == 1 || i == 2 {
			sequence += "0!"
	} 
}

return sequence

}
