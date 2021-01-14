package _0_zui_xiao_de_kge_shu_lcof

func GetLeastNumbers(arr []int, k int) []int {
	al := len(arr)
	if al <= k {
		return arr
	}

	return quickSelect(arr, 0, al-1, k)
}

func quickSelect(arr []int, l, r, k int) []int {
	i := partition(arr, l, r)
	if i == k {
		return arr[:k]
	} else if i < k {
		return quickSelect(arr, i+1, r, k)
	} else {
		return quickSelect(arr, l, i-1, k)
	}
}

func partition(arr []int, l, r int) int {
	pivot := arr[r]

	i := l
	for j := l; j < r; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[r] = arr[r], arr[i]

	return i
}
