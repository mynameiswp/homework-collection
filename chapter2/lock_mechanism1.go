package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. 定义共享计数器和互斥锁
	var count int
	var mu sync.Mutex
	var wg sync.WaitGroup

	// 2. 启动10个协程，每个协程执行1000次递增
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 每个协程执行1000次递增
			for j := 0; j < 1000; j++ {
				mu.Lock()         // 加锁：保护共享变量count
				count++           // 安全修改共享计数器
				mu.Unlock()       // 解锁：释放锁，让其他协程可以访问
			}
		}()
	}

	// 3. 等待所有协程完成
	wg.Wait()
	fmt.Printf("最终计数器的值：%d\n", count) // 预期输出：10*1000=10000
}
