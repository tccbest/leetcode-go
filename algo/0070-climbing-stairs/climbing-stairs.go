package _070_climbing_stairs

func ClimbStairs(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 1
	}

	a, b, sum := 1, 1, 2
	for i := 2; i < n; i++ {
		a = b
		b = sum
		sum = a + b
	}

	return sum
}