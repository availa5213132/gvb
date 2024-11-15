package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建两个无缓冲的通道，分别用于控制奇数和偶数的打印顺序
	oddCh := make(chan struct{})
	evenCh := make(chan struct{})

	// 打印奇数的 Goroutine
	go func() {
		for i := 1; i <= 10; i += 2 {
			<-oddCh              // 等待 oddCh 通知
			fmt.Println(i)       // 打印奇数
			evenCh <- struct{}{} // 通知偶数 Goroutine
		}
	}()

	// 打印偶数的 Goroutine
	go func() {
		for i := 2; i <= 10; i += 2 {
			<-evenCh            // 等待 evenCh 通知
			fmt.Println(i)      // 打印偶数
			oddCh <- struct{}{} // 通知奇数 Goroutine
		}
	}()

	// 开始打印，首先触发奇数
	oddCh <- struct{}{}

	// 让主程序等待足够时间让打印完成
	time.Sleep(1 * time.Second)
}
