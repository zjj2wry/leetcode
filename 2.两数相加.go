// +build OMIT
package main

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (36.11%)
 * Likes:    3472
 * Dislikes: 0
 * Total Accepted:    266.2K
 * Total Submissions: 737.1K
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 *
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 *
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 * 示例：
 *
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	res := new(ListNode)
	cur := res
	carry := 0
	for p1 != nil || p2 != nil {
		i := 0
		j := 0
		if p1 != nil {
			i = p1.Val
		}
		if p2 != nil {
			j = p2.Val
		}
		sum := i + j + carry
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return res.Next
}

func arrayToListNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	head := new(ListNode)
	walk := head
	for index, val := range list {
		walk.Val = val
		if index != len(list)-1 {
			walk.Next = new(ListNode)
			walk = walk.Next
		}
	}
	return head
}

/*
func print(l1 *ListNode) {
	cur := l1
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := arrayToListNode([]int{2, 4, 3})
	print(l1)
	l2 := arrayToListNode([]int{5, 6, 4})
	print(l2)
	print(addTwoNumbers(l1, l2))
}
*/

// @lc code=end
