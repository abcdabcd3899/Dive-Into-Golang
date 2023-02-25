package main

import (
	"fmt"
	"sync"
)

func main() {
	lock_mut_test()
	channel_to_lock()
}

func lock_mut_test() {
	var mut sync.Mutex
	wg := sync.WaitGroup{}
	count := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			mut.Lock()
			defer mut.Unlock() // 要注意这个 defer unlock 的位置
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count = ", count)
}

func channel_to_lock() {
	ch := make(chan int, 1) // using 1 size buffered channel as mutex
	wg := sync.WaitGroup{}
	count := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			ch <- 1
			count++
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count = ", count)
}
