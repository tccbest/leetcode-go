package _125_valid_palindrome

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(IsPalindrome(s))

	s = "race a car"
	fmt.Println(IsPalindrome(s))
}
