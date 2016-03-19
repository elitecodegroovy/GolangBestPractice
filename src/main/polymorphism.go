package main

import (
	"fmt"
)

type strCase struct{}

type intCase struct{}

func (text strCase) Add(x string, y string) string {
	return x + y
}

func (number intCase) Add(x int, y int) int {
	return x + y
}

func StartPolymorphism() {
	number := new(intCase)
	fmt.Print(" intCase Add: ", number.Add(1, 3), "\n")

	text := new(strCase)
	fmt.Print("strCase Add: ", text.Add("Add A", "Add B"), "\n")
}
