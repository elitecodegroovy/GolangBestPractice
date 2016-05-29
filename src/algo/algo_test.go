package algo

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var result int
	result = Fibonacci(6)
	if result != 8 {
		t.Error("Expected 8, got ", result)
	}
	fmt.Println(result)
}
