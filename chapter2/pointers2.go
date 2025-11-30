package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 接收整数切片的指针，将每个元素×2
func doubleSlice(s *[]int) {
	// 解引用切片指针，遍历并修改原切片元素
	for i := range *s {
		(*s)[i] *= 2
	}
}

func main() {
	fmt.Println("=== 切片指针元素×2工具 ===")
	fmt.Println("📌 输入格式：数组元素（英文逗号分隔），例如：1,2,3；输入 q 退出")
	fmt.Println("----------------------------------------")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("请输入切片元素（输入q退出）：")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		// 处理退出
		if input == "q" || input == "Q" {
			fmt.Println("👋 程序已退出～")
			break
		}
		if input == "" {
			fmt.Println("❌ 错误：输入不能为空！")
			continue
		}

		// 解析输入为整数切片
		parts := strings.Split(input, ",")
		var nums []int
		valid := true
		for _, part := range parts {
			numStr := strings.TrimSpace(part)
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("❌ 错误：「%s」不是有效整数！\n", numStr)
				valid = false
				break
			}
			nums = append(nums, num)
		}
		if !valid {
			continue
		}

		// 调用函数（传递切片指针）
		fmt.Printf("✅ 修改前：%v → 修改后（×2）：", nums)
		doubleSlice(&nums)
		fmt.Printf("%v\n", nums)
		fmt.Println("----------------------------------------")
	}
}
