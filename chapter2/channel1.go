package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. 创建无缓冲通道（用于传递int类型数据）
	ch := make(chan int)
	var wg sync.WaitGroup

	// 2. 启动生产者协程：生成1-10的整数并发送到通道
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("生产者协程启动：开始生成1-10的整数")
		for i := 1; i <= 10; i++ {
			ch <- i // 向通道发送数据（无缓冲通道会阻塞，直到消费者接收）
			fmt.Printf("生产者发送：%d\n", i)
		}
		close(ch) // 数据发送完成后关闭通道（避免消费者阻塞）
		fmt.Println("生产者协程结束：通道已关闭")
	}()

	// 3. 启动消费者协程：从通道接收数据并打印
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("消费者协程启动：等待接收数据")
		// 循环接收通道数据（直到通道关闭且数据被取完）
		for num := range ch {
			fmt.Printf("消费者接收：%d\n", num)
		}
		fmt.Println("消费者协程结束：通道已无数据")
	}()

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("程序执行完毕")
}
