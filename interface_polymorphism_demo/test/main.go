package main

import (
	"fmt"
	"os"
)

func main() {
	t := T{}
	ret := t.foo()
	fmt.Println("return value is", ret)
	var i A = &t
	i.foo()
	os.Exit(0)
}

// 接口在 golang 中本身就是一个指针
type A interface {
	foo() string
}

type B interface {
	A // 匿名接口
	bar() string
}

type T struct {
}

func (t *T) foo() string {
	fmt.Println("T's foo method")
	return "foo"
}

// bar 既可以实现，也可以不实现
func (t *T) bar() string {
	fmt.Println("T's bar method")
	return "bar"
}
