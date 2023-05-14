package calc

func Add(args ...int) int {
	var sum int = 0
	for _, v := range args {
		sum += v
	}
	return sum
}

func Sub(a, b int) int {
	return a - b
}
