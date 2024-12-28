package sprint

/* Create a Go function that finds all possible 
combinations of two two-digit numbers. 
The pairs should be in ascending order, and each number 
should be padded with a leading zero if it's less than 10. 
The pairs should be separated by a comma and a space. 
Each number within a pair should be separated by a space. 
The function should return a string containing these pairs. */

import "fmt"

func Pairs() string {

	var resultPairs string
	
	// Loop through the first and second numbers
	for i:= 0; i <= 99; i++ {
		for j:= i + 1; j <= 99; j++ {
			//Format the pair with leading zero
			resultPairs += fmt.Sprintf("%02d %02d", i, j) 

			// omit comma in the end
			if i != 98 || j != 99 {
				resultPairs += ", "
			}
		}
	}

	return resultPairs
}