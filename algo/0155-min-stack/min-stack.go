package _155_min_stack

import "math"

type MinStack struct {
	min   []int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)

	min := this.GetMin()
	if min > x {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, min)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}

	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}

	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 0
	}

	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
