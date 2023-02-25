package main

import (
	"fmt"
	"os"
)

func main() {
	// s := Student{"1000", "胖虎", 20}
	s := &Student{"1000", "胖虎", 20} // 也可以调用两个函数
	fmt.Printf("%x\n", &s.name)
	fmt.Println("--------")
	s.String1()
	fmt.Println("--------")
	s.String2()
	fmt.Println("--------")
	os.Exit(0)
}

type Student struct {
	Id   string
	name string
	year int
}

// 这个节省空间，建议采用该方法
func (s *Student) String1() {
	fmt.Printf("%x\n", &s.name)
	fmt.Println(s.Id, s.name, s.year)
}

// 这里会产生对象复制
func (s Student) String2() {
	fmt.Printf("%x\n", &s.name)
	fmt.Println(s.Id, s.name, s.year)
}
