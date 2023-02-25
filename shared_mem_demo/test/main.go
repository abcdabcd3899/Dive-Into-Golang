package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	lock_mut_test()
	lock_rwlock_test()
	atomic_test()
	// atomic_variable_test()
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

func lock_rwlock_test() {
	var rwlock sync.RWMutex = sync.RWMutex{}
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	count := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			rwlock.Lock()
			defer rwlock.Unlock() // 要注意这个 defer unlock 的位置
			count++
			// panic("Something Wrong!")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count = ", count)
}

func atomic_test() {
	var count int32 = 0
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&count, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count = ", count)
}

// https://gfw.go101.org/article/concurrent-atomic-operation.html
// github action 的 golang 版本为 1.18, 1.19 之后才开始支持 atomic.xxx 这些类型
// 因此注释下面这段代码
// func atomic_variable_test() {
// 	var count atomic.Int64
// 	var wg sync.WaitGroup = sync.WaitGroup{}
// 	for i := 0; i < 5000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			count.Add(1)
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("count = ", count.Load())
// }
