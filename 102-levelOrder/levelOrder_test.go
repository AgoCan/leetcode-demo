package levelOrder

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
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
	fmt.Println(levelOrder(root))
}
