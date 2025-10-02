package main

import (
	"fmt"
	"strings"
)

func Generate(s string) []string {
	//res := ""
	nStars := strings.Count(s, "*")

	if nStars != 4 {
		fmt.Println("Error: mask contains no asterisks or too many asterisks")
		return nil
	}

	firstStar := strings.Index(s, "*")

	if firstStar == -1 || firstStar != len(s)-nStars {
		fmt.Println("Error: mask contains no asterisks or too many asterisks")
		return nil
	}

	prefix := s[:firstStar]

	for i := 0; i < len(prefix); i++ {
		if prefix[i] < '0' || prefix[i] > '9' {
			fmt.Println("Error: prefix must contain only digits!")
			return nil
		}
	}
	fmt.Println(prefix)

	results := []string{}

	// перебор 000..999
	for i := 0; i < 1000; i++ {
		mid := fmt.Sprintf("%03d", i)
		candidate := prefix + mid
		results = append(results, candidate)
	}

	return results
	//return res
}
