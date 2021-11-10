package maxdepth

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   12,
			Left:  &TreeNode{Val: 121},
			Right: &TreeNode{Val: 122},
		},
		Right: &TreeNode{
			Val:   13,
			Left:  &TreeNode{Val: 131},
			Right: &TreeNode{Val: 132},
		},
	}
	fmt.Println(maxDepth(root))
}
