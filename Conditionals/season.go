package sprint

/* While if or if-else statements are very useful,
sometimes they are a little uncomfortable.
switch statement is very useful if one needs to perform more
checks for a value or if there are multiple cases for one check.
In this task you need to make a function that takes
a string, that can contain a month name.
If a month name is given, the season has to be returned.
Otherwise return "invalid input: " with the input appended to it. */

func Season(month string) string {
	var message string // declaring output string
	
	switch month {
	case "jan", "feb", "dec":
		message = "winter"
	case "mar", "apr", "may":
		message = "spring"
	case "jun", "jul", "aug":
		message = "summer"
	case "sep", "oct", "nov":
		message = "autumn"
	default:
		message = "invalid input: " + month // message concatenates the input string
	}
	return message
}