package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 核心函数：大整数加一
func plusOne(digits []int) []int {
	n := len(digits)
	// 从最后一位（最低位）开始处理
	for i := n - 1; i >= 0; i-- {
		// 当前位加1
		digits[i]++
		// 取模10，若不为0说明无进位，直接返回
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	// 若所有位都是9（比如[9,9]），需要新增一位1在开头
	return append([]int{1}, digits...)
}

func main() {
	fmt.Println("=== 大整数加一工具 ===")
	fmt.Println("说明：输入数组元素（英文逗号分隔），例如：1,2,3 或 9")
	fmt.Println("输入 q 可退出程序\n")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("请输入数字数组：")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		// 输入q退出
		if input == "q" || input == "Q" {
			fmt.Println("程序已退出～")
			break
		}
		if input == "" {
			fmt.Println("❌ 错误：请输入有效内容！\n")
			continue
		}

		// 解析数组
		parts := strings.Split(input, ",")
		var digits []int
		valid := true
		for _, part := range parts {
			numStr := strings.TrimSpace(part)
			num, err := strconv.Atoi(numStr)
			if err != nil || num < 0 || num > 9 {
				fmt.Printf("❌ 错误：「%s」不是有效的数字（需为0-9的整数）！\n\n", numStr)
				valid = false
				break
			}
			digits = append(digits, num)
		}
		if !valid {
			continue
		}

		// 计算并输出结果
		result := plusOne(digits)
		fmt.Printf("✅ 结果：加1后的数组是 %v\n\n", result)
	}
}
