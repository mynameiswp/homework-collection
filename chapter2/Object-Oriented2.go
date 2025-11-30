package main

import "fmt"

// 1. 定义Person结构体（包含Name和Age字段）
type Person struct {
	Name string
	Age  int
}

// 2. 定义Employee结构体：组合Person（嵌套Person），并添加EmployeeID字段
type Employee struct {
	Person     // 组合Person结构体（匿名嵌套，可直接访问Person的字段）
	EmployeeID string
}

// 3. 为Employee实现PrintInfo()方法，输出员工信息
func (e Employee) PrintInfo() {
	// 可直接访问组合的Person的字段（e.Name等价于e.Person.Name）
	fmt.Println("=== 员工信息 ===")
	fmt.Printf("姓名：%s\n", e.Name)
	fmt.Printf("年龄：%d\n", e.Age)
	fmt.Printf("员工ID：%s\n", e.EmployeeID)
}

func main() {
	// 创建Employee实例（同时初始化Person的字段）
	emp := Employee{
		Person: Person{
			Name: "张三",
			Age:  28,
		},
		EmployeeID: "EMP001",
	}

	// 调用PrintInfo()方法输出信息
	emp.PrintInfo()
}
