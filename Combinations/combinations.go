package sprint

/* Create a Go function that generates 
a series of unique triplets of digits. 
Each triplet should consist of different digits, 
and they should be in ascending order, from the 
smallest to the largest. The triplets should be 
separated by a comma and a space.
For example, the function should return triplets 
like 012, 013, 014, 015.... Avoid combinations with all 
the same digits, like 000 or 999, and exclude 
triplets in descending order, like 987. */

import "fmt"

func Combinations() string {

	var triplets string

	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			for z := j + 1; z <= 9; z++ {
				triplets += fmt.Sprintf("%d%d%d", i, j, z)

				if i != 7 || j != 8 || z != 9 {
					triplets += ", "
				}
			}
		}

	}
	return triplets
}