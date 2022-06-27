/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
package main
type MyStack struct {
	vals []int
}

func Constructor() MyStack {
	return MyStack{
		vals: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.vals = append(this.vals, x)
}

func (this *MyStack) Pop() int {
	val := this.vals[len(this.vals)-1]
	this.vals = this.vals[0 : len(this.vals)-1]

	return val
}

func (this *MyStack) Top() int {
	return this.vals[len(this.vals)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.vals) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

