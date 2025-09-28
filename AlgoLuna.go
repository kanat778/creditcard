package main

import (
	"fmt"
	"strings"
)

func AlgorithmLuna(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	fmt.Println("Arr =", s)
	sum := 0
	double := false
	for i := len(s) - 1; i >= 0; i-- {
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
