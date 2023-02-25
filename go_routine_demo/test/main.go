package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go PrintGoRoutine(i)
	}
	time.Sleep(time.Microsecond * 50)
	os.Exit(0)
}

// 使用 goroutine 平行打印一个整数
func PrintGoRoutine(i int) {
	fmt.Println(i)
}
