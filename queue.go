package main

/**
	队列 ：先进先出。
	slice实现数据结构 队列
 */
import "fmt"

type Element interface{}

type Queue interface {
	Offer(e Element)    //向队列中添加元素
	Poll()   Element    //移除队列中最前面的元素
	Clear()  bool       //清空队列
	Size()  int            //获取队列的元素个数
	IsEmpty() bool   //判断队列是否是空
}

type  sliceEntry struct{
	element []Element
}

func NewQueue() *sliceEntry {
	return &sliceEntry{}
}


//向队列中添加元素
func (entry *sliceEntry) Offer(e Element) {
	fmt.Println("add----",e)
	entry.element = append(entry.element,e)
}

//移除队列中最前面的额元素
func (entry *sliceEntry) Poll() Element {
	if entry.IsEmpty() {
		fmt.Println("queue is empty!")
		return nil
	}

	firstElement := entry.element[0]
	entry.element = entry.element[1:]
	fmt.Println("pop-----",firstElement)
	fmt.Printf("queue----%+v\n",entry.element)
	return firstElement
}

func (entry *sliceEntry) Clear() bool {
	if entry.IsEmpty() {
		fmt.Println("queue is empty!")
		return false
	}
	for i:=0 ; i< entry.Size() ; i++ {
		entry.element[i] = nil
	}
	entry.element = nil
	return true
}

func (entry *sliceEntry) Size() int {
	return len(entry.element)
}

func (entry *sliceEntry) IsEmpty() bool {
	if len(entry.element) == 0 {
		return true
	}
	return false
}

func queueHelper(q Queue,size int){
	for i:= 0;i< size;i++{
		q.Offer(i)
	}
	for i:= 0 ;i< size ;i++{
		q.Poll()
	}
}

func main() {
	queue := NewQueue()
	queueHelper(queue,10)
}