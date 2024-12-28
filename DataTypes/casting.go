package sprint

/* Write a function that accepts 
a float64 value, rounds it to the nearest integer, 
and then returns the result as an int. */

import (
	"math"
)

func Casting(n float64) int {
	res := math.Round(n)
	return int(res) // converting float32 into integer
}