package main

import "fmt"

/**
	链表实现
 */
type Node struct {
	Data interface{} `json:"data"`
	Next *Node
}

type List struct {
	Size     int   `json:"size"`
	FistNode *Node `json:"fist_node"`
	LastNode *Node `json:"last_node"`
}

func (list *List) Init() {
	(*list).Size = 0
	(*list).FistNode = nil
	(*list).LastNode = nil
}

func (list *List) add(node *Node) {
	if (*list).Size == 0 {
		(*list).Size = 1
		(*list).FistNode = node
		(*list).LastNode = node
	} else {
		(*list).Size++
		(*list).LastNode.Next = node
		(*list).LastNode = node
	}
}

func (list *List) Insert(node *Node, i int) bool {
	// 空的节点、索引超出范围和空链表都无法做插入操作
	if node == nil || i > (*list).Size || (*list).Size == 0 {
		return false
	}

	if i == 0 { // 直接排第一，也就领导小舅子才可以
		(*node).Next = (*list).FistNode
		(*list).FistNode = node
	} else {
		// 找到前一个元素
		preItem := (*list).FistNode
		for j := 1; j < i; j++ { // 数前面i个元素
			preItem = (*preItem).Next
		}
		// 原来元素放到新元素后面,新元素放到前一个元素后面
		(*node).Next = (*preItem).Next
		(*preItem).Next = preItem
	}

	(*list).Size++
	return true
}

func (list *List) Remove(i int, node *Node) bool {
	if i >= (*list).Size {
		return false
	}

	if i == 0 { // 删除头部
		node = (*list).FistNode
		(*list).FistNode = (*node).Next
		if (*list).Size == 1 { // 如果只有一个元素，那尾部也要调整
			(*list).LastNode = nil
		}
	} else {
		preItem := (*list).FistNode
		for j := 1; j < i; j++ {
			preItem = (*preItem).Next
		}

		node = (*preItem).Next
		(*preItem).Next = (*node).Next

		if i == ((*list).Size - 1) { // 若删除的尾部，尾部指针需要调整
			(*list).LastNode = preItem
		}
	}
	(*list).Size--
	return true
}



func (list *List) Get(i int) *Node {
	if i >= (*list).Size {
		return nil
	}

	item := (*list).FistNode
	for j := 0; j < i ; j++ {    // 从head数i个
		item = (*item).Next
	}
	fmt.Printf("item-----%+v",item)
	return item
}


func main(){
	lastNode := &Node{
		Data:"lastNode",
		Next:nil,
	}
	firstNode := &Node{
		Data:"firstNode",
		Next:lastNode,
	}
	list := &List{}
	list.add(lastNode)
	list.add(firstNode)
	list.Get(0)
	list.Get(1)
}