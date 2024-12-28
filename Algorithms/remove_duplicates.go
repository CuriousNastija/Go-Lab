package sprint

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