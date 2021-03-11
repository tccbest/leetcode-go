package _0_01_sorted_merge_lcci

//双指针逆向
func merge(A []int, m int, B []int, n int)  {
	if len(B) == 0 {
		return
	}

	allIndex := m + n - 1
	aIndex, bIndex := m-1, n-1
	for aIndex >= 0 || bIndex >= 0 {
		if aIndex == -1 {
			A[allIndex] = B[bIndex]
			bIndex--
		} else if bIndex == -1 {
			A[allIndex] = A[aIndex]
			aIndex--
		} else if A[aIndex] < B[bIndex] {
			A[allIndex] = B[bIndex]
			bIndex--
		} else {
			A[allIndex] = A[aIndex]
			aIndex--
		}
		allIndex--
	}
}
