package isSymmetric

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isSymmetric(root))
}
