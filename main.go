package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
/* Create a Go function that takes two integers, min and max, as input.
The function should return a slice of integers containing all values between min
(inclusive) and max (exclusive) using the make function to create the slice.
If min is greater than or equal to max, the function should return a nil slice. */

/* func GenerateRange(min, max int) []int {

	if min >= max {
		return nil // nil slice
	}

	sequence := make([]int, 0) // int slice, length 0

	for i := min; i < max; i++ {
		sequence = append(sequence, i)
	}

	return sequence

} */

/* Write a Go function that takes an array of float64, two indices,
and removes elements that lie between these indices from the array.
The lower index should be removed, and the upper index should be kept.
The function should return the resulting modified array. The indices can be
negative or larger than the length of the array, but the function should
still remove the elements that fall within the specified range.
The indices might also be in wrong order. */

/* func RemoveElementsInRange(arr []float64, from, to int) []float64 {

	arrLen := len(arr)

	// if from is negative
	if from < 0 {
		from = arrLen + from
	}

	if from < 0 { // if from goes out of bounds on the left
		from = 0
	} else if from >= arrLen { // if from goes out of bounds on the right
		from = arrLen - 1
	}

	// if to is negative
	if to < 0 {
		to = arrLen + to
	}

	if to < 0 { // if to goes out of bounds on the left
		to = 0
	} else if to >= arrLen { // if to goes out of bounds on the right
		to = arrLen - 1
	}

	// if order is wrong
	if from > to {
		from, to = to, from
	}

	// if the same nothing to remove
	if from == to {
		return arr
	}

	// concatenate elements before from :from and after to to:
	arr = append(arr[:from], arr[to:] ...)

	return arr

} */

/* Write a Go function that takes an array of booleans and adds the minimum number
of booleans necessary so that the count of true and false values in the array
is equal. The function should return the resulting modified array.
The order of the elements must remain the same and new elements should
be added at the end of the array. */

/* func BalanceOut(arr []bool) []bool {

	var countTrue int
	var countFalse int

	for _, b := range arr {
		if b == true {
			countTrue++
		} else {
			countFalse++
		}
	}

	if countTrue > countFalse {
		difference := countTrue - countFalse
		for difference > 0 {
			arr = append(arr, false)
			difference -= 1
		}
	} else if countTrue < countFalse {
		difference := countFalse - countTrue
		for difference > 0 {
			arr = append(arr, true)
			difference -= 1
		}
	} else {
		return arr
	}

	return arr

} */

/* Write a Go function that takes a 2D array of integers and an integer value.
The function should filter out all subarrays from the array (2D) in which
the sum of elements is lower than the given value. The resulting modified 2D array should be returned. */

/* func FilterBySum(arr [][]int, limit int) [][]int {

	result := [][]int{} //new slice to store new elements

	for i := 0; i < len(arr); i++ { // iterate through subarrays
		sum := 0
		for j := 0; j < len(arr[i]); j++ { // iterate through elements of subarray
			sum += arr[i][j]
		}

		if sum >= limit {
			result = append(result, arr[i]) // add suitable subarrays into new array
		}

	}

	return result

} */

// Write a function that sorts a slice of integers in ascending order.

/* func SortIntegerTable(table []int) []int {

	tableLen := len(table)

	for i := 0; i < tableLen - 1; i++{
		for j := i+1; j < tableLen; j++ {
			if table[j] < table[i] {
				table[j], table[i] = table[i], table[j]
			}
		}
	}

	return table

} */

/* Write a Go function that takes an array of strings, applies
the StrToInt function you wrote earlier to every element in the array,
and returns the resulting integer values as a new array. Note that you
cannot call the StrToInt function from the package in our current
learning environment, you have to copypaste the function instead. */

/* func BulkAtoi(arr []string) []int {

	intArr := []int{}

	for _, s := range arr {
		intArr = append(intArr, StrToInt(s))
	}

	return intArr

}

func StrToInt (s string) int {

	if len(s) == 0 {
		return 0
	}

	isNegative := false

	if s[0] == '-' {
		isNegative = true
		s = s[1:]
	}

	result := 0

	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		digit := int(s[i] - '0') //s[i] represented in ASCII
		result = result*10 + digit
	}

	if isNegative {
		result = -result
	}

	return result
} */

/* Dive into the world of number combinations!
Your goal is to create a function that reveals all
possible combinations of ascending digits of length n (see example). */

/*func CombN(n int) []string {

	result := []string{}

	combination := make([]int, n) // this will define our combination size

	for i := 0; i < n; i++ { // write 0, 1, 2, n-1 numbers into this new slice
		combination[i] = i
	}

	done := false
	for !done {
		currentCombination := ""
		for i := 0; i < n; i++ {
			currentCombination += fmt.Sprintf("%d", combination[i])
		}
		result = append(result, currentCombination)

		done = true
		for i := n-1; i >= 0; i-- {
			if combination[i] < 9-(n-1-i) {
				combination[i]++
				for j := i+1; j<n; j++ {
					combination[j] = combination[j-1]+1
				}
				done = false
				break
			}
		}
	}

return result
}*/

/* Create a Go function that takes a non-empty string
as input and returns the first character of the string as a rune. */

/* func GetFirstRune(s string) rune {

	runesS := []rune(s)

	firstRune := runesS[0]

	return firstRune

} */

/* Create a Go function that takes a non-empty
string as input and returns the last character of the string as a rune */

/* func GetLastRune(s string) rune {

	runesS := []rune(s)

	sLen := len(s)

	lastRune := runesS[sLen-1]

	return lastRune

} */

/* Create a Go function that takes a non-empty string as
its first argument and an index as its second argument.
The function should return the rune at the specified index in the string. */

func NRune(s string, i int) rune {

	runesS := []rune(s)

	iRune := runesS[i]

	return iRune

}

/* Create a Go function that accepts a string 
and returns two integers. The first integer represents 
the number of runes in the string, and the second integer 
represents the byte length of the string. The function 
should handle UTF-8 encoding for all characters. */

/* func StrLength(s string) []int {

	runesS := []rune(s)

	var result []int

	result = append(result, len(runesS), len(s))

	return result

}*/

/* Create a Go function that takes a string and returns the reversed version of that string. */

/* func StrReverse(s string) string {

	runesS := []rune(s)

	for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 { //i starts at the befinning and j starts at the end
		runesS[i], runesS[j] = runesS[j], runesS[i]
	}
	return string(runesS)

} */

/* Create a Go function that takes a string as its parameter. 
The function should return true if the string contains only 
lowercase characters, and false otherwise. */

func IsLower(s string) bool {

	runesS := []rune(s)

	for i := 0; i < len(runesS); i++ {
		if runesS[i] < 'a' || runesS[i] > 'z' {
			return false
		}
	}

	return true

}

/* Create a Go function that takes a string as its parameter. 
The function should return true if the string contains only numeric 
characters, and false otherwise. */

func IsNumeric(s string) bool {

	runesS := []rune(s)

	for i := 0; i < len(runesS); i++ {
		if runesS[i] < '0' || runesS[i] > '9' {
			return false
		}
	}

	return true
}

/* Create a Go function that takes a string as its parameter 
and returns a new string with each letter capitalized. */

func ToUpperCase(s string) string {

	runesS := []rune(s)
	var result string

	for i := 0; i < len(runesS); i++ {
		if runesS[i] >= 'a' && runesS[i] <= 'z' {
			result += string(runesS[i] - ('a' - 'A'))
		} else {
			result += string(runesS[i])
		}
	}

	return result
}

/* Create a Go function that takes a string as its parameter. 
The function should capitalize the first letter of each word 
while converting the rest of the word to lowercase.
In this task a word is defined as a sequence of alphanumeric characters. */

func ToCapitalCase(s string) string {

	runesS := []rune(s)

	result := ""

	newWord := true

	for i := 0; i < len(runesS); i++ {
		if runesS[i] >= 'a' && runesS[i] <= 'z' {
			if newWord {
				result += string(runesS[i] - 32)
				newWord = false
			} else {
				result += string(runesS[i])
			}
		} else {
			result += string(runesS[i])
			if (runesS[i] < 'A' || runesS[i] > 'Z') && (runesS[i] < '0' || runesS[i] > '9') {
				newWord = true
			} else {
				newWord = false
			}
		}
	
}
return result
}

/* Create a Go function that takes a slice of strings 
and a separator string as its parameters. 
The function should return a new string by concatenating 
all the strings in the slice, with each string separated by the provided separator.*/

func StrConcatWith(strs []string, sep string) string {

	result := ""

	for i := 0; i < len(strs); i++ {
		result += strs[i]
		if i < len(strs) - 1 {
			result += sep
		}
	}

	return result

}

/* Create a Go function that takes a string as its parameter. 
The function should split the string into words and store them 
in a string slice. Words are separated by spaces, tabs, and newlines. */

func SplitWhitespaces(s string) []string {

	result := []string{}
	word := ""

	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(r)
		}

	}

	if word != "" {
		result = append(result, word)
	}
	
	return result

}

/* Create a Go function that takes a string s and a 
separator sep as parameters. The function should split 
the string s into substrings using the separator sep and 
return a slice of strings containing the resulting substrings. */

func StrSplitBy(s, sep string) []string {

	result := []string{}
	substring := ""
	sepLen := len(sep)

	for i := 0; i < len(s); {
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			result = append(result, substring)
			substring = ""

			i += sepLen
		} else {
			substring += string(s[i]) 
			i++
		}

	}

	if substring != "" {
		result = append(result, substring)
	}

	return result

}

/* Create a Go function that behaves like the Index function.
It takes two strings as input: s and toFind. The function should 
find the index of the first occurrence of toFind in s and return 
that index as an integer. If toFind is not present in s, the function should return -1. */

func SubstrIndex(s string, toFind string) int {

	slen := len(s)
	toFindLen := len(toFind)

	if toFindLen == 0 {
		return 0
	}


	for i := 0; i <= slen - toFindLen; i++ {
		if s[i:i+toFindLen] == toFind {
			return i
		} 
	}

	return -1

}

/* Create a Go function that mimics the behavior of the Compare function. 
It takes two strings, a and b, as input and returns an integer.
The function should compare the two strings and return:

0 if the strings are equal,
-1 if a comes before b in lexicographic order,
1 if a comes after b in lexicographic order. */

func StrCompare(a, b string) int {

	if a == b {
		return 0
	}

	if a < b {
		return - 1
	} 
		
	return 1
	

}

/* Create a Go function that takes an integer n and 
a string base as parameters. The function should return 
the integer n represented in the specified base as a string.
Here are the validity rules for the base:

The base must contain at least 2 unique characters.
The base should not contain the characters + or -.
The function should handle negative numbers as well 
(see examples in the usage section). If the base is not valid, the function should return "NV" (Not Valid). */

func NbrBase(n int, base string) string {

	runesBase := []rune(base)
	baseLen := len(runesBase)

	if baseLen < 2 {
		return "NV"
	}

	for i := 0; i < baseLen; i++ {

		if runesBase[i] == '+' || runesBase[i] == '-' {
			return "NV"
		}
		
		for j := i + 1; j < baseLen; j++ {
			if runesBase[i] == runesBase[j] {
				return "NV"
			}
		}
	}

	if n == 0 {
		return string(runesBase[0])
	}

	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	var digits []rune
	for n > 0 {
		digit := n % baseLen
		digits = append(digits, runesBase[digit])
		n /= baseLen
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i + 1, j - 1 {
			digits[i], digits[j] = digits[j], digits[i]
	}
	
	result := string(digits)


	if isNegative {
		result = "-" + result
	}
	return result
	
}

/* Create a Go function that takes two parameters:

s: a numeric string in a specific base.
base: a string containing unique characters representing the available digits in that base.
The function should return the integer value of s in the given base. If the base is not valid, it returns 0.
Here are the validity rules for the base:
The base must contain at least 2 unique characters.
The base should not contain the characters + or -.
The function only works with valid string numbers that 
consist of elements present in the base. It does not handle negative numbers. */

func StrToDecimalBase(s, base string) int {
	runesBase := []rune(base)
	baseLen := len(runesBase)

	// Validate the base: must have at least two characters and no + or -
	if baseLen < 2 {
		return 0
	}

	for i := 0; i < baseLen; i++ {
		if runesBase[i] == '+' || runesBase[i] == '-' {
			return 0
		}
		for j := i + 1; j < baseLen; j++ {
			if runesBase[i] == runesBase[j] {
				return 0 // Base contains duplicate characters
			}
		}
	}

	// Convert the string `s` to its decimal integer value
	result := 0
	for _, r := range s {
		// Find the character's index in the base
		value := -1
		for i := 0; i < baseLen; i++ {
			if r == runesBase[i] {
				value = i
				break
			}
		}

		// If the character is not found in the base, return 0
		if value == -1 {
			return 0
		}

		// Update the result: result = result * baseLen + value
		result = result*baseLen + value
	}

	return result
}


/* You're tasked with creating an iterative function that 
calculates the factorial of an integer passed as a parameter. 
Make sure to handle errors, returning 0 for non-possible values or overflows. */

func FactorialIterative(n int) int {

	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	var factorial int = 1
	for i := 1; i <= n; i++ {
		if factorial > (1<<63-1)/i {
			return 0
		}
		factorial *= i
	}

	return factorial

}

/* You're tasked with creating a recursive function that 
calculates the factorial of an integer passed as a parameter. 
Make sure to handle errors, returning 0 for non-possible values 
or overflows. NB! You should achieve this without using a for loop.*/

func FactorialRecursive(n int) int {

	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	return n * FactorialRecursive(n-1)

}

/* You're tasked with creating an iterative function 
that calculates the power of an integer n to the given power. 
Handle negative powers by returning 0. */

func ToThePowerIterative(n int, power int) int {

	if power < 0 {
		return 0
	}

	if n == 0 {
		if power == 0 {
			return 1
		}
		return 0
	}

	if power == 0 {
		return 1
	}

	result := 1
	for i := 0; i < power; i++ {
		result *= n
	}

	return result

}

/* You're tasked with creating an recursive function 
that calculates the power of an integer n to the given power. 
Handle negative powers by returning 0. NB! You should achieve this without using a for loop. */

func ToThePowerRecursive(n int, power int) int {

	if power < 0 {
		return 0
	}

	if n == 0 {
		if power == 0 {
			return 1
		}
		return 0
	}

	if power == 0 {
		return 1
	}

	return n * ToThePowerRecursive(n, power-1)

}

/* You're tasked with creating a recursive function that returns the 
value at a given position index in the Fibonacci sequence. 
The first value is at index 0.
The Fibonacci sequence starts with 0 and 1, and each subsequent 
value is the sum of the two preceding values: 0, 1, 1, 2, 3, etc. 
Negative indices should be handled by returning -1. */

func NthFibonacci(index int) int {

	if index < 0 {
		return -1
	}

	if index == 0 {
		return 0
	}

	if index == 1 || index == 2 {
		return 1
	}

	return NthFibonacci(index-1) + NthFibonacci(index-2)

}

/* You're tasked with creating a function that takes an integer 
as a parameter and returns true if this number is prime and false if it's not.
You should optimize the function to avoid time-outs during testing.
Remember that prime numbers can be only natural numbers greater than 1. */

func IsPrime(n int) bool {

	if n <= 1 {
		return false
	}

	if n == 1 {
		return true
	}

	if n % 2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n % i == 0 {
			return false
		}
	}

	return true

}

/* You're tasked with creating a function that takes an integer as a parameter. 
If the given integer is a prime, the function should return that integer; 
otherwise, it should find and return the next prime number.
You should optimize the function to avoid time-outs during testing.
Remember that prime numbers can be only natural numbers greater than 1. */

func NextPrime(n int) int {

	if n <= 1 {
		return 2
	}

	current := n
	for {
		if prime(current){
		return current
	}
	current++
}
}

	func prime(num int) bool {

	if num <= 1 {
        return false
    }
    if num == 2 {
        return true  // 2 is the only even prime number.
    }
    if num % 2 == 0 {
        return false  // Exclude even numbers.
    }
    for i := 3; i*i <= num; i += 2 {
        if num % i == 0 {
            return false
        }
    }
    return true
}

/* You're tasked with creating a function that takes an integer 
as a parameter and returns its square root if the square 
root is a whole number; otherwise, it should return 0.*/

func Sqrt(n int) int {

	if n < 0 {
		return 0
	}
	
	if n == 0 || n == 1 {
		return n
	}

	for i := 2; i*i <= n; i++ {
		if i*i == n {
			return i
		}
	}

	return 0

}

/* You're tasked with creating a function that returns the 
solutions to the eight queens puzzle. Use recursivity to solve this problem.
The function returns the solutions as a string, 
with each solution on a single line. The index of the placement 
of a queen starts at 1 and reads from left to right, with each digit 
representing the position for each column. The solutions should be in ascending order. */

func isSafe(board [] int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board [i] == col || abs(board[i]-col) == abs(i-row) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func placeQuenns(board []int, row int, solutions *[]string) {
	if row == len(board) {
		solution := make([]string, len(board))
		for i, col := range board {
			solution[i] = strconv.Itoa(col + 1)
		}
		*solutions = append(*solutions, strings.Join(solution, ""))
		return
	}

	for col := 0; col < len(board); col++ {
		if isSafe(board, row, col) {
			board[row] = col
			placeQuenns(board, row+1, solutions)
		}
	}
}


func EightQueensSolver() string {
	board := make([]int, 8)
	var solutions []string
	placeQuenns(board, 0, &solutions)
	return strings.Join(solutions, "\n")

}

/* You're tasked with creating a function ArrAny that 
returns true for a string slice if at least one element 
returns true when passed through the f function. */


func IsUpper(s string) bool {

	if len(s) == 0 {
		return false
	}

	for _, e := range s {
		if e < 'A' || e > 'Z' {
		return false
		} 
	}
		return true

}

func IsAlphanumeric(s string) bool {

	if len(s) == 0 {
        return false
    }

    for _, e := range s {
        if !(e >= '0' && e <= '9') && !(e >= 'A' && e <= 'Z') && !(e >= 'a' && e <= 'z') {
            return false
        }
    }
    return true

}

func ArrAny(f func(string) bool, a []string) bool {

	for _, str := range a {
		if f(str) {
			return true
		}
	}

	return false 

}

/* Write a Go function that takes two integers as 
input and returns their greatest common divisor (GCD). 
The GCD is the largest positive integer that divides both of 
the input integers without leaving a remainder. */

func GCD(a, b int) int {

	if b == 0 {
		return a
	}

	tmp := a
	a = b
	b = tmp % a
	return GCD(a, b)

}

/* Write a Go function that takes two integers as input and 
returns their least common multiple (LCM). The LCM is the 
smallest positive integer that is divisible by both of the input integers. */

func LCM(a, b int) int {

	var result int

	result = (a * b)/GCD(a, b)

	return result

}

/* Write a Go function that takes a string as input and returns a 
boolean value indicating whether the input string is a palindrome or 
not. A palindrome is a string that reads the same forwards and backward, 
considering character case and white spaces. */

func IsPalindrome(s string) bool {


	for i := 0; i < len(s)/2; i++ {
		for j := len(s)-1; j >= 0; j-- {
			if s[i] != s[len(s)-i-1] {
				return false
			}
		}
	}

	return true

}

/* Write a Go function that takes two strings as input and returns 
a boolean value indicating whether the input strings are anagrams or not. 
Anagrams are words or phrases formed by rearranging the letters of another, 
and in this task, you should ignore differences in character case and whitespace. */

func AreAnagrams(str1, str2 string) bool {

	a := CleanAndSortString(str1)
	b := CleanAndSortString(str2)

	return a == b

}

func CleanAndSortString(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}

/* You're tasked with creating a function called SortWordArr 
that sorts a string slice in ascending order based on ASCII values. */

func SortWordArr(a []string) []string {

	n := len(a)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
		}

	return a

}

/* Write a Go function that takes a string as input and compresses 
it by replacing consecutive repeating characters with a count of 
repetitions followed by the character. However, if there is only 
one character in a sequence, it should not be prepended with a number. 
For example, "aaabbaac" should be compressed to "3a2b2ac", but "abcde" should remain unchanged as "abcde". */

func StrCompress(input string) string {

	count := 1
	result := ""

	for i := 0; i < len(input); i++ {
		if i < len(input)-1 && input[i] == input [i+1] {
			count++
		} else {
			if count > 1 {
				result += fmt.Sprintf("%d", count)
			}
			result += string(input[i])
			count = 1
			
		}
	}

	return result

}

/* Write a Go function that takes an array of 
integers as input and returns the same array with 
duplicate elements removed, preserving the order of 
the remaining elements. For example, if the input array is 
[3, 5, 2, 3, 8, 5, 2], the function should return [3, 5, 2, 8]. */

func RemoveDuplicates(arr []int) []int {

	result := []int{}

	for i := 0; i < len(arr); i++ {
		found := false

		for j := 0; j < len(result); j++ {
			if arr[i] == result[j] {
				found = true
				break
			}
		}

		if !found {
			result = append(result, arr[i])
		}
	}

	return result

}

/* Write a Go function that takes an array of integers and returns 
the longest contiguous subarray in the array in which the elements are increasing. 
The function should return the contiguous subarray. */

func LongestClimb(arr []int) []int {

	if len(arr) == 0 {
		return []int{}
	}

	longest := []int{}
	current :=[]int{arr[0]}

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			current = append(current, arr[i])
		} else {
			if len(current) > len(longest) {
				longest = current
			}

			current = []int{arr[i]}
		}
	}

	if len(current) > len(longest) {
		longest = current
	}

	return longest
	

}

/* Write a Go function that takes an integer amount and 
a slice of integers denominations. The function should 
use the values in denominations to pay exactly the amount. 
The denominations can be used any number of times, but 
higher denominations should be preferred. The function 
should return the resulting denominations as an array of 
integers in descending order. If the payout cannot be made, return an empty array. */

func Payout(amount int, denominations []int) (payout []int) {

	for amount > 0 {
		largestDenom := -1
	

	for _, denom := range denominations {
		if denom <= amount && (largestDenom == -1 || denom > largestDenom) {
			largestDenom = denom
		}
	}

	if largestDenom == -1 {
		return []int{}
	}
	
	payout = append(payout, largestDenom)
	amount -= largestDenom

}
	return payout

}

/* Write a Go function that takes a valid Roman numeral 
as input and converts it into an integer. The function 
should return the integer representation of the valid Roman numeral input. */

func GetValueFromRoman(ch rune) int {
    switch ch {
    case 'I':
        return 1
    case 'V':
        return 5
    case 'X':
        return 10
    case 'L':
        return 50
    case 'C':
        return 100
    case 'D':
        return 500
    case 'M':
        return 1000
    default:
        return 0
    }
}

func FromRoman(s string) int {

	var result int

	for i := 0; i < len(s); i++ {
		current := GetValueFromRoman(rune(s[i]))

		if i < len(s)-1 && current < GetValueFromRoman(rune(s[i+1])) {
			result -= current
		} else {
			result += current
		}
	}
	
	return result

}

/* Write a Go function that takes an integer and converts 
it into a Roman numeral. The function should return "Invalid input" 
if the input integer is less than 1 or more than 3999. Otherwise, 
it should return the Roman numeral representation of the valid input integer. */

func ToRoman(num int) string {

	if num < 1 || num > 3999 {
		return "Invalid input"
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			result += symbols[i]
			num -= values[i]
		}
	}

	return result
}

/* Write a Go function that takes an integer as input and returns a 2D 
integer array representing Pascal's triangle up to the specified number 
of levels. Pascal's triangle is a triangular array of binomial coefficients 
where each number is the sum of the two directly above it. The function should 
generate the specified number of levels of Pascal's triangle. */

func PascalsTriangle(n int) [][]int {

	triangle := [][]int{}

	if n == 0 {
		return triangle
	}

	for i := 0; i < n; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1

		for j := 1; j < i; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		triangle = append(triangle, row)
	}

	return triangle

}

/* Write a Go function that takes two strings as 
input and finds the longest common substring that occurs 
in both strings. The function should return the substring 
that occurs earlier in the first string passed if there 
are multiple substrings of the same length. */

func LongestCommonSubstr(str1, str2 string) string {

	substr := ""

	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			k := 0
			for i+k < len(str1) && j+k < len(str2) && str1[i+k] == str2[j+k] {
				k++
			}
			if k > len(substr) {
				substr = str1[i : i+k]
			}
			} 
		}
		return substr
	}

/* Write a Go function that takes a string containing various 
characters, including parentheses ()[]{}, and checks if the parentheses 
are balanced. The function should return a boolean value indicating 
whether the parentheses are balanced or not. */

func BalancedParentheses(input string) bool {

	stack := []rune{}

	for _, ch := range input {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else if ch == ')' || ch == ']' || ch == '}' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack) - 1]
			stack = stack[:len(stack)-1]

			if (ch == ')' && top != '(') || (ch == ']' && top != '[') || (ch == '}' && top != '{') {
				return false
			}
		}
	}

	return len(stack) == 0

}

/* Write a Go function that takes two arrays of integers 
and returns an array containing the elements that are 
common in both input arrays, sorted in ascending order. 
If an element occurs multiple times in both arrays, 
it should be included in the result array as many times as it occurs. */

func Overlap(arr1, arr2 []int) []int {

	result := []int{}
	used := make([]bool, len(arr2))
	
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] && !used[j] {
				result = append(result, arr1[i])
				used[j] = true
				break
			}
		}
	}

	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result

}

/* Write a Go function that takes a matrix 
represented as a 2D array and returns the transposed matrix. 
Transposing a matrix involves swapping its rows and columns. 
The function should take the original matrix and return the resulting transposed 2D array. */

func TransposeMatrix(matrix [][]int) [][]int {

	rows := len(matrix)
	cols := len(matrix[0])

	transpose := make([][]int, cols)

	for i := range transpose {
		transpose[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transpose[j][i] = matrix[i][j]
		}
	}

	return transpose

}

/* Write a Go function that solves the digital root problem. 
The function should take an integer as input and return an integer. 
The digital root is the recursive sum of the digits of a number until 
a single-digit result is achieved. For example, the digital root of 9875 
is 2 because 9 + 8 + 7 + 5 = 29, and 2 + 9 = 11, and finally 1 + 1 = 2. */

func DigitalRoot(n int) int {

	for n > 9 {
		n = sumOfDigits(n)
	}
	return n
}

func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}


func main() {

	fmt.Println(DigitalRoot(123456))
	//fmt.Println(TransposeMatrix([][]int{{1, 2, 3}, {4, 5, 6}}))
	//arr1 := []int{1, 2, 2, 3, 4, 5}
	//arr2 := []int{2, 3, 4, 4, 5, 6}
	//fmt.Println(Overlap(arr1, arr2))
	//fmt.Println(BalancedParentheses("Everything { is [ fine ()]}()"))
	//fmt.Println(LongestCommonSubstr("ABCBDAB", "BDCAB"))
	//fmt.Println(PascalsTriangle(2))
	//fmt.Println(ToRoman(128))
	//fmt.Println(FromRoman("CXXVIII"))
	//amount := 128
	//denominations := []int{1, 2, 5, 10, 20, 50, 100, 200}
	//fmt.Println(Payout(amount, denominations))
	//fmt.Println(LongestClimb([]int{8, 4, 2, 1, 2, 4, 8, 2, 4, 8}))
	//fmt.Println(RemoveDuplicates([]int{1, 2, 3, 2, 4, 8, 8, 1, 2, 0, 8}))
	//fmt.Println(StrCompress("abac"))
	//a := []string{"a", "F", "8", "b", "D", "9", "c", "E", "7"}
	//fmt.Println(SortWordArr(a))
	//fmt.Println(AreAnagrams("Listen", "S i l e n t"))
	//fmt.Println(IsPalindrome("level"))
	//fmt.Println(LCM(12, 18))
	//s := "HowYOUhaveYOUyouYOUbeen?"
	//fmt.Println(StrSplitBy(s, "YOU"))
	//fmt.Println(SplitWhitespaces("Hello! How have you been?"))
	//toConcat := []string{"Three", " Two", " One", " Go!"}
	//fmt.Println(StrConcatWith(toConcat, "."))
	//fmt.Println(ToCapitalCase("Hello! Great to see you! How-are-you-doing-2day?"))
	//fmt.Println(ToUpperCase("Hello! How's your day going?"))
	//fmt.Println(IsNumeric("01+23-45"))
	//fmt.Println(IsLower("hello"))
	//fmt.Println(StrReverse("Hello Coder!"))
	//fmt.Println(StrLength("Hello World!"))
	//fmt.Println(string(NRune("Coding is cool", 4)))
	//fmt.Println(string(GetLastRune("kood")))
	//fmt.Println(GetFirstRune("kood"))
	//fmt.Println(BulkAtoi([]string{"8", "kood", "-13"}))
	//fmt.Println(SortIntegerTable([]int{2, 0, 5, 4, 1, 3}))
	//fmt.Println(FilterBySum([][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, 9))
	//fmt.Println(BalanceOut([]bool{true, false, false, false}))
	//fmt.Println(RemoveElementsInRange([]float64{10., .8, -.4, 20., 7.7, 3.}, 4, 1))
	//fmt.Println(GenerateRange(-10, 9))
	
	
	
} 