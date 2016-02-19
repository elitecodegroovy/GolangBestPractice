package main

import (
	"algo"
	"fmt"
	"strcon"
)

func DefineMap() {
	dict := make(map[string]string)
	dict["go"] = "Golang"
	dict["cs"] = "CSharp"
	dict["rb"] = "Ruby"
	dict["py"] = "Python"
	dict["js"] = "JavaScript"
	fmt.Print("\n")
	for k, v := range dict {
		fmt.Printf("Key: %s Value: %s\t", k, v)
	}
	fmt.Print("\n")
	if lan, ok := dict["go"]; ok {
		fmt.Println(lan, ok)
	}
}
func DefineSlice() {
	//x := make([]int, 5,10)
	//A Slice Initializes for a Specific Length Without Providing Elements
	//x2 := []int{4: 0}
	//append and copy
	x := []int{10, 20, 30}
	y := append(x, 40, 50)
	fmt.Println(x, y)

	x1 := []int{1, 2, 3, 4, 5}
	y1 := make([]int, 6)
	copy(y1, x1)
	fmt.Println(x1, y1)

	x2 := []int{10, 20, 30}
	for k, v := range x2 {
		fmt.Printf("Index: %d Value: %d\t ", k, v)
	}
}

func DefineArray() {
	//define array, default is zero.
	x1 := [5]int{0: 101, 2: 12, 4: 22}
	fmt.Println("define array:", x1)
}

func Init() {
	if algo.Fibonacci(6) == 8 {
		fmt.Println("fibonacci(6)== 8 ")
	}
	var _s = "Golang world, I are coming!"
	s := strcon.SwapCase(_s)
	fmt.Println("Converted string is :", s)
	DefineArray()
	DefineSlice()
	DefineMap()
}

func main() {
	Init()
}
