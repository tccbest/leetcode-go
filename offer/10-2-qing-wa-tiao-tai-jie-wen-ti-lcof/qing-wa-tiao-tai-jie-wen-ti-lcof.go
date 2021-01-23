package _0_2_qing_wa_tiao_tai_jie_wen_ti_lcof

func NumWays(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 1
	}

	a, b, sum := 0, 1, 1
	for i := 1; i < n; i++ {
		a = b
		b = sum
		sum = (a + b) % 1000000007
	}

	return sum
}
