package main


import (
	"errors"
	"fmt"
)


type Stack struct {
	Element []interface{} //Element
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack)Push(value ...interface{}){
	stack.Element = append(stack.Element,value...)
}

//返回下一个元素
func (stack *Stack)Top()(value interface{}){
	if stack.Size() > 0 {
		return stack.Element[stack.Size() - 1]
	}
	return nil //read empty stack
}

//返回下一个元素,并从Stack移除元素
func (stack *Stack)Pop()(err error){
	if stack.Size()> 0 {
		stack.Element = stack.Element[:stack.Size() - 1]
		return nil
	}
	return errors.New("Stack为空.") //read empty stack
}

//交换值
func (stack *Stack)Swap(other *Stack){
	switch{
	case stack.Size() == 0 && other.Size() == 0:
		return
	case other.Size() == 0 :
		other.Element = stack.Element[:stack.Size()]
		stack.Element = nil
	case stack.Size()== 0 :
		stack.Element = other.Element
		other.Element = nil
	default:
		stack.Element,other.Element = other.Element,stack.Element
	}
	return
}

//修改指定索引的元素
func (stack *Stack)Set(idx int,value interface{})(err error){
	if idx >= 0 && stack.Size() > 0 && stack.Size() > idx{
		stack.Element[idx] = value
		return nil
	}
	return errors.New("Set失败!")
}

//返回指定索引的元素
func (stack *Stack)Get(idx int)(value interface{}){
	if idx >= 0 && stack.Size() > 0 && stack.Size() > idx {
		return stack.Element[idx]
	}
	return nil //read empty stack
}

//Stack的size
func (stack *Stack)Size()(int){
	return len(stack.Element)
}

//是否为空
func (stack *Stack)Empty()(bool){
	if stack.Element == nil || stack.Size() == 0 {
		return true
	}
	return false
}

//打印
func (stack *Stack)Print(){
	for i := len(stack.Element) - 1; i >= 0; i--{
		fmt.Println(i,"=>",stack.Element[i])
	}
}
