package _0_1_fei_bo_na_qi_shu_lie_lcof

func Fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	a, b, sum := 0, 1, 1

	for i := 2; i < n; i++ {
		a = b
		b = sum
		sum = (a + b) % 1000000007
	}

	return sum
}
