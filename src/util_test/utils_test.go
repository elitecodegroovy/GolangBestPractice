package util_test

import (
	"fmt"
	"testing"
	"time"

	. "util"
)

// Test case for the SwapCase function to execute in parallel
func TestSwapCaseInParallel(t *testing.T) {
	t.Parallel()
	// Delaying 1 second for the sake of demonstration
	time.Sleep(1 * time.Second)
	input, expected := "Hello, World", "hELLO, wORLD"
	result := SwapCase(input)
	if result != expected {
		t.Errorf("SwapCase(%q) == %q, expected %q", input, result, expected)
	}
}

// Test case for the Reverse function to execute in parallel
func TestReverseInParallel(t *testing.T) {
	t.Parallel()
	// Delaying 2 seconds for the sake of demonstration
	time.Sleep(2 * time.Second)
	input, expected := "Hello, World", "dlroW ,olleH"
	result := Reverse(input)
	if result != expected {
		t.Errorf("Reverse(%q) == %q, expected %q", input, result, expected)
	}
}

// Test case for the SwapCase function
func TestSwapCase(t *testing.T) {
	input, expected := "Hello, World", "hELLO, wORLD"
	result := SwapCase(input)
	if result != expected {
		t.Errorf("SwapCase(%q) == %q, expected %q", input, result, expected)
	}
}

// Test case for the Reverse function
func TestReverse(t *testing.T) {
	input, expected := "Hello, World", "dlroW ,olleH"
	result := Reverse(input)
	if result != expected {
		t.Errorf("Reverse(%q) == %q, expected %q", input, result, expected)
	}
}

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

func TestLongRun(t *testing.T) {
	// Checks whether the short flag is provided
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}
	// Long running implementation goes here
	time.Sleep(5 * time.Second)
}

func TestSubString(t *testing.T) {

	str := "ReadAtLeast reads from r into buf until it has read at least min bytes. It returns the number of bytes copied and an error if fewer bytes were read. The error is EOF only if no bytes were read. If an EOF happens after reading fewer than min bytes, ReadAtLeast returns ErrUnexpectedEOF. If min is greater than the length of buf, ReadAtLeast returns ErrShortBuffer. On return, n >= min if and only if err == nil."
	expected := "ReadAtLeast"
	result := Substring(str, 11)
	if result != expected {
		t.Errorf("Substring result: %q, expected: %q", result, expected)
	}
	expected = "ReadAtLeast reads"
	result = Substring(str, len(expected))
	if result != expected {
		t.Errorf("Substring result: %q, expected: %q", result, expected)
	}

	str = "上海新政：外来者购房条件由社保2年调为5年"
	expected = "上海新政"
	//TODO ...
	result = Substring(str, 4)
	fmt.Println("Geting上海新政 is ", result == "上海新政", result)
}
