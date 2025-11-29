package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 两数之和核心函数：返回和为target的两个数的下标
func twoSum(nums []int, target int) []int {
	// 用map存储“数值:对应的下标”
	numIndexMap := make(map[int]int)
	for idx, num := range nums {
		// 计算需要的补数
		complement := target - num
		// 若补数已在map中，直接返回两个下标
		if existIdx, ok := numIndexMap[complement]; ok {
			return []int{existIdx, idx}
		}
		// 否则将当前数值和下标存入map
		numIndexMap[num] = idx
	}
	// 题目保证有解，此处不会执行
	return nil
}

func main() {
	fmt.Println("=== 两数之和计算工具 ===")
	fmt.Println("说明：输入格式为「数组元素（英文逗号分隔）, 目标值」，例如：2,7,11,15, 9")
	fmt.Println("输入 q 可退出程序\n")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("请输入数组和目标值：")
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

		// 分割输入：前半部分是数组，最后一个是目标值
		parts := strings.Split(input, ",")
		if len(parts) < 2 {
			fmt.Println("❌ 错误：输入格式不正确！请按「数组元素, 目标值」输入\n")
			continue
		}

		// 解析数组
		var nums []int
		valid := true
		for i := 0; i < len(parts)-1; i++ {
			numStr := strings.TrimSpace(parts[i])
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("❌ 错误：「%s」不是有效的整数！\n\n", numStr)
				valid = false
				break
			}
			nums = append(nums, num)
		}
		if !valid {
			continue
		}

		// 解析目标值
		targetStr := strings.TrimSpace(parts[len(parts)-1])
		target, err := strconv.Atoi(targetStr)
		if err != nil {
			fmt.Printf("❌ 错误：「%s」不是有效的目标值！\n\n", targetStr)
			continue
		}

		// 计算并输出结果
		result := twoSum(nums, target)
		fmt.Printf("✅ 结果：和为 %d 的两个数下标是 %v\n\n", target, result)
	}
}
