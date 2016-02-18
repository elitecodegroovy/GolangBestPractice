package main

import (
	"algo"
	"fmt"
	"strcon"
)

func main() {
	if algo.Fibonacci(6) == 8 {
		fmt.Println("fibonacci(6)== 8 ")
	}
	var _s = "Golang world, I are coming!"
	s := strcon.SwapCase(_s)
	fmt.Println("Converted string is :", s)
	fmt.Printf("%s", _s)
}
