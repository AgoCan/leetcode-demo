package stackDemo

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	fmt.Println("isEmpty", s.IsEmpty())
	s.Push("a")
	s.Push("b")
	s.Push("c")
	fmt.Println("pop: ", s.Pop())
	fmt.Println("size:", s.Size())
	fmt.Println("peek:", s.Peek())
	fmt.Println("size:", s.Size())
	fmt.Println("isEmpty", s.IsEmpty())
}
