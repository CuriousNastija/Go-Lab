package sprint

import "fmt"

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