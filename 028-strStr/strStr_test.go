package strStrDemo

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	ret := strStr("hello", "ll")
	fmt.Println(ret)
}
