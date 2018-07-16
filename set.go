package main

import (
	"fmt"
	"sort"
	"sync"
)

/**
  集合： 由map实现
  [很多函数需要修改才能使用，只完成了并集操作]
//todo  map为什么无序     slice为什么有序     锁     hashset   哈希表
 */

type Set struct {
	sync.RWMutex
	m map[int]interface{}
}

// 新建集合对象
func New(items ...interface{}) *Set {
	s := &Set{
		m: make(map[int]interface{}, len(items)),
	}
	s.Add(items...)
	return s
}

// 添加元素
func (s *Set) Add(items ...interface{}) {
	s.Lock()
	defer s.Unlock()
	count := len(s.m)
	for k, v := range items {
		s.m[k+count+1] = v
	}
}

// 删除元素
func (s *Set) Remove(items ...interface{}) {

}

// 判断元素是否存在
func (s *Set) Has(items ...interface{}) int {
	s.RLock()
	defer s.RUnlock()
	for k, v := range items {
		for _,val := range s.m {
			if v == val {
				return k
			}
		}
	}
	return -1
}

// 元素个数
func (s *Set) Count() int {
	return len(s.m)
}

// 清空集合
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]interface{}{}
}

// 空集合判断
func (s *Set) Empty() bool {
	return len(s.m) == 0
}

// 无序列表
func (s *Set) List() []interface{} {
	s.RLock()
	defer s.RUnlock()
	list := make([]interface{}, 0, len(s.m))
	for _,item := range s.m {
		list = append(list, item)
	}
	return list
}

// 排序列表
func (s *Set) SortList() []string {
	s.RLock()
	defer s.RUnlock()
	list := make([]string, 0, len(s.m))
	for _,item := range s.m {
		list = append(list, item.(string))
	}
	sort.Strings(list)
	return list
}

// 并集
func (s *Set) Union(sets ...*Set) *Set {
	r := New(s.List()...)
	for _, set := range sets {
		for _,v := range set.m {
			if s.Has(v) == -1 {
				r.Add(v)
			}
		}
	}
	return r
}

// 差集
//func (s *Set) Minus(sets ...*Set) *Set {
//
//}
//

// 交集
//func (s *Set) Intersect(sets ...*Set) *Set {
//
//}
//

// 补集
//func (s *Set) Complement(full *Set) *Set {
//
//}

func main() {
	s1 := New("a","g","fe","d","6")
	s2 := New("a","b","c","af","3")
	//s3 := New(5, 6, 9, 10)
	r1 := s1.Union(s2)     // 获取并集
	//r2 := s1.Minus(s2, s3)     // 获取差集
	//r3 := s1.Intersect(s2, s3) // 获取交集
	//r4 := s3.Complement(s1)    // 获取 s3 相对于 s1 的补集
	//fmt.Println(s1.SortList())
	fmt.Println(s1)
	fmt.Println(r1.List())
	fmt.Println(r1.SortList())


	//fmt.Println(r2.SortList())
	//fmt.Println(r3.SortList())
	//fmt.Println(r4.SortList())
}
