package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]
	algo := AlgorithmLuna(args)
	fmt.Println(algo)

	if AlgorithmLuna(args) {
		fmt.Println("OK!")
	}
}
