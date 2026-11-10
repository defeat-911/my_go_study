package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		// 要点,必须确保defer 也就是wg.Done()动作在线程最后
		defer wg.Done()
		fmt.Println("work1 begin.")
		time.Sleep(2 * time.Second)
		fmt.Println("work1 done!")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("work2 begin.")
		time.Sleep(2 * time.Second)
		fmt.Println("work2 done!")
	}()

	wg.Wait() // 阻塞至所有 routine 结束
	fmt.Println("All works done!")
}
