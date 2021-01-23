package _070_climbing_stairs

func ClimbStairs(n int) int {
	a, b, sum := 0, 0, 1
	for i := 0; i < n; i++ {
		a = b
		b = sum
		sum = a + b
	}

	return sum
}