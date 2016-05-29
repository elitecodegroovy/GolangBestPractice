package algo

//Fibonacci The Fibonacci numbers
//for n=1, 2, ... are 1, 1, 2, 3, 5, 8, 13, 21, ... (OEIS A000045).
func Fibonacci(x int) int {
	if x == 0 || x == 1 {
		return x
	} else {
		return Fibonacci(x-1) + Fibonacci(x-2)
	}
}
