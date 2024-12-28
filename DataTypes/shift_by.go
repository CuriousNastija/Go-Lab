package sprint

/* Create a Go function that takes 
a single lowercase letter as a rune 
and an int 'step'. 
The function should shift the letter in the alphabet 
by the specified 'step' value, and return 
the resulting letter. For example, if you shift 'a' 
by 4 steps, it should become 'e'. 
Remember to handle looping back to 
the start of the alphabet if needed.
Only lower case characters are tested. */



func ShiftBy(r rune, step int) rune {
	// 'a' unicode value is 97. r-'a' calculates the position of the rune in the alphabet
	// int step should be converted to rune
	// modulus 26 ensures, that we will not go beyond lowercase alphabet
	return 'a' + (r - 'a' + rune(step))%26 
}