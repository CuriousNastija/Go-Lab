package sprint

/* Create a function that takes a positive 
integer as input and returns the corresponding 
number of letters from the Latin alphabet. 
The input integer won't be larger than the alphabet's length. */



func AlphabetMastery(n int) string {

	//var alphabet string = "abcdefghijklmnopqrstuvwxyz"
	var resultString string
	
	for i := 0; i < n; i++ {
		resultString += string('a' + i)
	} 
	return resultString
}