package sprint

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