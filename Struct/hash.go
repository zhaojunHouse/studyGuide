package Struct

import "fmt"

const maxCapacity int = 100     //bucket的最大容量
const loadFactor float32 = 0.75 //负载因子

var nowCapacity int = 10 //当前容量

type Entry struct {
	key   int
	value string
	next  *Entry
}

type HashMap struct {
	size    int
	buckets []Entry
}

func CreateHashMap() *HashMap {
	hm := &HashMap{}
	hm.size = 0
	hm.buckets = make([]Entry, nowCapacity, maxCapacity)
	return hm
}

//获取key的hash
func getHash(k int) int {
	return k % nowCapacity
}

func (hm *HashMap) GetSize() int {
	return hm.size
}

//由key在hashMap中找到其Entry指针
func (hm *HashMap) GetEntry(k int) (*Entry, bool) {
	p := hm.buckets[getHash(k)].next
	for p != nil {
		//找到,返回
		if p.key == k {
			return p, true
		} else {
			p = p.next
		}
	}
	//此时p==nil，没找到
	return nil, false
}

//不判断是否达到负载因子,直接往hashmap中插entry
func (hm *HashMap) insert(e *Entry) {

	hash := getHash(e.key)
	//从表头中的指针开始遍历
	var p *Entry = &hm.buckets[hash]
	for p.next != nil {
		//如果找到相同的key，则覆盖其中的值,完成insert，return
		if p.next.key == e.key {
			p.next.value = e.value
			return
		} else {
			p = p.next
		}
	}
	//此时p指向尾部结点
	p.next = e
	hm.size++
}

//判断是否达到负载因子再insert
func (hm *HashMap) Put(k int, v string) {
	e := &Entry{k, v, nil}
	hm.insert(e)
	//达到负载因子且还能扩容时，扩容并迁移数据
	if float32(hm.size)/float32(nowCapacity) >= loadFactor && nowCapacity < maxCapacity {
		if 2*nowCapacity > maxCapacity {
			nowCapacity = maxCapacity
		} else {
			nowCapacity = 2 * nowCapacity
		}

		newHm := CreateHashMap()
		var index int
		for index = 0; index < len(hm.buckets); index++ {
			p := hm.buckets[index].next
			for p != nil {
				//临时保存p的next
				pNext := p.next
				p.next = nil
				newHm.insert(p)
				p = pNext
			}
		}

		var b1 int
		var b2 int = len(newHm.buckets)

		//把hm的切片扩容至newHm一样大
		for b1 = len(hm.buckets); b1 < b2; b1++ {
			hm.buckets = append(hm.buckets, Entry{})
		}
		//移回数据
		for b1 = 0; b1 < b2; b1++ {
			hm.buckets[b1].next = newHm.buckets[b1].next
			newHm.buckets[b1].next = nil
		}
	}
}

//删除指定Entry
func (hm *HashMap) DeleteEntry(e *Entry) {
	//注意这里不能是p:=hm.buckets[getHash(e.key)].next,因为表头存的指针不一定为nil
	p := &hm.buckets[getHash(e.key)]
	for p != nil {
		if p.next == e {
			fmt.Println("正在删除")
			p.next = p.next.next
			fmt.Println("删除成功")

			hm.size--
			return
		} else {
			p = p.next
		}
	}
	fmt.Println("删除失败")
}

//删除指定key的Entry
func (hm *HashMap) Delete(k int) bool {
	e, ok := hm.GetEntry(k)
	if ok {
		fmt.Println("有该Entry,其hash为：", getHash(e.key), "值为：", e)
		hm.DeleteEntry(e)
		return true
	} else {
		fmt.Println("没有该Entry，删除失败")
		return false
	}
}

//遍历hashMap
func (hm *HashMap) Traverse() {
	var index int
	for index = 0; index < nowCapacity; index++ {
		p := hm.buckets[index].next
		if p == nil {
			fmt.Println(index, ":")
		} else {
			fmt.Print(index, ":")
		}
		for p != nil {
			fmt.Print("----->", p)
			p = p.next
			if p == nil {
				fmt.Println()
			}
		}
	}
	fmt.Println()
}