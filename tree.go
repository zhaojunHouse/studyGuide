package main

import "fmt"

/**
二叉树
	1。定义
	2。先序 中序  后序
	3。实现二叉树
	4。优缺点，应用。复杂度
 */

type TreeNode struct {
	IsRoot   bool
	Data  int
	Left  *TreeNode
	Right *TreeNode
}


func NewTree() *TreeNode{
	return &TreeNode{
		IsRoot:true,
	}
}


func (n *TreeNode) insert(data int){
	if n.IsRoot == true {
		n.IsRoot = false
		n.Data = data
	} else {
		newNode := &TreeNode{
			IsRoot:false,
			Data: data,
		}
		current := n
		for {
			parent := current
			if data < parent.Data {
				current = parent.Left
				if current == nil {
					parent.Left = newNode
					return
				}
			} else {
				current = parent.Right
				if current == nil {
					parent.Right = newNode
					return
				}
			}
		}
	}
}


func main() {
	node := NewTree()
	node.insert(8)
	node.insert(7)
	node.insert(12)
	node.insert(9)
	node.insert(15)
	fmt.Printf("root----%+v-----\n",(*node).Data)
	left1 := (*node).Left
	right1 := (*node).Right
	fmt.Printf("left1----%+v-----\n",left1.Data)
	fmt.Printf("right1----%+v-----\n",right1.Data)
	fmt.Printf("left2----%+v-----\n",right1.Left.Data)
	fmt.Printf("right2----%+v-----\n",right1.Right.Data)

}
