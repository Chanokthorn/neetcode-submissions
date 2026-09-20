/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}	
	return f(head.Next, head)
}

func f(fast, slow *ListNode) bool {
	if fast == nil || fast.Next == nil {
		return false
	}
	if fast.Val == slow.Val {
		return true
	}
	return f(fast.Next.Next, slow.Next)
}
