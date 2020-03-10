// +build OMIT
package main

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (65.75%)
 * Likes:    664
 * Dislikes: 0
 * Total Accepted:    125.8K
 * Total Submissions: 191.3K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 *
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
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

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}

/*
type ListNode struct {
	Val  int
	Next *ListNode
}

func (o *ListNode) Insert(val int) *ListNode {
	node := new(ListNode)
	node.Val = val
	if o == nil {
		o = node
		return o
	}
	node.Next = o
	o = node
	return o
}

func (o *ListNode) Display() {
	for o != nil {
		fmt.Print(o.Val, " ")
		o = o.Next
	}
	fmt.Println()
}

func main() {
	var list *ListNode
	for i := 0; i < 10; i++ {
		list = list.Insert(i)
	}

	list.Display()
	list = reverseList(list)
	list.Display()
}
*/

// @lc code=end
