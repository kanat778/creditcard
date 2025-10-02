package main

func AlgorithmLuna(s string) bool {

	sum := 0
	double := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
		digit := int(s[i] - '0')

		if double {
			digit = digit * 2
			if digit > 9 {
				digit = digit - 9
			}
		}
		sum = sum + digit
		double = !double
	}
	return sum%10 == 0
}
