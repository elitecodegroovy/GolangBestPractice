package util

import (
	"fmt"
	"testing"
)

func TestSwapCase(t *testing.T) {
	input, expected := "Hello, World", "hELLO, wORLD"
	result := SwapCase(input)
	if result != expected {
		t.Errorf("SwapCase(%q) == %q, expected %q", input, result, expected)
	}
}

func TestReverse(t *testing.T) {
	input, expected := "Hello, World", "dlroW ,olleH"
	result := Reverse(input)
	if result != expected {
		t.Errorf("Reverse (%q) == %q, expected %q", input, result, expected)
	}
}

// -------------------benchmark ------------------------------------
//Benchmark for SwapCase function
func BenchmarkSwapCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SwapCase("Hello, World")
	}
}

//Benchmark for Reverse function
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("Hello, World")
	}
}

// -------------------verifyting ------------------------------------
//Example code for Reverse function
func ExampleReverse() {
	fmt.Println(Reverse("Hello, World"))
	// Output: dlroW ,olleH
}

//Example code for Reverse function
func ExampleSwapCase() {
	fmt.Println(SwapCase("Hello, World"))
	// Output: hELLO, wORLD
}
