package _003_longest_substring_without_repeating_characters

/**
3. 无重复字符的最长子串

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
示例 4:

输入: s = ""
输出: 0
 

提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
*/
func LengthOfLongestSubstring(s string) int {
	ret := 0
	length := len(s)
	if length == 0 {
		return ret
	}

	m := map[byte]int{} //记录是否存在
	left, right := 0, -1	//初始化左右指针，右指针从最左侧开始
	for left < length {
		//左指针往后移，删除第一个字符
		if left != 0 {
			delete(m, s[left - 1])
		}

		for right+1 < length && m[s[right+1]] == 0 {
			//指针右移
			m[s[right+1]]++
			right++
		}

		ret = max(ret, right-left+1)

		left++
	}

	return ret
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
