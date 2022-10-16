package method

import (
	"fmt"
	"math"
)

type Point7 struct{ X, Y float64 }

//这是给struct Point7类型定义一个方法
func (p Point7) Distance7(q Point7) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point7) Add(another Point7) Point7 {
	return Point7{p.X + another.X, p.Y + another.Y}
}

func (p Point7) Sub(another Point7) Point7 {
	return Point7{p.X - another.X, p.Y - another.Y}
}

func (p Point7) Print() {
	fmt.Printf("{%f, %f}\n", p.X, p.Y)
}

//定义一个Point7切片类型 Path
type Path []Point7

//方法的接收器 是Path类型数据, 方法的选择器是TranslateBy(Point7, bool)
func (path Path) TranslateBy(another Point7, add bool) {
	var op func(p, q Point7) Point7 //定义一个 op变量 类型是方法表达式 能够接收Add,和 Sub方法
	// if add == true {
	if add {
		op = Point7.Add //给op变量赋值为Add方法
	} else {
		op = Point7.Sub //给op变量赋值为Sub方法
	}

	for i := range path {
		//调用 path[i].Add(another) 或者 path[i].Sub(another)
		path[i] = op(path[i], another)
		path[i].Print()
	}
}

func MethodExample7() {

	Point7s := Path{
		{10, 10},
		{11, 11},
	}

	anotherPoint7 := Point7{5, 5}

	Point7s.TranslateBy(anotherPoint7, false)

	fmt.Println("------------------")

	Point7s.TranslateBy(anotherPoint7, true)
}

/*
{5.000000, 5.000000}
{6.000000, 6.000000}
------------------
{10.000000, 10.000000}
{11.000000, 11.000000}
*/
