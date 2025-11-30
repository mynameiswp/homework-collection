package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义任务类型：接收任务名，返回执行耗时（单位：毫秒）
type Task func(taskName string) int64

// 任务调度器：接收任务列表，并发执行并统计每个任务的执行时间
func taskScheduler(tasks map[string]Task) {
	var wg sync.WaitGroup

	// 遍历所有任务，启动协程执行
	for name, task := range tasks {
		wg.Add(1) // 注册任务到WaitGroup
		go func(taskName string, t Task) {
			defer wg.Done() // 任务完成后通知

			// 统计任务执行时间
			start := time.Now()
			t(taskName) // 执行任务
			cost := time.Since(start).Milliseconds()

			fmt.Printf("任务「%s」执行完成，耗时：%dms\n", taskName, cost)
		}(name, task) // 注意：循环变量要通过参数传递，避免闭包引用问题
	}

	wg.Wait() // 等待所有任务完成
	fmt.Println("所有任务调度完成！")
}

func main() {
	// 示例任务1：模拟耗时任务（打印数字）
	task1 := func(name string) int64 {
		fmt.Printf("任务「%s」开始：打印1-5\n", name)
		for i := 1; i <= 5; i++ {
			fmt.Printf("任务「%s」输出：%d\n", name, i)
			time.Sleep(100 * time.Millisecond) // 模拟任务耗时
		}
		return 0
	}

	// 示例任务2：模拟另一个耗时任务（打印字母）
	task2 := func(name string) int64 {
		fmt.Printf("任务「%s」开始：打印A-E\n", name)
		for c := 'A'; c <= 'E'; c++ {
			fmt.Printf("任务「%s」输出：%c\n", name, c)
			time.Sleep(150 * time.Millisecond) // 模拟任务耗时
		}
		return 0
	}

	// 示例任务3：模拟短耗时任务
	task3 := func(name string) int64 {
		fmt.Printf("任务「%s」开始：计算1+1\n", name)
		result := 1 + 1
		fmt.Printf("任务「%s」结果：1+1=%d\n", name, result)
		return 0
	}

	// 构造任务列表
	tasks := map[string]Task{
		"数字打印": task1,
		"字母打印": task2,
		"简单计算": task3,
	}

	// 启动调度器
	fmt.Println("=== 任务调度器启动 ===")
	taskScheduler(tasks)
}
