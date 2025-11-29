package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 修正核心验证函数：确保参数正确且逻辑无误
func isValid(s string) bool {
	if s == "" { // 空字符串视为有效
		return true
	}
	// 括号映射表：右括号对应左括号
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []rune{} // 用切片模拟栈

	for _, char := range s {
		// 如果是右括号，检查是否与栈顶左括号匹配
		if targetLeft, ok := bracketMap[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != targetLeft {
				return false // 栈空或不匹配 → 无效
			}
			stack = stack[:len(stack)-1] // 匹配成功，弹出栈顶
		} else if char == '(' || char == '[' || char == '{' {
			// 左括号直接入栈
			stack = append(stack, char)
		} else {
			// 包含非括号字符 → 无效
			return false
		}
	}
	// 所有左括号必须被匹配（栈为空才有效）
	return len(stack) == 0
}

func main() {
	fmt.Println("=== 括号验证工具（输入 q 退出） ===")
	scanner := bufio.NewScanner(os.Stdin) // 更可靠的输入读取方式

	for {
		// 循环提示输入
		fmt.Print("请输入括号字符串：")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text()) // 去除输入前后的空格

		// 输入 q 或 Q 退出程序
		if input == "q" || input == "Q" {
			fmt.Println("程序已退出～")
			break
		}

		// 验证并输出结果
		if isValid(input) {
			fmt.Printf("✅ 结果：\"%s\" 是有效的括号\n\n", input)
		} else {
			fmt.Printf("❌ 结果：\"%s\" 不是有效的括号\n\n", input)
		}
	}
}
