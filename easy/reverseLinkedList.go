package main

import (
	"fmt"
)

func main() {
	head := reverseList(&ListNode{Val: 1,
		Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}})

	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var previous *ListNode = nil

	for head != nil {
		aux := head.Next
		head.Next = previous
		previous = head
		head = aux
	}

	return previous
}
