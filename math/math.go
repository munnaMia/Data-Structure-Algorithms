package math

// Give a number factorial.
func Factorial (n int) int {
	if n == 1 {
		return n
	}
	return n * Factorial(n-1)
}