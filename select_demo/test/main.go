package main

import (
	"fmt"
	"sync"
)

// func AsyncServices() chan int {
// 	ch := make(chan int, 1)
// 	go func() {
// 		time.Sleep(time.Millisecond * 50) // 不管有没有这个走的都是default
// 		ch <- 1
// 	}()
// 	return ch
// }

// func main() {
// 	select {
// 	case ch := <-AsyncServices():
// 		fmt.Println(ch)
// 	case <-time.After(time.Millisecond * 100):
// 		fmt.Println("time out")
// 	default:
// 		fmt.Println("Receive nothing.")
// 	}
// }

// It's not guaranteed that the code output will be "The default case!" as the order of execution of the goroutines is non-deterministic.
// The select statement will select the first case that is ready to receive a value, either ch1 or ch2 will receive a value, and the corresponding case will be executed.
// If both ch1 and ch2 are not ready to receive a value, the default case will be executed.
// However, since the order of execution of the goroutines is non-deterministic, it is possible that one of the non-default cases will be executed before the default case.
// https://golangdocs.com/select-statement-in-golang 下面这个例子来自这个链接，里面的解释是错误的
// func g1(ch chan int) {
// 	ch <- 42
// }

// func g2(ch chan int) {
// 	ch <- 43
// }

// func main() {

// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go g1(ch1)
// 	go g2(ch2)

// 	select {
// 	case v1 := <-ch1:
// 		fmt.Println("Got: ", v1)
// 	case v2 := <-ch2:
// 		fmt.Println("Got: ", v2)
// 	default:
// 		fmt.Println("The default case!")
// 	}
// }

// func main() {
// 	ch := make(chan int, 1)
// 	go func() {
// 		time.Sleep(time.Millisecond * 1)
// 		ch <- 1
// 	}()
// 	go func() {
// 		time.Sleep(time.Millisecond * 1)
// 		ch <- 1
// 	}()
// 	select {
// 	case ch := <-ch:
// 		fmt.Println(ch)
// 	case <-time.After(time.Millisecond * 100):
// 		fmt.Println("Time out.")
// 	default:
// 		fmt.Println("Receive nothing.")
// 	}
// }

// func main() {
// 	var wg *sync.WaitGroup = new(sync.WaitGroup)
// 	var c = make(chan int)
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		c <- 12
// 	}()
// 	select {
// 	case x := <-c:
// 		fmt.Println("received", x)
// 	default:
// 		fmt.Println("default case executed")
// 	}
// 	wg.Wait()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	done := make(chan bool)
// 	c := make(chan int, 1)
// 	wg.Add(1)
// 	go func() {
// 		c <- 12
// 		done <- true
// 		wg.Done()
// 	}()
// 	select {
// 	case x := <-c:
// 		fmt.Println("received", x)
// 	case <-done:
// 		fmt.Println("Done")
// 	}
// 	wg.Wait()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	done := make(chan bool)
// 	c := make(chan int)
// 	wg.Add(1)
// 	go func() {
// 		c <- 12
// 		done <- true
// 		wg.Done()
// 	}()
// 	select {
// 	case x := <-c:
// 		fmt.Println("received", x)
// 	case <-done:
// 		fmt.Println("Done")
// 	default:
// 		fmt.Println("default case executed")
// 	}
// 	wg.Wait()
// }

func main() {
	var wg sync.WaitGroup
	done := make(chan bool)
	c := make(chan int)
	wg.Add(1)
	go func() {
		c <- 12
		done <- true
		wg.Done()
	}()
	select {
	case x := <-c:
		fmt.Println("received", x)
	}
	select {
	case <-done:
		fmt.Println("Done")
	}
	wg.Wait()
}
