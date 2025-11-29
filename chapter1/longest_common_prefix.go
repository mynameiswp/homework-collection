
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 核心函数：计算最长公共前缀
func longestCommonPrefix(strs []string) string {
	// 边界条件：空数组直接返回空
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串为初始前缀
	prefix := strs[0]

	// 遍历剩余字符串，逐步缩短前缀
	for i := 1; i < len(strs); i++ {
		j := 0
		// 逐字符比较，直到出现不匹配或任一字符串结束
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}
		// 截断前缀到匹配的长度
		prefix = prefix[:j]

		// 前缀为空时提前退出
		if prefix == "" {
			break
		}
	}

	return prefix
}

func main() {
	fmt.Println("=== 最长公共前缀计算工具 ===")
	fmt.Println("说明：输入多个字符串（用英文逗号分隔，比如 flower,flow,flight）")
	fmt.Println("输入 q 可退出程序\n")

	// 创建输入扫描器，支持空格/特殊字符
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// 提示输入
		fmt.Print("请输入字符串列表：")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		// 输入q退出
		if input == "q" || input == "Q" {
			fmt.Println("程序已退出～")
			break
		}

		// 空输入处理
		if input == "" {
			fmt.Println("❌ 错误：请输入至少一个字符串！\n")
			continue
		}

		// 分割输入为字符串数组（按英文逗号分割）
		strList := strings.Split(input, ",")
		// 去除每个字符串前后的空格（兼容用户输入 "flower, flow, flight" 这种情况）
		for i := range strList {
			strList[i] = strings.TrimSpace(strList[i])
		}

		// 计算并输出结果
		result := longestCommonPrefix(strList)
		if result == "" {
			fmt.Printf("✅ 结果：输入的字符串无公共前缀\n\n")
		} else {
			fmt.Printf("✅ 结果：最长公共前缀为 \"%s\"\n\n", result)
		}
	}
}
