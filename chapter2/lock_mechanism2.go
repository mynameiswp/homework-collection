package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// 1. 定义原子计数器（必须用int64类型，对应atomic的函数）
	var count int64
	var wg sync.WaitGroup

	// 2. 启动10个协程，每个协程执行1000次原子递增
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 每个协程执行1000次原子递增
			for j := 0; j < 1000; j++ {
				// atomic.AddInt64：原子地将count加1，返回新值
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	// 3. 等待所有协程完成
	wg.Wait()
	// 原子读取最终值（也可以直接用count，这里用LoadInt64更规范）
	fmt.Printf("最终计数器的值：%d\n", atomic.LoadInt64(&count)) // 预期输出：10*1000=10000
}
