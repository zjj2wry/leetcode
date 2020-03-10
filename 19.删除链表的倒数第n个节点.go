package main

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 *
 * https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (36.51%)
 * Likes:    605
 * Dislikes: 0
 * Total Accepted:    96.5K
 * Total Submissions: 264.2K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
 *
 * 示例：
 *
 * 给定一个链表: 1->2->3->4->5, 和 n = 2.
 *
 * 当删除了倒数第二个节点后，链表变为 1->2->3->5.
 *
 *
 * 说明：
 *
 * 给定的 n 保证是有效的。
 *
 * 进阶：
 *
 * 你能尝试使用一趟扫描实现吗？
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
// 1->2->3
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	first := dummy
	second := dummy

	for i := 0; i <= n; i++ {
		second = second.Next
	}

	for second != nil {
		first = first.Next
		second = second.Next
	}

	first.Next = first.Next.Next

	return dummy.Next
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
	list = removeNthFromEnd(list, 2)
	list.Display()
}
*/
// @lc code=end
