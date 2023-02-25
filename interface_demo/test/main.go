package main

import (
	"fmt"
	"os"
)

func main() {
	a := &Cat{}
	var i Animal = a // Animal 对象的本质是一根指针
	fmt.Println("cat sound is", i.Sound())
	fmt.Println("cat sound is", a.Sound())
	WhatType(a)
	WhatType(i)
	fmt.Println("--------cat seperation--------")
	b := new(Dog)
	i = b
	fmt.Println("dog sound is", b.Sound())
	fmt.Println("dog sound is", i.Sound())
	WhatType(b)
	WhatType(i)
	os.Exit(0)
}

// interface
type Animal interface {
	Sound() string
}

func WhatType(a Animal) {
	fmt.Printf("%T\n", a)
}

type Cat struct {
}

func (c *Cat) Sound() string {
	return "miaomiao"
}

type Dog struct {
}

func (d *Dog) Sound() string {
	return "wangwang"
}
