// +build OMIT
package main

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * https://leetcode-cn.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (33.23%)
 * Likes:    1485
 * Dislikes: 0
 * Total Accepted:    228.6K
 * Total Submissions: 687.9K
 * Testcase Example:  '123'
 *
 * 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 *
 * 示例 1:
 *
 * 输入: 123
 * 输出: 321
 *
 *
 * 示例 2:
 *
 * 输入: -123
 * 输出: -321
 *
 *
 * 示例 3:
 *
 * 输入: 120
 * 输出: 21
 *
 *
 * 注意:
 *
 * 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回
 * 0。
 *
 */

// @lc code=start
func reverse(x int) int {
	rev := 0
	for ; x != 0; x = x / 10 {
		rev = rev*10 + (x % 10)
	}
	if rev > math.MaxInt32 || rev < math.MinInt32 {
		rev = 0
	}
	return rev
}

// @lc code=end
