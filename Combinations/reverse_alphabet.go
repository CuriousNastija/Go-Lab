package sprint

/* Create a function that takes 
a step value n as input. Starting from z 
in the Latin alphabet, it should return 
the lowercase alphabet in reverse order 
as a string. For each letter, you skip n-1 
letters. If n is 0 or negative, use a default step of 1. */

func ReverseAlphabet(step int) string {

	var resultString string

	if step <= 0 {
		step = 1
	} 

	// stepRune := rune(step) - we can convert step to rune earlier for readability

	for i := 'z'; i >= 'a'; i -= rune(step) { // here i is a rune
		resultString += string(i) // converting i into string
	}
	return resultString

}