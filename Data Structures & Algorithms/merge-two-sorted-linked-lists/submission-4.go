/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var more, less *ListNode
	if list1.Val  > list2.Val {
		more = list1
		less = list2
	} else {
		more = list2
		less = list1
	}
	less.Next = mergeTwoLists(less.Next, more)
	return less
}
