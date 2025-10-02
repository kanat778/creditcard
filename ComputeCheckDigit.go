package main

func ComputeCheckDigit(s string) int {
	checkDigit := 0
	sum := 0
	double := true

	for i := 0; i < len(s); i++ {
		digit := int(s[i] - '0')
		if double {
			sum = digit * 2
			if digit > 9 {
				digit = digit - 9
			}
		}
		sum = sum + digit
		double = !double
	}
	checkDigit = (10 - (sum % 10)) % 10
	return checkDigit
}
