package main

func main() {
	middleNode(&ListNode{Val: 1, Next: &ListNode{Val: 2}})
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cache := map[int]*ListNode{}
	count := 0
	for head != nil {
		cache[count] = head
		head = head.Next
		count++
	}

	return cache[count/2]

}
