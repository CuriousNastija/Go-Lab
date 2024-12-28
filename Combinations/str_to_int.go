package sprint

/* Create a Go function that mimics the behavior 
of the Atoi function. Atoi converts a string 
representing a number into an integer.

The function should return the integer value.
If the input string is not a valid number, it should return 0.
The function must handle signs such as + or -.
The exercise does not require you to return an error. */

func StrToInt(s string) int {

	result := 0
	negativeNumber := false

	// handle empty string case
	if len(s) == 0 {
		return 0
	}

	// check for negative or positive sign at the beginning
	if s[0] == '-' {
			negativeNumber = true
			s = s[1:] // remove the negative sign
		} else if s[0] == '+' {
			s = s[1:] // // remove the positive sign
		}

	for i:= 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0 // return 0 if there is any invalid character
		}
		
		// convert character to digit and build the result
		digit := int(s[i] - '0')
		result = result*10 + digit

		 
	}
	
	// apply negative result if needed
	if negativeNumber {
		result = -result
	}

	return result
}