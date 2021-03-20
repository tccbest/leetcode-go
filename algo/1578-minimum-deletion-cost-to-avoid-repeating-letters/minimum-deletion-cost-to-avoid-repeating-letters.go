package _578_minimum_deletion_cost_to_avoid_repeating_letters

func MinCost(s string, cost []int) int {
	ret, left, right := 0, 0, 1
	for right < len(cost) {
		if s[left] == s[right] {
			if cost[left] <= cost[right] {
				ret += cost[left]
				left = right
				right++
			} else {
				ret += cost[right]
				right++
			}
		} else {
			left = right
			right++
		}
	}

	return ret
}