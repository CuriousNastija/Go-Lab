package sprint

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