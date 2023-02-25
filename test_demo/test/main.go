package main // 必须是 main 包才能编译 main 函数

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("hello,", os.Args[1])
	}
	os.Exit(0)
}
