// 编写一个 Go 程序，启动 3 个 goroutine，每个 goroutine 打印自己的编号（1, 2, 3）
// 并使用 sync.WaitGroup 确保主程序在所有 goroutine 完成后退出
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		fmt.Println("routine 1")
		defer wg.Done()
	}()

	go func() {
		fmt.Println("routine 2")
		defer wg.Done()
	}()

	go func() {
		fmt.Println("routine 3")
		defer wg.Done()
	}()
	// or for
	//for i := 0; i <= 3; i++ {
	//  go func() {
	//    wg.Add(1)
	//    defer wg.Done()
	//    fmt.Println("routine ", i)
	//  }()
	//}
	wg.Wait()

}
