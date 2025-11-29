package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 核心函数：原地删除有序数组重复项，返回新长度
func removeDuplicates(nums []int) int {
	// 数组为空则返回0
	if len(nums) == 0 {
		return 0
	}
	// 慢指针：标记唯一元素的最后位置
	slow := 0
	// 快指针：遍历数组找新元素
	for fast := 1; fast < len(nums); fast++ {
		// 快指针找到新元素（与慢指针元素不同）
		if nums[fast] != nums[slow] {
			slow++ // 慢指针前移
			nums[slow] = nums[fast] // 覆盖重复项，保存新元素
		}
	}
	// 新长度是慢指针+1（从0开始计数）
	return slow + 1
}

func main() {
	fmt.Println("=== 有序数组去重工具 ===")
	fmt.Println("说明：输入有序数组元素（英文逗号分隔），例如：1,1,2 或 0,0,1,1,1,2")
	fmt.Println("输入 q 可退出程序\n")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("请输入有序数组：")
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
		var nums []int
		valid := true
		for _, part := range parts {
			numStr := strings.TrimSpace(part)
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

		// 调用核心函数，获取新长度
		newLen := removeDuplicates(nums)
		// 输出结果（前newLen个元素是去重后的数组）
		fmt.Printf("✅ 结果：去重后新长度为 %d，去重后的数组前 %d 位是 %v\n\n", 
			newLen, newLen, nums[:newLen])
	}
}
