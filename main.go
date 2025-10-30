package main

import (
	"creditcard/data"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	input, err := data.ParseCommandLine(args)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%+v\n", input)

}

//func AlgorithmLuna(s string) bool {
//	sum := 0
//	double := false
//
//	for i := len(s) - 1; i >= 0; i-- {
//		if s[i] < '0' || s[i] > '9' {
//			return false
//		}
//		digit := int(s[i] - '0')
//
//		if double {
//			digit = digit * 2
//			if digit > 9 {
//				digit = digit - 9
//			}
//		}
//		sum = sum + digit
//		double = !double
//	}
//	return sum%10 == 0
//}
