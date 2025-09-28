package main

import (
	"fmt"
	"strings"
)

func AlgorithmLuna(arr string) bool {
	arr = strings.ReplaceAll(arr, " ", "")
	fmt.Println("Arr =", arr)
	sum := 0
	double := false
	for i := len(arr) - 1; i >= 0; i-- {
		//if arr[i] < '0' || arr[i] > '9' {
		//	continue
		//}
		digit := int(arr[i] - '0')

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
