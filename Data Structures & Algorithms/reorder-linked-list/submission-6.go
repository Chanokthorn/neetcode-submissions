/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(n *ListNode) {
	prior, latter := cutHalf(n)
	reversedLatter := mutateReverse(latter)
	merge(prior, reversedLatter)
}

func cutHalf(n *ListNode) (prior, latter *ListNode) {
	var cutPoint, fast, slow *ListNode
	fast = n
	slow = n
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		cutPoint = slow
		slow = slow.Next
	}
	fmt.Println("signal 1")
	fmt.Printf("cut point: %d\n", cutPoint.Val)
	prior = n
	latter = cutPoint.Next
	cutPoint.Next = nil
	return prior, latter
}

func mutateReverse(n *ListNode) *ListNode{
	if n == nil {
		return nil
	}
	if n.Next == nil {
		return n
	}

	reversedHead := mutateReverse(n.Next)
	n.Next.Next = n
	n.Next = nil
	return reversedHead
}

func merge(n1, n2 *ListNode) {
	for n1 != nil && n2 != nil {
		if n1 != nil {
			prevNext := n1.Next
			n1.Next = n2
			n1 = prevNext
		}
		if n2 != nil {
			prevNext := n2.Next
			n2.Next = n1
			n2 = prevNext
		}
	}
}

func newLL(arr []int) *ListNode{
	var head, tail *ListNode
	for _, val := range arr {
		node := &ListNode{
			Val: val,
		}
		if head == nil {
			head = node
			tail = head
			continue
		}
		tail.Next = node
		tail = tail.Next
	}
	return head
}

func printLL(n *ListNode) {
	for n != nil {
		fmt.Printf("%d ", n.Val)
		n = n.Next
	}
	fmt.Println("")
}
