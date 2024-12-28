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

	var firstPair []string //slice
	var secondPair []string
	var resultPairs string

	// Generate first pair of numbers (0-99)
		for i := 0; i <= 99; i ++ {
			if i < 10 {
			firstPair = append(firstPair, fmt.Sprintf("0%d", i))
		} else {
			firstPair = append(firstPair, fmt.Sprintf("%d", i))
		}
	}
	
	// Generate second pair of numbers (0-99)
	for i := 0; i <= 99; i ++ {
			if i < 10 {
			secondPair = append(secondPair, fmt.Sprintf("0%d", i))
		} else {
			secondPair = append(secondPair, fmt.Sprintf("%d", i))
		}
	}
	
	// Loop through both pairs, ensuring that i always less then j
	for i:= 0; i < len(firstPair); i++ {
		for j:= i + 1; j < len(secondPair); j++ {
			resultPairs += fmt.Sprintf("%s %s", firstPair[i], secondPair[j]) 

			// omit comma in the end
			if i != 98 || j != 99 {
				resultPairs += ", "
			}
		}
	}

	return resultPairs
}