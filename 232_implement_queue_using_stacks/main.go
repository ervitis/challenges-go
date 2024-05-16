package main

import "fmt"

type MyQueue struct {
	data []int
}

func Constructor() MyQueue {
	return MyQueue{
		data: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	// [1, 2, 3]
	// push(4) -> [4, 1, 2, 3]
	this.data = append([]int{x}, this.data...)
}

func (this *MyQueue) Pop() int {
	if len(this.data) == 0 {
		return -1
	}
	x := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return x
}

func (this *MyQueue) Empty() bool {
	return len(this.data) == 0
}

func (this *MyQueue) Peek() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data)-1]
}

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}
