package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Adds the values of the two nodes
func sum(t1 *TreeNode, t2 *TreeNode) int {
	val := 0

	if t1 != nil {
		val += t1.Val
	}

	if t2 != nil {
		val += t2.Val
	}

	return val
}

func merge(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	// If both nodes are nil, then the result is a nil node.
	if t1 == nil && t2 == nil {
		return nil
	}

	val := sum(t1, t2)

	// Create a new node with the combined value
	newNode := TreeNode{val, nil, nil}

	// Since the Left and Right of the trees are also trees,
	// You can apply the merge operation on them and assign the merged result
	// to the left and right of this parent tree respectively.
	//
	// So we’ll basically merge two branches recursively.

	// If the tree is not nil, take its left and right nodes
	// otherwise assume a nil Tree
	var t1Left, t1Right, t2Left, t2Right *TreeNode
	if t1 != nil {
		t1Left = t1.Left
		t1Right = t1.Right
	}
	if t2 != nil {
		t2Left = t2.Left
		t2Right = t2.Right
	}

	// Recursively merge the Left subtrees:
	newNode.Left = merge(t1Left, t2Left)

	// Recursively merge the right subtrees:
	newNode.Right = merge(t1Right, t2Right)

	// Note that the recursion ends when there is nothing to merge
	// (i.e. when both of the tree nodes are nil;
	// Because merging a node with a nil node, or merging two non-nil nodes will
	// give a new TreeNode that needs you to `sum` them, whereas merging two nil nodes
	// will only return a nil node that you can return right away w/o any computation.

	// The function returns a pointer type, so we take the address of the newNode.
	return &newNode
}

func main() {
	treeA := TreeNode{21, nil, nil}
	treeB := TreeNode{21, nil, nil}

	fmt.Println(merge(&treeA, &treeB))
}
