package fib

import (
	"errors"
	"fmt"
)

// 可以在一个 packahe 内写多个 init 方法，init 方法会在 main 函数之前执行
// 一个源码文件中可能会有多个 init 方法，它的执行顺序 golang 并没有规定
// 因此在代码过程中应该避免依赖 init 方法的顺序
// https://zhuanlan.zhihu.com/p/34211611
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

var ErrorFibNegNumber = errors.New("input is less than 0")

func Fib(n int) ([]int, error) {
	if n <= 0 {
		return nil, ErrorFibNegNumber
	}
	if n == 1 {
		return []int{1}, nil
	}
	if n == 2 {
		return []int{1, 2}, nil
	}
	var res = []int{1, 1}
	for i := 3; i <= n; i++ {
		res = append(res, res[i-2]+res[i-3])
	}
	return res, nil
}
