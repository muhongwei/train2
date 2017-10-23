package main

import (
	"stack"
	"fmt"
)
func main(){
	var stack stack.Stack
	stack.Initstack()
	fmt.Println("----执行入栈：0~9--------------------------")
	for i:=0;i<10;i++{
		stack.Push(i)
	}
	fmt.Println("------------------------------------------")
	fmt.Println("------------------------------------------")
	fmt.Println("栈的长度length()",stack.Length())
	fmt.Println("栈顶元素topValue()值为：",stack.TopValue())
	fmt.Println("------------------------------------------")
	fmt.Println("-----执行出栈次数：12----------------------")
	for i:=12;i>0;i--{
		stack.Pop()
	}
}