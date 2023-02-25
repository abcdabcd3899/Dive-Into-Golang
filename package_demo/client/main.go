package main

import (
	fib "abcdabcd3899/package_demo/test"
	"fmt"

	"github.com/easierway/concurrent_map"
)

func main() {
	var n int = 5
	if res, err := fib.Fib(n); err == nil {
		fmt.Println(res)
	}
	fmt.Printf("concurrent_map.ConcurrentMap: %v\n", concurrent_map.ConcurrentMap{})
}
