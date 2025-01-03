package sprint

/* Write a Go function that takes an integer 
representing a year and returns a boolean 
value indicating whether that year is a 
leap year or not. A leap year is a year that is 
evenly divisible by 4, except for years that are divisible by 100 but not by 400. */

func IsLeapYear(year int) bool {

	if year % 400 == 0 { // check if year divisible by 400
		return true
	} 
	
	if year % 100 == 0 { // check if year divisible by 100
		return false
	}
	
	if year % 4 == 0 { // check if year divisible by 4
		return true
	} 
	
	return false
}