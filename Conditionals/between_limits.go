package sprint

/* Build a function that accepts two runes as input 
and returns a string containing all the runes that come 
between these two runes in the alphabet, in the correct order. 
For instance, if the input runes are 'f' and 'j', 
the function should return "ghi" regardless of the order 
in which the runes are given. */

func BetweenLimits(from, to rune) string {
	
	var result string
	var runeValue rune
	
	if from > to {
		for i := to + 1; i < from; i++ {
			runeValue = rune(i)
			result += string(runeValue) // appends a rune to a string
		} 
		} else if from < to {
			for i := from + 1; i < to; i++ {
			runeValue = rune(i)
			result += string(runeValue)
		} 
		}

	return result
}

/*

// Swap if 'from' is greater than 'to'
	if from > to {
		from, to = to, from
	}

	var result string

	// Loop through the runes between 'from' and 'to'
	for i := from + 1; i < to; i++ {
		result += string(i)
	}

	return result
}

*/