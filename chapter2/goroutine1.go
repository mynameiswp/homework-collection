package main

import (
	"fmt"
	"sync"
	"time"
)

func printOdds(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("奇数协程开始：")
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数：%d\n", i)
		time.Sleep(10 * time.Millisecond) // 加微小延迟，触发协程切换
	}
}

func printEvens(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("偶数协程开始：")
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数：%d\n", i)
		time.Sleep(10 * time.Millisecond) // 加微小延迟，触发协程切换
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printOdds(&wg)
	go printEvens(&wg)
	wg.Wait()
	fmt.Println("所有协程执行完毕")
}
