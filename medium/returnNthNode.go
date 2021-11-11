package main

func main() {
	removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	// the 2nd pointer will advance only after
	// we processed the first n nodes
	first, second := dummy, dummy

	var count int
	for first.Next != nil {
		first = first.Next

		count++
		if count > n {
			second = second.Next
		}
	}

	second.Next = second.Next.Next

	return dummy.Next
}

// func removeNthFromEnd(head *ListNode, n int) *ListNode {

// 	start := head
// 	var lastNode *ListNode
// 	count := 0
// 	for head != nil {
// 		head = head.Next

// 		if lastNode != nil {
// 			lastNode = lastNode.Next
// 		}

// 		if count == n {
// 			lastNode = start
// 		}
// 		count++

// 	}

// 	if lastNode == nil && count == n {
// 		return start.Next
// 	} else if lastNode == nil {
// 		return nil
// 	}

// 	if lastNode.Next.Next == nil {
// 		lastNode.Next = nil
// 	} else {
// 		lastNode.Next = lastNode.Next.Next
// 	}

// 	return start
// }
