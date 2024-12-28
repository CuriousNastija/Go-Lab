package sprint

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