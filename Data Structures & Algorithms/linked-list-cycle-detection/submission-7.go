/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {	
	return f(head, head)
}

func f(fast, slow *ListNode) bool {
	if fast == nil || fast.Next == nil || fast.Next.Next == nil {
		return false
	}
	fast = fast.Next.Next
	slow = slow.Next
	if fast.Val == slow.Val {
		return true
	}
	return f(fast, slow)
}
