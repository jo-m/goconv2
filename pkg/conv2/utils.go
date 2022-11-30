package conv2

// modulo which also works for negative numbers
func wrap(m, n int) int {
	return ((m % n) + n) % n
}

func center(x int) int {
	return (x - 1) / 2
}
