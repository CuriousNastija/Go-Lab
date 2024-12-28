package sprint

/* Write a Go function that takes a lowercase 
letter rune ('a'-'z') as input and returns its reverse 
letter in the Latin alphabet as a rune integer value. 
For example, 'a' should be transformed to 'z', 'y' 
should become 'b', and so on. The function should 
maintain the case of the input letter. */

func ReverseAlphabetValue(ch rune) rune {
	if ch >= 'a' && ch <= 'z' { // lowercase alphabet range
		return 'z' - (ch - 'a') // сh - 'a' calculates the position of ch in the alphabet
		// 'z' - (ch - 'a') reverses position of the letter
	} else if ch >= 'A' && ch <= 'Z' {
		return 'Z' - (ch - 'A')
	}
	return ch // return the original character if it is not a letter
}