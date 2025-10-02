package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//args := os.Args[1]
	if len(os.Args) < 2 {
		fmt.Println("ERROR: Please enter one or more arguments!")
		return
	}
	for _, args := range os.Args[1:] {
		arg := strings.ReplaceAll(args, " ", "")
		generate := Generate(arg)
		fmt.Println(generate)

		if AlgorithmLuna(arg) {
			fmt.Printf("%s -> OK\n", arg)
		} else {
			fmt.Printf("%s -> INCORRECT\n", arg)
		}
	}

}
