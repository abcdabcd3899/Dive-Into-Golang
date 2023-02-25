package main

import (
	"fmt"
	"os"
	"time"
)

// 函数在 golang 中是一等公民，即 function 可以作为参数和返回值等
func main() {
	f := timeSpent(test)
	fmt.Println("The function return value is", f(10))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	TestDefer()
	os.Exit(0)
}

// 函数同时作为参数和返回值
func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent is", time.Since(start).Seconds())
		return ret
	}
}

func test(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func sum(ops ...int) int {
	sum := 0
	for _, ele := range ops {
		sum += ele
	}
	return sum
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer() {
	defer Clear()
	fmt.Println("Start")
	// panic("err")
}
