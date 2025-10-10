package main

import (
	"creditcard/data"
	"fmt"
)

func main() {
	args := data.ParseArgs()
	fmt.Println(args)
}
