package mymath

func MyPower(base int, exponent int) int {
	result := 1
	for i := 1; i <= exponent; i++ {
		result = result * base
	}
	return result
}

func myabs(number int) int { //절대값
	if number < 0 {
		return number * -1
	}
	return number
}
