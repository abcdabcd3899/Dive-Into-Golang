package main

import (
	"fmt"
	"sync"
)

// 单生产者+单消费者
func main() {
	cond := sync.NewCond(new(sync.Mutex))
	count := 0
	// 开启一个消费者线程，消费 count，即使得 count--
	go func() {
		// 死循环消费 count
		for {
			cond.L.Lock()
			for count == 0 {
				cond.Wait()
			}
			count--
			fmt.Println("count: ", count)
			cond.L.Unlock()
			cond.Signal()
		}
	}()
	// 主线程作为生产者，生产 count，即 count++
	for {
		cond.L.Lock()
		for count == 3 {
			cond.Wait()
		}
		count++
		fmt.Println("count: ", count)
		cond.L.Unlock()
		cond.Signal()
	}
}
