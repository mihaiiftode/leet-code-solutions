package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummyHead := &ListNode{}
	curr := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			l1, l2 = l2, l1
		}
		curr.Next = l1
		curr, l1 = curr.Next, l1.Next
	}

	if l2 != nil {
		curr.Next = l2
	}

	return dummyHead.Next
}
