package sprint

/* Create a Go function that converts a valid 
string representation of a number into an integer.
If the string is not a valid number, the function 
should return 0. The input will only contain one 
or more digits, and you don't need to handle signs like + or -. */

func SimpleStrToInt(s string) int {

	result := 0

	// iterate through each character in the string
	for i:= 0; i < len(s); i++ {
		// check if the character is a valid digit ('0' to '9')
		// if s[i] < 0 || s[i] > 9 {
		//return 0
	//}

	// convert character to corresponding digit
		digit := int(s[i] - '0')

	// build the integer result
		result = result*10 + digit
	}

	return result
}