package goutil

// Max Return maximum value between 2 number
func Max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

// Min Return minimum value between 2 number
func Min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}

// Abs Return absolute value
func Abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

// Pow Return power value of base and exponent
//
// If exponent < 0, return 0
func Pow(base, exponent int) int {
	if exponent < 0 {
		return 0
	}
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	return base * Pow(base, exponent-1)
}
