package sprint

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
