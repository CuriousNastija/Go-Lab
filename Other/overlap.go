package sprint


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