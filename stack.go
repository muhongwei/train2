package stack

import (
	"fmt"
)
const(
	ERROR = -666
)
type Stack struct{
	size int
	top int
	array []int
}

// type Stacker interface{
// 	Push(data Element)
// 	Pop() Element
// 	Length() Element
// 	TopValue() Element
// }
func (stack *Stack)Initstack(){
	stack.top=-1
	stack.array=make([]int,0)
}

func (stack *Stack)Push(data int){
	stack.array=append(stack.array,data)
	stack.top +=1
	fmt.Println("入栈：",stack.array[stack.top])
	
}
func (stack *Stack)Pop() bool{	
	if(stack.top==-1){
		fmt.Println("栈为空")
		return false
	}
	
	fmt.Println("出栈：",stack.array[stack.top])
	stack.top -=1
	return true
}
func (stack *Stack)Length() int{
	len:=stack.top+1
	return len
}
func (stack *Stack)TopValue() int{
	if(stack.top<0){
		fmt.Println("栈为空")
		return ERROR
	}
	tValue:=stack.array[stack.top]
	return tValue
}
