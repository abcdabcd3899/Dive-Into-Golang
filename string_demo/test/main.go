package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// https://go.dev/blog/strings
	var s string = "中"
	fmt.Println([]rune(s)) // unicode code point
	fmt.Println([]byte(s)) // UTF-8 byte, 从输出我们知道一个中文字需要三个 byte

	// https://pkg.go.dev/strings
	s = "a, b, c"
	sl := strings.Split(s, ", ")
	fmt.Println(sl)

	// https://pkg.go.dev/strconv
	s = "10"
	if i, err := strconv.Atoi(s); err == nil {
		fmt.Println(i)
	}
}
