package Struct

import "fmt"

/**
二叉树
	1。定义
		1、每个结点最多有两颗子树，结点的度最大为2。
		2、左子树和右子树是有顺序的，次序不能颠倒。
		3、即使某结点只有一个子树，也要区分左右子树。

	2。先序 中序  后序 遍历

	3。实现

	4。应用
		  在处理大批量的动态的数据是比较有用。

	5。优缺点
		  数组的搜索比较方便，可以直接用下标，但删除或者插入某些元素就比较麻烦。
		  链表与之相反，删除和插入元素很快，但查找很慢。
		  二叉排序树就既有链表的好处，也有数组的好处。

		缺点：
			顺序存储可能会浪费空间(在非完全二叉树的时候)，但是读取某个指定的节点的时候效率比较高O(0)
			链式存储相对二叉树比较大的时候浪费空间较少，但是读取某个指定节点的时候效率偏低O(nlogn)

	6。复杂度
          二叉树支持动态的插入和查找，保证操作在O(height)时间
 */

type TreeNode struct {
	IsRoot bool
	Data   int
	Left   *TreeNode
	Right  *TreeNode
}

func NewTree() *TreeNode {
	return &TreeNode{
		IsRoot: true,
	}
}

func (n *TreeNode) insert(data int) {
	if n.IsRoot == true {
		n.IsRoot = false
		n.Data = data
	} else {
		newNode := &TreeNode{
			IsRoot: false,
			Data:   data,
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

//先序遍历 （递归法） （根-左-右）
func preOrder(t *TreeNode) {
	fmt.Printf("先序遍历----%v\n", t.Data)
	if t.Left != nil {
		preOrder(t.Left)
	}
	if t.Right != nil {
		preOrder(t.Right)
	}
}

//先序遍历（非递归，利用栈结构）
func preOrderTwo(t *TreeNode) {
	//todo
}

//中序遍历（左-根-右）
func midOrder(t *TreeNode) {
	if t.Left != nil {
		midOrder(t.Left)
	}
	fmt.Printf("中序遍历-----%v\n", t.Data)
	if t.Right != nil {
		midOrder(t.Right)
	}
}

//后序遍历 (左-右-根)
func postOrder(t *TreeNode) {
	if t.Left != nil {
		postOrder(t.Left)
	}
	if t.Right != nil {
		postOrder(t.Right)
	}
	fmt.Printf("后序遍历-----%v\n", t.Data)
}

func main() {
	node := NewTree()
	node.insert(8)
	node.insert(7)
	node.insert(12)
	node.insert(9)
	node.insert(15)
	fmt.Printf("root----%+v-----\n", (*node).Data)
	left1 := (*node).Left
	right1 := (*node).Right
	fmt.Printf("left1----%+v-----\n", left1.Data)
	fmt.Printf("right1----%+v-----\n", right1.Data)
	fmt.Printf("left2----%+v-----\n", right1.Left.Data)
	fmt.Printf("right2----%+v-----\n", right1.Right.Data)

	preOrder(node)
	fmt.Printf("------------\n")
	midOrder(node)
	fmt.Printf("------------\n")
	postOrder(node)
	fmt.Printf("------------\n")
}
