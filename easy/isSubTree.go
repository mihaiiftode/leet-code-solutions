package main

import "fmt"

func main() {
	// isSubtree(&TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 5}}, &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}})
	result := isSubtree(&TreeNode{Val: 1, Left: &TreeNode{Val: 1}}, &TreeNode{Val: 1})
	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		return false
	}
	if p.Val == q.Val {
		if isEqual(p, q) {
			return true
		}
	}

	return isSubtree(p.Left, q) || isSubtree(p.Right, q)
}

func isEqual(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isEqual(p.Left, q.Left) && isEqual(p.Right, q.Right)
}
