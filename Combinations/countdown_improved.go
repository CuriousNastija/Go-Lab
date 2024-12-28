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

	for i := n; i >= 1; i -= step {
		sequence += string('0' + i) // convert digit to character
		sequence += ", "
	} 
	sequence += "0!" // append 0! at the end

return sequence

}
