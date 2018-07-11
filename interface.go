package main

import (
	"fmt"
	"math"
)

/**
	接口示例.实现接口所有的方法，就实现一个接口
 */

type Shape interface {
	area() float64
}


type Square struct {
	a float64
	b float64
}

type Circle struct {
	r float64
}

func (s *Square)area() float64{
	return s.a*s.b
}

func (c Circle)area() float64{
	return math.Pi * c.r *c.r
}

func getArea(shape Shape) float64{
	return shape.area()
}

func main(){
	square := &Square{a:2,b:3}
	area := getArea(square)
	fmt.Println(area)
}