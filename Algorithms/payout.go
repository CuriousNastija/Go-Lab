package sprint

func Payout(amount int, denominations []int) (payout []int) {

	for amount > 0 {
		largestDenom := -1
	

	for _, denom := range denominations {
		if denom <= amount && (largestDenom == -1 || denom > largestDenom) {
			largestDenom = denom
		}
	}

	if largestDenom == -1 {
		return []int{}
	}
	
	payout = append(payout, largestDenom)
	amount -= largestDenom

}
	return payout

}