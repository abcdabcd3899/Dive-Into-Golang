package main

import (
	"errors"
	"os"
)

/*
golang 没有和其它语言一样提供 try...catch... finally 这样子的错误处理结构，
而是保持和 C/C++ 一样采用自定义错误的办法处理。其他语言的 try...catch...finally 结构
存在很大的隐患。举个栗子：Java 程序员采用 JDBC 操作 PostgreSQL 数据库，向数据库循环插入 100 万行
数据，并且为了防止错误发生，采用 try catch 包裹，每次插入数据库后端都会开启一个子事务，当插入没有
错误时，子事务全部提交，就产生 100 万个子事务，PostgreSQL 子事务结构当子事务数量大于 64 时，
会将这些结构落盘，这会导致大量的 IO，拖垮系统性能
*/

func main() {
	var n int = 10
	// for {
	// 	fmt.Scanf("%d", &n)
	// 	if res, err := Fib(n); err == nil {
	// 		fmt.Println(res)
	// 	} else {
	// 		fmt.Println(err.Error())
	// 	}
	// 	if n >= 0x3f3f3f3f {
	// 		break
	// 	}
	// }
	Fib(n)
	os.Exit(0)
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
