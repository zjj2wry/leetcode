// +build OMIT
package main

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (32.06%)
 * Likes:    2754
 * Dislikes: 0
 * Total Accepted:    279.7K
 * Total Submissions: 872.2K
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	var max, start, end, length = 0, 0, 0, len(s)
	for start < length && end < length {
		if index := index(s[start:end], rune(s[end])); index < 0 {
			end = end + 1
			if end-start > max {
				max = end - start
			}
		} else {
			start = start + 1 + index
		}
	}
	return max
}

func index(s string, r rune) int {
	for index, value := range s {
		if value == r {
			return index
		}
	}
	return -1
}

// @lc code=end
