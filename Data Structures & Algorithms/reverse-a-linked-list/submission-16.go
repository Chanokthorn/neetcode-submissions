/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(n *ListNode) *ListNode {
	if n == nil {
		return nil
	}
	if n.Next == nil {
		return n
	}
	reversedHead := reverseList(n.Next)
	n.Next.Next = n
	n.Next = nil
	return reversedHead
}
