package sprint

func StrLength(s string) []int {

	var byteLength int
	var runeLength int
	resultArr := []int{}


	byteLength = len(s)
	runeLength = len([]rune(s))

	resultArr = append(resultArr, runeLength, byteLength)

	return resultArr
	
}