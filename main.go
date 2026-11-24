package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// TopView returns, for every vertical column, the value seen from above —
// the FIRST node BFS meets in that column.
func TopView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	type item struct {
		n  *TreeNode
		hd int
	}
	first := map[int]int{}
	minHd, maxHd := 0, 0
	queue := []item{{root, 0}}
	for len(queue) > 0 {
		it := queue[0]
		queue = queue[1:]
		if _, seen := first[it.hd]; !seen {
			first[it.hd] = it.n.Val // write only on first sight: topmost wins
		}
		if it.hd < minHd {
			minHd = it.hd
		}
		if it.hd > maxHd {
			maxHd = it.hd
		}
		if it.n.Left != nil {
			queue = append(queue, item{it.n.Left, it.hd - 1})
		}
		if it.n.Right != nil {
			queue = append(queue, item{it.n.Right, it.hd + 1})
		}
	}
	out := make([]int, 0, maxHd-minHd+1)
	for hd := minHd; hd <= maxHd; hd++ {
		out = append(out, first[hd])
	}
	return out
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(1, tn(2, tn(4, nil, nil), tn(5, nil, nil)), tn(3, tn(6, nil, nil), tn(7, nil, nil)))
	fmt.Println(TopView(root))
}
