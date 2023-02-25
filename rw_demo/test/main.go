package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

// https://www.golinuxcloud.com/golang-mutex/
func main() {
	cond := sync.NewCond(new(sync.Mutex))
	var count int32 = 0
	go reader("read1", &count, cond)
	go reader("read2", &count, cond)
	go writer("writer1", &count, cond)
	go writer("writer2", &count, cond)
	time.Sleep(time.Second * 10)
	os.Exit(0)
}

func reader(str string, count *int32, cond *sync.Cond) {
	for {
		cond.L.Lock()
		for atomic.LoadInt32(count) <= 0 {
			cond.Wait()
		}
		atomic.AddInt32(count, -1)
		fmt.Println(str, "count:", *count)
		cond.L.Unlock()
		cond.Broadcast()
	}
}

func writer(str string, count *int32, cond *sync.Cond) {
	for {
		cond.L.Lock()
		for atomic.LoadInt32(count) >= 3 {
			cond.Wait()
		}
		atomic.AddInt32(count, 1)
		fmt.Println(str, "count:", *count)
		cond.L.Unlock()
		cond.Broadcast()
	}
}
