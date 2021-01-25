package _020_valid_parentheses

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	fmt.Println(IsValid("()"))	//true
	fmt.Println(IsValid("()[]{}"))	//true
	fmt.Println(IsValid("(]"))	//false
	fmt.Println(IsValid("([)]"))	//false
	fmt.Println(IsValid("{[]}"))	//true
}
