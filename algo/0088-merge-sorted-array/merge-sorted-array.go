package _088_merge_sorted_array

func Merge(nums1 []int, m int, nums2 []int, n int)  {
	//初始化三个指针，分别指向nums1最后一个有效元素，nums2最后一个元素，nums1最后一个元素，
	p1, p2, p3 := m - 1, n -1, m + n - 1
	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p3] = nums1[p1]
			p1--
		} else {
			nums1[p3] = nums2[p2]
			p2--
		}
		p3--
	}
}