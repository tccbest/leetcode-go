package _402_remove_k_digits

import (
	"fmt"
	"testing"
)

func TestRemoveKdigits(t *testing.T) {
	fmt.Println(RemoveKdigits("1432219", 3))
	fmt.Println(RemoveKdigits("10200", 1))
	fmt.Println(RemoveKdigits("10", 2))
}
