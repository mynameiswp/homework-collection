package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 核心函数：接收整数指针，值+10
func addTen(num *int) {
	*num += 10
}

func main() {
	fmt.Println("=== 整数指针加10工具 ===")
	fmt.Println("📌 输入整数即可计算+10后的结果；输入 q / Q 可退出程序")
	fmt.Println("----------------------------------------")

	// 创建scanner（复用避免重复初始化）
	scanner := bufio.NewScanner(os.Stdin)

	// 循环接收输入
	for {
		fmt.Print("请输入（输入q退出）：")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		// 1. 处理退出指令
		if input == "q" || input == "Q" {
			fmt.Println("👋 程序已退出，下次见～")
			break
		}

		// 2. 处理空输入
		if input == "" {
			fmt.Println("❌ 错误：输入不能为空，请输入有效整数！")
			continue
		}

		// 3. 解析输入为整数（校验有效性）
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("❌ 错误：「%s」不是有效整数，请重新输入！\n", input)
			continue
		}

		// 4. 执行加10操作并输出结果
		fmt.Printf("✅ 修改前：%d → 修改后（+10）：", num)
		addTen(&num)
		fmt.Printf("%d\n", num)
		fmt.Println("----------------------------------------")
	}

	// 捕获scanner可能的错误（非必需，但更健壮）
	if err := scanner.Err(); err != nil {
		fmt.Printf("❌ 读取输入时出错：%v\n", err)
	}
}
