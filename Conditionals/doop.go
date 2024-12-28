package sprint

/* You're tasked with creating a function called Doop. 
This function takes three parameters:

A value (integer).
An operator, which can be one of the following: +, -, /, *, %.
Another value (integer).
In case of an invalid operator, 
invalid values, or incorrect number of arguments, 
the function returns 0. 
The program must also handle modulo and division by 0.*/

var result int

func Doop(a int, op string, b int) int {
	
	var result int
	args := []interface{}{a, op, b}

	if len(args) != 3 {
		return 0
	} else if b == 0 {
		return 0
	} else {

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "/":
		result = a / b
	case "*":
		result = a * b
	case "%":
		result = a % b
	default:
		result = 0
	}
}
	return result
}

/*
	// Perform the operation based on the operator
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		return a / b
	case "*":
		return a * b
	case "%":
		return a % b
	default:
		return 0  // Invalid operator
	}
} /*