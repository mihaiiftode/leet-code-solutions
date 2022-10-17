package main

func main() {
	reorderList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}})
}

func reorderList(head *ListNode) {

	array := []*ListNode{}
	array = append(array, head)

	for start := head.Next; start != nil; start = start.Next {
		array = append(array, start)
	}
	n := len(array)

	for i := 0; i < n/2; i++ {

		array[i].Next = array[n-i-1]
		array[n-i-1].Next = array[i+1]

	}

	array[n/2].Next = nil

}

type ListNode struct {
	Val  int
	Next *ListNode
}
