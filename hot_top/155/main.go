package main

import "math"

//MinStack() 初始化堆栈对象。
//void push(int val) 将元素val推入堆栈。
//void pop() 删除堆栈顶部的元素。
//int top() 获取堆栈顶部的元素。
//int getMin() 获取堆栈中的最小元素。

// 结构体
type MinStack struct {
	Stack   []int // 栈空间
	MinArry []int // 辅助栈
}

func Constructor() MinStack {
	return MinStack{
		Stack:   []int{},
		MinArry: []int{math.MaxInt64}, // 重点，此位置需要放一个最大值，防止第一次入栈时，取不到值
	}
}

func (this *MinStack) Push(val int) {

	this.Stack = append(this.Stack, val)
	top := this.MinArry[len(this.MinArry)-1]
	if val < top {
		this.MinArry = append(this.MinArry, val)
	} else {
		this.MinArry = append(this.MinArry, top)
	}
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.MinArry = this.MinArry[:len(this.MinArry)-1]
}

func (this *MinStack) Top() int {
	val := this.Stack[len(this.Stack)-1]
	//this.Pop()  // 获取后，并不弹出栈最上方的值
	return val
}

func (this *MinStack) GetMin() int {
	return this.MinArry[len(this.Stack)]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
