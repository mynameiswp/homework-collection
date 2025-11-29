package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 合并区间核心函数
func merge(intervals [][]int) [][]int {
	// 空数组直接返回
	if len(intervals) == 0 {
		return nil
	}

	// 按区间的start升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果集，放入第一个区间
	result := [][]int{intervals[0]}

	// 遍历剩余区间
	for i := 1; i < len(intervals); i++ {
		// 取结果集最后一个区间
		last := result[len(result)-1]
		// 当前区间的start <= 最后区间的end → 合并
		if intervals[i][0] <= last[1] {
			// 合并后的end取两者的较大值
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else {
			// 不重叠，直接加入结果集
			result = append(result, intervals[i])
		}
	}

	return result
}

func main() {
	fmt.Println("=== 合并区间工具 ===")
	fmt.Println("说明：输入格式为「区间1,区间2,...」，每个区间用-分隔，例如：1-3,2-6,8-10,15-18")
	fmt.Println("输入 q 可退出程序\n")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("请输入区间集合：")
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

		// 解析输入为区间数组
		intervalStrs := strings.Split(input, ",")
		var intervals [][]int
		valid := true
		for _, s := range intervalStrs {
			s = strings.TrimSpace(s)
			// 分割区间的start和end（用-分隔）
			se := strings.Split(s, "-")
			if len(se) != 2 {
				fmt.Printf("❌ 错误：区间「%s」格式不正确（应为start-end）！\n\n", s)
				valid = false
				break
			}
			// 解析start和end
			start, err1 := strconv.Atoi(strings.TrimSpace(se[0]))
			end, err2 := strconv.Atoi(strings.TrimSpace(se[1]))
			if err1 != nil || err2 != nil || start > end {
				fmt.Printf("❌ 错误：区间「%s」无效（start需≤end且为整数）！\n\n", s)
				valid = false
				break
			}
			intervals = append(intervals, []int{start, end})
		}
		if !valid {
			continue
		}

		// 合并区间并输出结果
		merged := merge(intervals)
		fmt.Printf("✅ 结果：合并后的区间是 %v\n\n", merged)
	}
}
