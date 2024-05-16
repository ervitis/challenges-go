package main

import "fmt"

type MyStack struct {
	data []int
}

func Constructor() MyStack {
	return MyStack{
		data: make([]int, 0),
	}
}
func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
}
func (this *MyStack) Pop() int {
	if len(this.data) == 0 {
		return -1
	}
	x := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return x
}

func (this *MyStack) Top() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.data) == 0
}

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}
