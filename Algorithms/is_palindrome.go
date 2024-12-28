package sprint

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