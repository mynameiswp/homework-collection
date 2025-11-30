package main

import (
	"fmt"
	"math"
)

// 1. 定义Shape接口：包含Area()和Perimeter()方法
type Shape interface {
	Area() float64      // 计算面积
	Perimeter() float64 // 计算周长
}

// 2. 定义Rectangle结构体（矩形）
type Rectangle struct {
	Width  float64 // 宽度
	Height float64 // 高度
}

// 实现Shape接口的Area()方法（矩形面积=长×宽）
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 实现Shape接口的Perimeter()方法（矩形周长=2×(长+宽)）
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 3. 定义Circle结构体（圆形）
type Circle struct {
	Radius float64 // 半径
}

// 实现Shape接口的Area()方法（圆形面积=πr²）
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 实现Shape接口的Perimeter()方法（圆形周长=2πr）
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// 创建Rectangle实例
	rect := Rectangle{Width: 5, Height: 3}
	// 创建Circle实例
	circle := Circle{Radius: 4}

	// 调用方法并输出（直接通过结构体实例调用）
	fmt.Println("=== 矩形信息 ===")
	fmt.Printf("宽度：%.2f，高度：%.2f\n", rect.Width, rect.Height)
	fmt.Printf("面积：%.2f\n", rect.Area())
	fmt.Printf("周长：%.2f\n", rect.Perimeter())

	fmt.Println("\n=== 圆形信息 ===")
	fmt.Printf("半径：%.2f\n", circle.Radius)
	fmt.Printf("面积：%.2f\n", circle.Area())
	fmt.Printf("周长：%.2f\n", circle.Perimeter())

	// （拓展）利用接口实现多态：用Shape接口类型存储不同结构体实例
	var shape Shape
	shape = rect
	fmt.Println("\n通过Shape接口调用矩形面积：", shape.Area())

	shape = circle
	fmt.Println("通过Shape接口调用圆形周长：", shape.Perimeter())
}
