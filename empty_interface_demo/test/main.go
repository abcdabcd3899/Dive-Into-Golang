package main

import (
	"fmt"
	"os"
)

func main() {
	DoSomething(10)
	DoSomething("10")
	os.Exit(0)
}

// https://go.dev/ref/spec#Type_assertions
func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("string", s)
		return
	}
	fmt.Println("Unknow Type")
	// switch v := p.(type) {
	// case int:
	// 	fmt.Println("Integer", v)
	// case string:
	// 	fmt.Println("String", v)
	// default:
	// 	fmt.Println("Unknow Type")
	// }
}
