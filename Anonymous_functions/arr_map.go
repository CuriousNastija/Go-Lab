package sprint

func IsNegative(n int) bool {
	if n >= 0 {
		return false
	} else {
		return true
	}
}

func IsPrime(n int) bool {

	if n <= 1 {
        return false
    }

    if n == 2 {  
        return true
    }

    if n % 2 == 0 {  
        return false
    }

    // Check for factors starting from 3
    for i := 3; i*i <= n; i += 2 {
        if n % i == 0 {
            return false
        }
    }

    return true

}

func ArrMap(f func(int) bool, a []int) []bool {

	result := make([]bool, len(a))
	for i, e := range a {
		result[i] = f(e)

	}
	return result

}