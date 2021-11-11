package main

func main() {
	connect(&Node{Val: 1, Left: &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}}, Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}}})
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}

	queue := make([]*Node, 0)

	level := 0
	nextLevel := 1 // 2**0

	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[level]
		level++

		if level == nextLevel {

			for i := 0; i < nextLevel-1; i++ {
				queue[i].Next = queue[i+1]
			}
			queue = queue[nextLevel:]
			nextLevel = nextLevel * 2
			level = 0
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}

	return root

}
func connect2(root *Node) *Node {

	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			if i < levelSize-1 {
				currentNode.Next = queue[0]
			} else if i == levelSize-1 {
				currentNode.Next = nil
			}

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}

			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
	}
	return root
}
